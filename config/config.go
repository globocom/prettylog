package config

import (
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	settings Settings
	mu       sync.RWMutex
)

func Load() error {
	setFieldDefaults("timestamp", &Field{"time", true, 0})
	setFieldDefaults("logger", &Field{"logger", true, 0})
	setFieldDefaults("caller", &Field{"caller", false, 0})
	setFieldDefaults("level", &Field{"level", true, 0})
	setFieldDefaults("message", &Field{"msg", true, 0})

	viper.SetConfigType("yaml")
	viper.SetConfigName(".prettylog")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = updateSettings()
	if err != nil {
		return err
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		updateSettings()
	})
	viper.WatchConfig()

	return nil
}

func updateSettings() error {
	mu.Lock()
	defer mu.Unlock()

	err := viper.Unmarshal(&settings)
	if err != nil {
		return err
	}
	return nil
}

func GetSettings() *Settings {
	mu.RLock()
	defer mu.RUnlock()

	return &settings
}

func setFieldDefaults(name string, field *Field) {
	viper.SetDefault(name+".key", field.Key)
	viper.SetDefault(name+".visible", field.Visible)
	viper.SetDefault(name+".padding", field.Padding)
}

type Settings struct {
	Timestamp TimestampField
	Logger    Field
	Caller    Field
	Level     Field
	Message   Field
}

type TimestampField struct {
	Field  `mapstructure:",squash"`
	Format string
}

type Field struct {
	Key     string
	Visible bool
	Padding int
}
