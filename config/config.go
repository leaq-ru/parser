package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Grpc        grpc
	MongoDB     mongodb
	Vk          vk
	Service     service
	S3          s3
	LogLevel    string `envconfig:"LOGLEVEL"`
	ServiceName string
}

type grpc struct {
	Port string `envconfig:"GRPC_PORT"`
}

type mongodb struct {
	URL string `envconfig:"MONGODB_URL"`
}

type vk struct {
	UserTokens string `envconfig:"VK_USERTOKENS"`
}

type service struct {
	Image      string `envconfig:"SERVICE_IMAGE"`
	City       string `envconfig:"SERVICE_CITY"`
	Category   string `envconfig:"SERVICE_CATEGORY"`
	Technology string `envconfig:"SERVICE_TECHNOLOGY"`
	User       string `envconfig:"SERVICE_USER"`
}

type s3 struct {
	DownloadBucketName string `envconfig:"S3_DOWNLOADBUCKETNAME"`
	Endpoint           string `envconfig:"S3_ENDPOINT"`
	AccessKeyID        string `envconfig:"S3_ACCESSKEYID"`
	SecretAccessKey    string `envconfig:"S3_SECRETACCESSKEY"`
	Secure             string `envconfig:"S3_SECURE"`
	Region             string `envconfig:"S3_REGION"`
}

func NewConfig() (Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, err
	}

	cfg.ServiceName = "parser"
	return cfg, nil
}
