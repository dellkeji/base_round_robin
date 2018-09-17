package config

// Server :
type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// Init : init default conf
func (c *Server) Init() error {
	c.Host = "0.0.0.0"
	c.Port = 8080
	return nil
}
