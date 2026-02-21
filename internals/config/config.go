package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Host           string `mapstructure:"host"`
	Port           string `mapstructure:"port"`
	StatsVizPort   string `mapstructure:"stats_viz_port"`
	Log            `mapstructure:",squash"`
	RedisConfig    `mapstructure:",squash"`
	DatabaseConfig `mapstructure:",squash"`
}

type Log struct {
	Level string `mapstructure:"log_level"`
}

type RedisConfig struct {
	Address string `mapstructure:"redis_address"`
}

type DatabaseConfig struct {
	Host         string `mapstructure:"db_host"`
	Port         string `mapstructure:"db_port"`
	User         string `mapstructure:"db_user"`
	Password     string `mapstructure:"db_password"`
	Name         string `mapstructure:"db_name"`
	SSLMode      string `mapstructure:"db_sslmode"`
	MaxOpenConns string `mapstructure:"db_maxopenconns"`
	MaxIdleConns string `mapstructure:"db_maxidleconns"`
}

var Cfg *Config

func Load() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("no .env found in the folder structure")
	}

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", "8080")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (db *DatabaseConfig) GetDSN() string {
	sslMode := "disable"
	if db.SSLMode == "true" {
		sslMode = "require"
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", db.Host, db.Port, db.User, db.Password, db.Name, sslMode)
}
