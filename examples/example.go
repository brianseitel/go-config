package main

import (
	"github.com/sanity-io/litter"
	"github.com/shipt/config/lib"
	"github.com/shipt/config/lib/options"
)

type ExampleConfig struct {
	Environment   string `default:"development"`
	Host          string `default:"localhost"`
	Redis         string `default:"localhost:6379"`
	PrivateFields string `mapstructure:"private_fields"`

	// mapstructure is used because viper is format-insensitive
	// and some formats don't allow underscores, so you have to
	// define it as a mapstructure instead of `json`
	ServiceName string `mapstructure:"service_name",default:"foobar"`
}

func main() {
	var cfg ExampleConfig
	c := lib.New("example", &cfg, yourMomOption, options.ValidateOption)
	litter.Dump(c)
}

var yourMomOption = func(c *lib.Configuration) {
	c.Set("ServiceName", "your mom")
}
