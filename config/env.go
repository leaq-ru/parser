package config

import (
	"github.com/kelseyhightower/envconfig"
)

const ServiceName = "parser"

type c struct {
	Grpc     grpc
	Redis    redis
	STAN     stan
	NATS     nats
	MongoDB  mongodb
	Vk       vk
	Service  service
	S3       s3
	LogLevel string `envconfig:"LOGLEVEL"`
}

type grpc struct {
	Port string `envconfig:"GRPC_PORT"`
}

type redis struct {
	URL string `envconfig:"REDIS_URL"`
}

type stan struct {
	ClusterID               string `envconfig:"STAN_CLUSTERID"`
	SubjectReviewModeration string `envconfig:"STAN_SUBJECTREVIEWMODERATION"`
	SubjectURL              string `envconfig:"STAN_SUBJECTURL"`
	URLMaxInFlight          string `envconfig:"STAN_URLMAXINFLIGHT"`
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
	Image      string `envconfig:"SERVICE_IMAGE"`
	User       string `envconfig:"SERVICE_USER"`
	Classifier string `envconfig:"SERVICE_CLASSIFIER"`
	Wappalyzer string `envconfig:"SERVICE_WAPPALYZER"`
}

type s3 struct {
	DownloadBucketName string `envconfig:"S3_DOWNLOADBUCKETNAME"`
	Endpoint           string `envconfig:"S3_ENDPOINT"`
	AccessKeyID        string `envconfig:"S3_ACCESSKEYID"`
	SecretAccessKey    string `envconfig:"S3_SECRETACCESSKEY"`
	Secure             string `envconfig:"S3_SECURE"`
	Region             string `envconfig:"S3_REGION"`
}

var Env c

func init() {
	envconfig.MustProcess("", &Env)
}
