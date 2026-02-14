package server

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

func NewServerMux(routers []Router) *chi.Mux {
	mux := chi.NewRouter()

	for _, router := range routers {
		router.Register(mux)
	}

	return mux
}

type Router interface {
	Register(r chi.Router)
}

func AsRouter(router any) any {
	return fx.Annotate(
		router,
		fx.As(new(Router)),
		fx.ResultTags(`group:"routers"`),
	)
}
