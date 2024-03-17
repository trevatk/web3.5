package setup

// MessageBroker
type MessageBroker struct {
	WalletPath string `env:"MSG_BROKER_WALLET_PATH"`
	Address    string `env:"MSG_BROKER_ADDR"`
}
