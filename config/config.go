package config

var (
	defaultCfg = &Config{
		DBType:     "mysql",
		DBAddr:     "root:golang666@tcp(127.0.0.1:3306)/students",
		ServerAddr: ":12345",
	}
	AppCfg *Config
)

func InitConfig(path string) error {
	if path == "" {
		AppCfg = defaultCfg
	}
	return nil
}

type Config struct {
	DBType     string `yaml:"DBType"`
	DBAddr     string `yaml:"DBAddr"`
	ServerAddr string `yaml:"ServerAddr"`
}
