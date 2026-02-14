package echo

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type EchoModule struct{}

func NewModule() EchoModule {
	return EchoModule{}
}

func (e EchoModule) Register(r chi.Router) {
	r.Get("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("echo"))
	})
}
