package config

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment string `mapstructure:"ENVIRONMENT"`

	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`

	DBMaxOpenConnections int    `mapstructure:"DB_MAX_OPEN_CONNECTIONS"`
	DBMaxIdleConnections int    `mapstructure:"DB_MAX_IDLE_CONNECTIONS"`
	DBMaxIdleTime        string `mapstructure:"DB_MAX_IDLE_TIME"`
	MigrationURL         string `mapstructure:"MIGRATION_URL"`

	HTTPServerAddress       string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	HTTPServerPort          string        `mapstructure:"HTTP_SERVER_PORT"`
	ReadHeaderTimeout       time.Duration `mapstructure:"READ_HEADER_TIMEOUT"`
	GracefulShutdownTimeout time.Duration `mapstructure:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	RunSwagger              bool          `mapstructure:"RUN_SWAGGER"`
	RequestLog              bool          `mapstructure:"REQUEST_LOG"`

	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
