package config

type Config struct {
	Logging Logging `yaml:"logging"`
}

type Logging struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}
