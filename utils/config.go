package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment           string        `mapstructure:"ENVIRONMENT"`
	DBDriver              string        `mapstructure:"DB_DRIVER"`
	DBSource              string        `mapstructure:"DB_SOURCE"`
	MigrationURL          string        `mapstructure:"MIGRATION_URL"`
	RedisAddress          string        `mapstructure:"REDIS_ADDRESS"`
	GIN_SERVER_ADDRESS    string        `mapstructure:"GIN_SERVER_ADDRESS"`
	HTTPServerAddress     string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress     string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey     string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration   time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration  time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EMAIL_SENDER_NAME     string        `mapstructure:"EMAIL_SENDER_NAME"`
	EMAIL_SENDER_ADDRESS  string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EMAIL_SENDER_PASSWORD string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
	EMAIL_TO_ADDRESS      string        `mapstructure:"EMAIL_TO_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
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

func LoadConfigENV() (config Config, err error) {
	time, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_DURATION"))
	if err != nil {
		fmt.Println(err)
		return
	}
	env := Config{
		DBDriver:            os.Getenv("DB_DRIVER"),
		DBSource:            os.Getenv("DB_SOURCE"),
		HTTPServerAddress:   os.Getenv("SERVER_ADDRESS"),
		TokenSymmetricKey:   os.Getenv("TOKEN_SYMMETRIC_KEY"),
		RedisAddress:        os.Getenv("REDIS_ADDRESS"),
		AccessTokenDuration: time,
		MigrationURL:        os.Getenv("MIGRATION_URL"),
	}
	return env, err
}
