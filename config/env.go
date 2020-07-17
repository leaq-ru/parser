package config

import (
	"github.com/kelseyhightower/envconfig"
)

const ServiceName = "parser"

type c struct {
	Grpc     grpc
	STAN     stan
	Mongo    mongo
	Vk       vk
	LogLevel string `envconfig:"LOGLEVEL"`
}

type grpc struct {
	Port string `envconfig:"GRPC_PORT"`
}

type stan struct {
	URL       string `envconfig:"STAN_URL"`
	ClusterID string `envconfig:"STAN_CLUSTERID"`
}

type mongo struct {
	URL string `envconfig:"MONGO_URL"`
}

type vk struct {
	UserTokens string `envconfig:"VK_USERTOKENS"`
}

var Env c

func init() {
	envconfig.MustProcess("", &Env)
}
