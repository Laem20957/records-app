package configuration

import (
	"time"

	"github.com/Laem20957/records-app/pkg/logger"
	"github.com/spf13/viper"
)

var logs = logger.CreateLogs()

type Config struct {
	TTL             time.Duration `mapstructure:"ttl"`
	TokenTTL        time.Duration `mapstructure:"token_ttl"`
	RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl"`
	LocalServerHost string        `mapstructure:"local_server_host"`
	LocalServerPort int           `mapstructure:"local_server_port"`
	Salt            string        `mapstructure:"salt"`
	SigningKey      string        `mapstructure:"sign_key"`
	PSQLHost        string        `mapstructure:"psql_host"`
	PSQLPort        int           `mapstructure:"psql_port"`
	PSQLDBName      string        `mapstructure:"psql_dbname"`
	PSQLUsername    string        `mapstructure:"psql_username"`
	PSQLPassword    string        `mapstructure:"psql_password"`
	PSQLModeSSL     string        `mapstructure:"psql_mode_ssl"`
}

func (config *Config) getConfigENV() {
	viper.AddConfigPath("records_app/configuration")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		logs.Fatal(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		logs.Fatal(err)
	}
}

func (config *Config) getConfigYAML() {
	viper.AddConfigPath("records_app/configuration")
	viper.SetConfigName("env")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		logs.Fatal(err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		logs.Fatal(err)
	}
}

func InitConfigs() *Config {
	call := &Config{}
	call.getConfigENV()
	call.getConfigYAML()
	return call
}
