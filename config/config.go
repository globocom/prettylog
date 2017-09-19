package config

import (
	"sync"

	"strings"

	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	settings Settings
	mu       sync.RWMutex
)

func Load(verbose bool) {
	setDefaults()

	viper.SetConfigType("yaml")
	viper.SetConfigName(".prettylog")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/")
	err := viper.ReadInConfig()
	if err != nil && verbose {
		fmt.Fprintf(os.Stderr, "error: failed to read configuration file: %v\n", err)
	}

	updateSettings(verbose)

	viper.OnConfigChange(func(e fsnotify.Event) {
		updateSettings(verbose)
	})
	viper.WatchConfig()
}

func updateSettings(verbose bool) {
	mu.Lock()
	defer mu.Unlock()

	err := viper.Unmarshal(&settings)
	if err != nil && verbose {
		fmt.Fprintf(os.Stderr, "error: failed to load configuration file: %v\n", err)
	}
}

func GetSettings() *Settings {
	mu.RLock()
	defer mu.RUnlock()

	return &settings
}

func setDefaults() {
	// Common settings
	setFieldDefaults("timestamp", "ts", true, 0, color.FgYellow, color.Faint)
	setFieldDefaults("logger", "logger", true, 10, color.FgWhite, color.Faint)
	setFieldDefaults("caller", "caller", false, 20, color.FgWhite, color.Faint)
	setFieldDefaults("level", "level", true, 5)
	setFieldDefaults("message", "msg", true, 40)

	// Level-specific settings
	viper.SetDefault("level.colors.debug", []color.Attribute{color.FgMagenta})
	viper.SetDefault("level.colors.info", []color.Attribute{color.FgBlue})
	viper.SetDefault("level.colors.warn", []color.Attribute{color.FgYellow})
	viper.SetDefault("level.colors.error", []color.Attribute{color.FgRed})
	viper.SetDefault("level.colors.fatal", []color.Attribute{color.FgRed, color.Bold})
}

func setFieldDefaults(name string, key string, visible bool, padding int, colorAttrs ...color.Attribute) {
	viper.SetDefault(name+".key", key)
	viper.SetDefault(name+".visible", visible)
	viper.SetDefault(name+".padding", padding)
	viper.SetDefault(name+".color", colorAttrs)
}

type Field struct {
	Key     string
	Visible bool
	Padding int
	Color   []color.Attribute
}

type Settings struct {
	Timestamp TimestampField
	Logger    Field
	Caller    Field
	Level     LevelField
	Message   Field
}

type TimestampField struct {
	Field  `mapstructure:",squash"`
	Format string
}

type LevelField struct {
	Field  `mapstructure:",squash"`
	Colors map[string][]color.Attribute
}

func (f *LevelField) GetColorAttr(level string) []color.Attribute {
	if c, exists := f.Colors[strings.ToLower(level)]; exists {
		return c
	} else {
		return []color.Attribute{}
	}
}
