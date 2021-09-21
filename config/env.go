package config

import "github.com/kelseyhightower/envconfig"

// config singleton, populated from environment by init function
var config Config

func init() {
	err := envconfig.Process("", &config)
	if err != nil {
		// panic is only okay here because the application
		// can't start without config.
		panic(err)
	}
}

type Config struct {
	NetflixUsername string `split_words:"true" required:"true"`
	NetflixPassword string `split_words:"true" required:"true"`
}

// Get returns a copy of the config singleton
func Get() Config {
	return config
}
