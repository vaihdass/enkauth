package platform

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

const (
	configPathFlagKey = "config"
	configPathEnvKey  = "CONFIG_PATH"
)

type Config struct {
	Env     string        `yaml:"env" env-required:"true"`
	Auth    AuthConfig    `yaml:"auth" env-required:"true"`
	Storage StorageConfig `yaml:"storage" required:"true"`
	GRPC    GRPCConfig    `yaml:"grpc" required:"true"`
}

type AuthConfig struct {
	TokenTTL time.Duration `yaml:"token_ttl" env-required:"true"`
}

type StorageConfig struct {
	DSN string `yaml:"dsn" env-required:"true"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-required:"true"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func LoadConfig() (*Config, error) {
	const op = "internal.platform.LoadConfig"

	cfgPath := fetchConfigPath()
	if cfgPath == "" {
		return nil, fmt.Errorf("%s: config file path is empty", op)
	}
	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("%s: config file with path '%s' does not exist", op, cfgPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		return nil, fmt.Errorf("%s: config file read error: %w", op, err)
	}

	return &cfg, nil
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default. Default value is empty string.
func fetchConfigPath() string {
	var res string

	// example: --config="path/to/config.yaml"
	flag.StringVar(&res, configPathFlagKey, "", "config file path")
	flag.Parse()

	// example: CONFIG_PATH=./path/to/config.yaml app
	if res == "" {
		res = os.Getenv(configPathEnvKey)
	}

	return res
}
