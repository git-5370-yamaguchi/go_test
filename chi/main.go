package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DELETE request received")
}

func main() {
	r := chi.NewRouter()

	// // A good base middleware stack
	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	// r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)

	// // Set a timeout value on the request context (ctx), that will signal
	// // through ctx.Done() that the request has timed out and further
	// // processing should be stopped.
	// r.Use(middleware.Timeout(60 * time.Second))

	// r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Delete("/delete", handleDelete)

	http.ListenAndServe(":1324", r)
}
