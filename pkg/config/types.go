package config

type Logger struct {
	Use         string `yaml:"use"`
	Environment string `yaml:"environment"`
	Loglevel    string `yaml:"loglevel"`
	Filename    string `yaml:"filename"`
}
type Cache struct {
	Use   string `yaml:"use"`
	Redis Redis  `yaml:"redis"`
}
type Redis struct {
	URL  string `yaml:"url"`
	Pass string `yaml:"password"`
	DB   int    `yaml:"db"`
}
type Server struct {
	Host string `yaml:"host"`
}
