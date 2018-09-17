package config

// RedisConf : redis configure
type RedisConf struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	Password       string `yaml:"password"`
	DefaultDB      int    `yaml:"default_db"`
	MaxPoolSize    int    `yaml:"max_pool_size"`
	MaxConnTimeout int    `yaml:"max_conn_timeout"`
	IdleTimeout    int    `yaml:"idle_timeout"`
	ReadTimeout    int    `yaml:"read_timeout"`
	WriteTimeout   int    `yaml:"write_timeout"`
}

// Init : for dev env
func (client *RedisConf) Init() error {
	client.Host = "127.0.0.1"
	client.Port = 6379
	client.Password = ""
	client.DefaultDB = 0

	client.MaxPoolSize = 100
	client.MaxConnTimeout = 5
	client.IdleTimeout = 600
	client.ReadTimeout = 10
	client.WriteTimeout = 10
	return nil
}
