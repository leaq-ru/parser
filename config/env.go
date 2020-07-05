package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/nnqq/scr-parser/logger"
)

type c struct {
	Grpc  grpc
	Mongo mongo
	Vk    vk
}

type grpc struct {
	Port string `envconfig:"GRPC_PORT"`
}

type mongo struct {
	URI string `envconfig:"MONGO_URI"`
}

type vk struct {
	Tokens string `envconfig:"VK_TOKENS"`
}

var Env c

func init() {
	err := envconfig.Process("", &Env)
	logger.Must(err)
}
