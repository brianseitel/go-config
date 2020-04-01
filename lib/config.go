package lib

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"gopkg.in/mcuadros/go-defaults.v1"
)

type ConfigStruct interface{}

type Configuration struct {
	App  ConfigStruct
	Base BaseConfig
}

type Option func(c *Configuration)

func New(name string, cfg ConfigStruct, options ...Option) Configuration {
	initConfig(name)

	// load app-specific config
	err := viper.Unmarshal(cfg)

	if err != nil {
		log.Fatalf("unable to decode into app config struct, %v", err)
	}

	defaults.SetDefaults(cfg)

	var baseConfig BaseConfig
	err = viper.Unmarshal(&baseConfig)
	if err != nil {
		log.Fatalf("unable to decode into base config struct, %v", err)
	}

	c := &Configuration{App: cfg, Base: baseConfig}

	for _, opt := range options {
		opt(c)
	}

	return *c
}

// Set a field in the config struct with a particular value
func (c *Configuration) Set(field string, value interface{}) error {
	ok, err := setField(c.App, field, value)
	if !ok && err != nil {
		ok, err = setField(c.Base, field, value)
		if !ok || err != nil {
			return err
		}
	}
	return nil
}

// Has attempts to retrieve a field in the config struct
func (c *Configuration) Has(field string) bool {
	out := getField(c.App, field)
	if out == nil {
		out = getField(c.Base, field)
	}

	return out != nil
}

// Get retrieves a field in the config struct
func (c *Configuration) Get(field string) interface{} {
	out := getField(c.App, field)
	if out == nil {
		out = getField(c.Base, field)
	}

	return out
}

// initConfig reads in config file and ENV variables if set.
func initConfig(name string) {
	if len(viper.AllKeys()) == 0 {
		viper.SetConfigType("json")

		if name != "" {
			// Use config file from the flag.
			viper.SetConfigFile(fmt.Sprintf(".%s", name))
		} else {
			// Find home directory.
			home, err := homedir.Dir()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Search config in home directory with name ".prism" (without extension).
			viper.AddConfigPath(home)
			viper.SetConfigName(".config")
		}

		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.WatchConfig()
		viper.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else {
			panic(err)
		}
	}
}
