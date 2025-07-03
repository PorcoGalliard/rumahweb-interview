package config

import (
	"log"

	"github.com/spf13/viper"
)

type loaderOption struct {
	ConfigPath string
	ConfigFile string
	ConfigType string
}

type LoaderOption func(*loaderOption)

func WithConfigPath(configPath string) LoaderOption {
	return func(lo *loaderOption) {
		lo.ConfigPath = configPath
	}
}

func WithConfigFile(configFile string) LoaderOption {
	return func(lo *loaderOption) {
		lo.ConfigFile = configFile
	}
}

func WithConfigType(configType string) LoaderOption {
	return func(lo *loaderOption) {
		lo.ConfigType = configType
	}
}

func LoadConfig[T any] (target *T, opts ...LoaderOption) *T {
	opt := &loaderOption{
		ConfigPath: ".",
		ConfigFile: "config",
		ConfigType: "yaml",
	}

	for _, loaderFunc := range opts {
		loaderFunc(opt)
	}

	viper.AddConfigPath(opt.ConfigPath)
	viper.SetConfigName(opt.ConfigFile)
	viper.SetConfigType(opt.ConfigType)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("❌ Failed reading config: %v", err)
	}

	if err := viper.Unmarshal(&target); err != nil {
		log.Fatalf("❌ Failed unmarshal config: %v", err)
	}

	log.Printf("✅ Config loaded from %s.%s\n", opt.ConfigFile, opt.ConfigType)
	return target
}