package config

type Options struct {
	Service  *Service  `yaml:"service"`
	Database *Database `yaml:"database"`
}

type Service struct {
	HttpPort           string `yaml:"http_port"`
	CurrentEnvironment string `yaml:"current_environment"`
}

type Database struct {
	Url            string `yaml:"url"`
	Host           string `yaml:"host"`
	Port           string `yaml:"port"`
	DefaultDb      string `yaml:"default_db"`
	UserName       string `yaml:"username"`
	Password       string `yaml:"password"`
	SSLMode        string `yaml:"ssl_mode"`
	ConnectTimeout int    `yaml:"connect_timeout"`
}
