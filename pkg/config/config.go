package config

type Config struct {
	BaseUrl string `yaml:"base_url"`
}

func ParseConfig() Config {
	return Config{
		BaseUrl: "https://shrt.torfstack.com",
	}
}
