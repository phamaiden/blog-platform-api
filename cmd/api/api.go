package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type application struct {
	config config
}

type config struct {
	addr string
	_    string
}

func (app *application) mount() http.Handler {
	mux := mux.NewRouter()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	return mux
}

func (app *application) run(r http.Handler) error {
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: r,
	}

	log.Printf("server is running on %s", srv.Addr)

	return srv.ListenAndServe()
}
