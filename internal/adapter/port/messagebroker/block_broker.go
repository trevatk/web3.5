package messagebroker

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	wallet "github.com/trevatk/go-wallet"
	"github.com/trevatk/web3.5/internal/adapter/setup"
	"github.com/trevatk/web3.5/internal/core/domain"

	pb "github.com/trevatk/go-pkg/proto/messaging/v1"
)

// BlockBroker client implementation
type BlockBroker struct {
	conn   *grpc.ClientConn
	wallet *wallet.Wallet
}

// New return new block broker client
func New(cfg *setup.Config) (*BlockBroker, error) {

	path := filepath.Clean(cfg.MessageBroker.WalletPath)
	_, err := os.Stat(path)
	if err != nil && err == os.ErrNotExist {
		// create new wallet and marshal to file
		w := wallet.New()
		err = w.MarshalToFile(path)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal wallet to file %v", err)
		}

	} else if err != nil {
		return nil, fmt.Errorf("failed to get wallet file info %v", err)
	}

	w, err := wallet.UnmarshalFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal wallet from file %v", err)
	}

	conn, err := grpc.Dial(
		cfg.MessageBroker.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to dial broker address %v", err)
	}

	return &BlockBroker{
		wallet: w,
		conn:   conn,
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
