package config

import (
	"log"
	"sync"
	"time"

	"github.com/Astemirdum/person-service/pkg/postgres"

	"github.com/Astemirdum/person-service/pkg/logger"

	"github.com/kelseyhightower/envconfig"
)

type HTTPServer struct {
	Host         string        `yaml:"host" envconfig:"HTTP_HOST"`
	Port         string        `yaml:"port" envconfig:"HTTP_PORT"`
	ReadTimeout  time.Duration `yaml:"readTimeout" envconfig:"HTTP_READ"`
	WriteTimeout time.Duration
}

type Config struct {
	Server   HTTPServer  `yaml:"server"`
	Database postgres.DB `yaml:"db"`
	Log      logger.Log  `yaml:"log"`
}

var (
	once sync.Once
	cfg  *Config
)

// NewConfig reads config from environment.
func NewConfig(ops ...Option) *Config {
	once.Do(func() {
		var config Config
		for _, op := range ops {
			op(&config)
		}
		err := envconfig.Process("", &config)
		if err != nil {
			log.Fatal("NewConfig ", err)
		}
		cfg = &config
		// printConfig(cfg)
	})

	return cfg
}

// func printConfig(cfg *Config) {
//	jscfg, _ := json.MarshalIndent(cfg, "", "	") //nolint:errcheck
//	fmt.Println(string(jscfg))
//}
