package config

type Config struct {
	HttpAddr string
}

func ReadConfig() *Config {
	return &Config{
		HttpAddr: ":8080",
	}
}
