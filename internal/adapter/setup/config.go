package setup

// Config
type Config struct {
	Server        *Server
	MessageBroker *MessageBroker
}

// New
func New() *Config {
	return &Config{
		Server:        &Server{},
		MessageBroker: &MessageBroker{},
	}
}
