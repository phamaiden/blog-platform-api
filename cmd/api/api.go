package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phamaiden/blog-platform-api/internal/handlers"
)

type application struct {
	config config
}

type config struct {
	addr  string
	dbUrl string
}

func (app *application) mount(bh *handlers.BlogHandler) http.Handler {
	mux := mux.NewRouter()

	mux.Use(jsonContentTypeMiddleware)

	mux.HandleFunc("/", bh.GetAllBlogs)
	mux.HandleFunc("/posts", bh.PostBlog).Methods("POST")
	mux.HandleFunc("/posts/{id}", bh.PutBlog).Methods("PUT")
	mux.HandleFunc("/posts/{id}", bh.DeleteBlog).Methods("DELETE")
	mux.HandleFunc("/posts/{id}", bh.GetBlogById).Methods("GET")

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
