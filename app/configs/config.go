package configs

import (
	"github.com/spf13/viper"
	"log"
)

type MongoConfig struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`
}

type MCPClient struct {
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
	BaseUrl string `mapstructure:"base_url" json:"base_url" yaml:"base_url"`
	Address string `mapstructure:"address" json:"address" yaml:"address"`
	SSE     bool   `mapstructure:"sse" json:"sse" yaml:"sse"`
}

type Config struct {
	Mongo MongoConfig `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	MCP   MCPClient   `mapstructure:"mcp" json:"mcp" yaml:"mcp"`
}

func LoadConfig(path string, env string) Config {
	var Cfg Config
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config error: %v", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("Unmarshal config error: %v", err)
	}

	if env != "" {
		viper.SetConfigName("config." + env)
		if err := viper.MergeInConfig(); err != nil {
			log.Fatalf("Merge config error: %v", err)
		}
		if err := viper.Unmarshal(&Cfg); err != nil {
			log.Fatalf("Unmarshal config error: %v", err)
		}
	}

	log.Println("Config loaded successfully")
	return Cfg
}
