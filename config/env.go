package config

import (
	"github.com/kelseyhightower/envconfig"
)

type c struct {
	Grpc     grpc
	Mongo    mongo
	Vk       vk
	LogLevel string `envconfig:"LOGLEVEL"`
}

type grpc struct {
	Port string `envconfig:"GRPC_PORT"`
}

type mongo struct {
	URI string `envconfig:"MONGO_URI"`
}

type vk struct {
	GroupTokens string `envconfig:"VK_GROUPTOKENS"`
}

var Env c

func init() {
	envconfig.MustProcess("", &Env)
}
