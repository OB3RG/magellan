package config

// Config ...
type Config struct {
	Environment string `arg:"env:ENVIRONMENT"`
	ServerConfig
	DBConfig
}

// ServerConfig ...
type ServerConfig struct {
	Port string `arg:"env:SERVER_PORT"`
}

// DBConfig ...
type DBConfig struct {
	DBUser string `arg:"env:DB_USER"`
	DBPwd  string `arg:"env:DB_PASSWORD"`
	DBHost string `arg:"env:DB_HOST"`
	DBPort string `arg:"env:DB_PORT"`
	DBName string `arg:"env:DB_NAME"`
}

// DefaultCfg ...
func DefaultCfg() *Config {
	return &Config{
		Environment: "dev",
		ServerConfig: ServerConfig{
			Port: "5000",
		},
		DBConfig: DBConfig{
			DBUser: "postgres",
			DBPwd:  "postgres",
			DBHost: "localhost",
			DBPort: "5432",
			DBName: "boilerplate",
		},
	}
}
