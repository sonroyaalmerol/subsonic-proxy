package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Env  string `mapstructure:"env"`
}

type DatabaseConfig struct {
	URL     string `mapstructure:"url"`
	MaxConn int    `mapstructure:"max_conn"`
}

var AppConfig = &Config{}

func init() {
	_, _ = Load()
}

func getSystemConfigPath() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("ProgramData"), "subsonic-proxy")
	}
	return "/etc/subsonic-proxy"
}

func Load() (*Config, error) {
	configDir := getSystemConfigPath()
	confD := filepath.Join(configDir, "config.d")

	viper.SetConfigType("toml")
	viper.SetEnvPrefix("SUBSONIC_PROXY")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.env", "production")

	mainConfig := filepath.Join(configDir, "config.toml")
	viper.SetConfigFile(mainConfig)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Printf("Main config not found at %s, proceeding with defaults", mainConfig)
		} else {
			return nil, fmt.Errorf("error reading main config: %w", err)
		}
	}

	if err := mergeConfigDir(confD); err != nil {
		log.Printf("Warning: failed to merge config.d: %v", err)
	}

	if err := viper.Unmarshal(AppConfig); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config updated: %s", e.Name)
		_ = mergeConfigDir(confD)
		if err := viper.Unmarshal(AppConfig); err != nil {
			log.Printf("Error updating config: %v", err)
		}
	})
	viper.WatchConfig()

	return AppConfig, nil
}

func mergeConfigDir(dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, f := range files {
		if !f.IsDir() && filepath.Ext(f.Name()) == ".toml" {
			fullPath := filepath.Join(dir, f.Name())
			viper.SetConfigFile(fullPath)
			if err := viper.MergeInConfig(); err != nil {
				log.Printf("Error merging %s: %v", f.Name(), err)
			} else {
				// TODO: We also want to watch these supplemental files
			}
		}
	}
	return nil
}
