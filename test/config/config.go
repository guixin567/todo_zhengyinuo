package config

type TestConfig struct {
	*ServiceConfig `mapstructure:"service"`
}

type ServiceConfig struct {
	AppMode string `mapstructure:"app_mode"`
}
