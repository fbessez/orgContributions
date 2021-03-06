// Package config will load the ENV variables for constants we want to set dynamically.
package config

import (
	"github.com/kelseyhightower/envconfig"
)

var CONSTANTS Constants

type (
	Constants struct {
		Username string `yaml:"username" envconfig:"username" required:"true"`
		Password string `yaml:"password" envconfig:"password" required:"true"`
		OrgName  string `yaml:"orgname"  envconfig:"orgname"  required:"true"`
		Redis    Redis
	}

	Redis struct {
		Address string `yaml:"address" envconfig:"redis_address" required:"true"`
		Port    string `yaml:"port"    envconfig:"redis_port"    required:"true"`
	}
)

func init() {
	err := envconfig.Process("github", &CONSTANTS)
	if err != nil {
		panic(err)
	}
}
