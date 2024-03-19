package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/multierr"

	"github.com/trevatk/go-pkg/logging"

	"github.com/trevatk/web3.5/internal/adapter/port/http/router"
	"github.com/trevatk/web3.5/internal/adapter/port/http/server"
	"github.com/trevatk/web3.5/internal/adapter/port/messagebroker"
	"github.com/trevatk/web3.5/internal/adapter/setup"
	"github.com/trevatk/web3.5/internal/core/application"
	"github.com/trevatk/web3.5/internal/core/domain"
)

func main() {
	fx.New(
		fx.Provide(context.TODO),
		fx.Provide(logging.NewLogger),
		fx.Provide(setup.New),
		fx.Provide(fx.Annotate(messagebroker.New, fx.As(new(domain.MessageBroker)))),
		fx.Provide(application.NewSurveyService),
		fx.Provide(fx.Annotate(router.New, fx.As(new(http.Handler)))),
		fx.Provide(server.New),
		fx.Invoke(registerHooks),
	).Run()
}

func registerHooks(lc fx.Lifecycle, s1 *http.Server) error {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := s1.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						log.Fatalf("failed to start http server %v", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {

				var result error

				err := s1.Shutdown(ctx)
				if err != nil {
					result = multierr.Append(result, fmt.Errorf("failed to shutdown http server %v", err))
				}

				return result
			},
		},
	)
	return nil
}
