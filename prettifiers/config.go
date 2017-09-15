package prettifiers

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName(".prettylog")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/")
	viper.ReadInConfig()

	viper.SetDefault("fields.timestamp", "time")
	viper.SetDefault("fields.logger", "logger")
	viper.SetDefault("fields.level", "level")
	viper.SetDefault("fields.caller", "caller")
	viper.SetDefault("fields.message", "msg")

	viper.SetDefault("show.timestamp", true)
	viper.SetDefault("show.caller", false)
}
