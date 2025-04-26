package config

type Config struct {
	Logging Logging `yaml:"logging"`
	Http    Http    `yaml:"http"`
}

type Logging struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type Http struct {
	Address string `yaml:"address"`
}
