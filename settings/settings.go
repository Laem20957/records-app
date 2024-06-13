package settings

import (
	"time"

	"records-app/internal/logger"

	"github.com/spf13/viper"
)

var logs = logger.CreateLogs()

type Settings struct {
	Debug           bool          `mapstructure:"DEBUG"`
	AppHost         string        `mapstructure:"APP_HOST"`
	AppPort         int           `mapstructure:"APP_PORT"`
	DBHost          string        `mapstructure:"DB_HOST"`
	DBPort          int           `mapstructure:"DB_PORT"`
	DBName          string        `mapstructure:"DB_NAME"`
	DBSchemaName    string        `mapstructure:"DB_SCHEMA_NAME"`
	DBUsername      string        `mapstructure:"DB_USERNAME"`
	DBPassword      string        `mapstructure:"DB_PASSWORD"`
	TTL             time.Duration `mapstructure:"TTL"`
	TTLToken        time.Duration `mapstructure:"TTL_TOKEN"`
	TTLRefreshToken time.Duration `mapstructure:"TTL_REFRESH_TOKEN"`
	Salt            string        `mapstructure:"SALT"`
	SignKey         string        `mapstructure:"SING_KEY"`
}

func GetSettings() *Settings {
	var settings Settings

	viper.AddConfigPath("settings")
	viper.SetConfigName("env")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logs.Log().Fatal(err)
	}

	if err := viper.Unmarshal(&settings); err != nil {
		logs.Log().Fatal(err)
	}
	return &settings
}
