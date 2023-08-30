package backend

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type host struct {
	listenAddr string
}

func NewHost(listenAddr string) *host {
	return &host{
		listenAddr: listenAddr,
	}
}

func (h *host) Run() {
	router := chi.NewRouter().With(middleware.Logger)
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})
	fmt.Println("Sever starting at port: " + h.listenAddr)
	http.ListenAndServe(":"+h.listenAddr, router)
}
