package main

import (
	"fmt"
	"khaibaoyte/apis"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Homepage"))
	})

	r.Route("/list", func(r chi.Router) {
		// Get List
		r.With(paginate).Get("/person", apis.ListPerson)
		r.Get("/province", apis.ListProvince)
		r.Get("/town", apis.ListTown)
		r.Get("/village", apis.ListVillage)
		r.Get("/sex", apis.ListSex)
		r.Get("/nationality", apis.ListNationality)
		r.Get("/question", apis.ListQuestion)
		r.Get("/health_declaration", apis.ListHealthDeclaration)
	})

	r.Route("/person", func(r chi.Router) {
		r.Post("/", apis.AddPerson)
		r.Get("/{id}", apis.GetPerson)
		r.Put("/{id}", apis.UpdatePerson)
	})

	r.Route("/health_declaration", func(r chi.Router) {
		r.Post("/", apis.AddHealthDeclaration)
		r.Get("/{id}", apis.GetHealthDeclaration)
		r.Put("/{id}", apis.UpdateHealthDeclaration)

		r.Post("/question", apis.AddQuestion)
		r.Get("/question/{id}", apis.GetQuestion)
		r.Put("/question/{id}", apis.UpdateQuestion)
	})

	// Mount the admin sub-router
	// r.Mount("/admin", adminRouter())

	// Listen server at port 8080
	fmt.Println("Listen server at port 8080")
	http.ListenAndServe(":8080", r)
}

func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}
