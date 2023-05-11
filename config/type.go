package config

type MysqlConfig struct {
	Root     string `mapstructure:"root"`
	Password string `mapstructure:"password"`
	Path     string `mapstructure:"path"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
}

type Config struct {
	MysqlConfig MysqlConfig `mapstructure:"db"`
}
