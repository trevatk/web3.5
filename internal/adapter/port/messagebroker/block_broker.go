package messagebroker

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/trevatk/wallet"
	"github.com/trevatk/web3.5/internal/adapter/setup"
	"github.com/trevatk/web3.5/internal/core/domain"

	pb "github.com/trevatk/go-pkg/proto/messaging/v1"
)

// BlockBroker client implementation
type BlockBroker struct {
	conn *grpc.ClientConn
	addr string
}

// New return new block broker client
func New(cfg *setup.Config) (*BlockBroker, error) {
	// TODO:
	// persist wallet
	w := wallet.NewWallet()
	addr, err := w.Address()
	if err != nil {
		return nil, fmt.Errorf("failed to get wallet address %v", err)
	}

	conn, err := grpc.Dial(
		cfg.MessageBroker.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial broker address %v", err)
	}

	return &BlockBroker{
		addr: addr,
		conn: conn,
	}, nil
}

// interface compliance
var _ domain.MessageBroker = (*BlockBroker)(nil)

// Publish message
func (bb *BlockBroker) Publish(ctx context.Context, topic string, payload []byte) error {

	bb.conn.Connect()

	client := pb.NewMessagingServiceV1Client(bb.conn)

	timeout, cancel := context.WithTimeout(ctx, time.Millisecond*250)
	defer cancel()

	_, err := client.Publish(timeout, &pb.Envelope{
		Topic:   topic,
		Payload: payload,
	})
	if err != nil {
		return fmt.Errorf("failed to publish message %v", err)
	}

	return nil
}

// Close client connection
func (bb *BlockBroker) Close() error {
	return bb.conn.Close()
}
