package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/andresmeireles/go-template/internal/core/server"
	"github.com/andresmeireles/go-template/internal/feature/echo"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			NewHttpServer,
			zap.NewExample,
			fx.Annotate(
				server.NewServerMux,
				fx.ParamTags(`group:"routers"`),
			),
			server.AsRouter(echo.NewModule),
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

func NewHttpServer(lc fx.Lifecycle, mux *chi.Mux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: fmt.Sprintf(":%s", os.Getenv("APP_PORT")), Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}

			log.Info("Start Http Server", zap.String("Addr", srv.Addr))
			go srv.Serve(ln)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("o server morreu! ahhhh")

			return nil
		},
	})

	return srv
}
