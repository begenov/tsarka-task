package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	POSTGRES PostgresConfig   `required:"true"`
	REDIS    RedisConfig      `required:"true"`
	HTTP     HTTPServerConfig `required:"true"`
}

type PostgresConfig struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     int    `envconfig:"POSTGRES_PORT" required:"true"`
	User     string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Database string `envconfig:"POSTGRES_DATABASE" required:"true"`
}

type HTTPServerConfig struct {
	Host           string        `envconfig:"HTTP_HOST" required:"true"`
	Port           string        `envconfig:"HTTP_PORT" required:"true"`
	WriteTimeout   time.Duration `envconfig:"HTTP_WRITETIMEOUT" required:"true"`
	ReadTimeout    time.Duration `envconfig:"HTTP_READTIMEOUT" required:"true"`
	MaxHeaderBytes int           `envconfig:"HTTP_MAXHEADERBYTES" required:"true"`
}

type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST" required:"true"`
	Port     string `envconfig:"REDIS_PORT" required:"true"`
	Password string `envconfig:"REDIS_PASSWORD"`
	DB       int    `envconfig:"REDIS_DB" required:"true"`
}

func NewConfig() Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

func GetPort() string {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func GetWriteTimeout() time.Duration {
	timeout := os.Getenv("HTTP_WRITETIMEOUT")
	if timeout == "" {
		return 10 * time.Second
	}
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		log.Fatalln(err)
	}
	return time.Duration(timeoutInt) * time.Second
}

func GetReadTimeout() time.Duration {
	timeout := os.Getenv("HTTP_READTIMEOUT")
	if timeout == "" {
		return 10 * time.Second
	}
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		log.Fatalln(err)
	}
	return time.Duration(timeoutInt) * time.Second
}

func GetMaxHeaderBytes() int {
	bytes := os.Getenv("HTTP_MAXHEADERBYTES")
	if bytes == "" {
		return 1
	}
	bytesInt, err := strconv.Atoi(bytes)
	if err != nil {
		log.Fatalln(err)
	}
	return bytesInt
}
