package messagebroker

import (
	"github.com/trevatk/web3.5/internal/core/domain"
	"google.golang.org/grpc"
)

// BlockBroker
type BlockBroker struct {
	conn *grpc.ClientConn
}

func New() *BlockBroker {
	return &BlockBroker{}
}

// interface compliance
var _ domain.MessageBroker = (*BlockBroker)(nil)

// Close client connection
func (bb *BlockBroker) Close() error {
	return bb.conn.Close()
}
