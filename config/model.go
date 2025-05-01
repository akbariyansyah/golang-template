package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Dbname   string
	Sslmode  string
}

type Configuration struct {
	Address    int
	Database   DatabaseConfig
	JwtConfig  JwtConfig
	SmtpConfig SmtpConfig
}

type JwtConfig struct {
	SecretKey       string
	TokenExpiration string
}

type SmtpConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

type LogConfig struct {
	FileLocation    string
	FileTDRLocation string
	FileMaxSize     int
	FileMaxBackup   int
	FileMaxAge      int
	Stdout          bool
}

func (c *Configuration) SetConfig(path string) error {
	viper.AddConfigPath(".")
	viper.SetConfigName(path)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("fatal error config file herreeeeee: %w", err)
	}

	// Decode the configuration
	if err := viper.Unmarshal(c); err != nil {
		return err
	}

	return nil
}
