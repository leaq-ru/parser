package config

import (
	"github.com/kelseyhightower/envconfig"
)

const ServiceName = "parser"

type c struct {
	Grpc     grpc
	STAN     stan
	NATS     nats
	MongoDB  mongodb
	Vk       vk
	Service  service
	LogLevel string `envconfig:"LOGLEVEL"`
}

type grpc struct {
	Port string `envconfig:"GRPC_PORT"`
}

type stan struct {
	ClusterID string `envconfig:"STAN_CLUSTERID"`
}

type nats struct {
	URL string `envconfig:"NATS_URL"`
}

type mongodb struct {
	URL string `envconfig:"MONGODB_URL"`
}

type vk struct {
	UserTokens string `envconfig:"VK_USERTOKENS"`
}

type service struct {
	Image    string `envconfig:"SERVICE_IMAGE"`
	City     string `envconfig:"SERVICE_CITY"`
	Category string `envconfig:"SERVICE_CATEGORY"`
}

var Env c

func init() {
	envconfig.MustProcess("", &Env)
}
