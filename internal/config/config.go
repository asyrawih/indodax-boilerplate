package config

type Database struct {
	Dsn string `yaml:"Dsn"`
}

type Config struct {
	Name     string   `yaml:"Name"`
	Host     string   `yaml:"Host"`
	Port     string   `yaml:"Port"`
	Database Database `yaml:"Database"`
}
