package setup

// Config
type Config struct {
	MessageBroker *MessageBroker
}

// New
func New() *Config {
	return &Config{
		MessageBroker: &MessageBroker{},
	}
}
