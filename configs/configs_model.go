package configs

// Config is the instance of configuration load from config file
var Config Configuration

// Configuration 模块配置
type Configuration struct {
	AppName         string
	Port            string
	Stage           string
	URL             string
	Limiter         *Limiter
	BasicAuth       *BasicAuth
	PostgresConf    *DatabaseConfig
	EmailNotifyList []string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

type BasicAuth struct {
	Account  string
	Password string
}

type Limiter struct {
	Limit int
	Burst int
}
