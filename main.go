package main

import (
	"fmt"
	"khaibaoyte/apis"
	_ "khaibaoyte/docs"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Health Declarations API
// @version 1.0
// @description This is a service for health declaration
// @host localhost:8080
// @BasePath /
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Homepage"))
	})

	r.Route("/api/v1/", func(r chi.Router) {
		// Provinces
		r.Get("/provinces", apis.GetProvinces)

		// Towns
		r.Get("/towns", apis.GetTowns)

		// Villages
		r.Get("/villages", apis.GetVillages)

		// Genders
		r.Get("/genders", apis.GetGenders)

		// Nationalitys
		r.Get("/nationalitys", apis.GetNationalitys)

		// Person Info
		r.With(paginate).Get("/persons", apis.GetPersons)
		r.Post("/persons", apis.AddPerson)
		r.Get("/persons/{id}", apis.GetPerson)
		r.Put("/persons/{id}", apis.UpdatePerson)

		// Health Declarations
		r.Get("/health-declarations", apis.GetHealthDeclarations)
		r.Post("/health-declarations", apis.AddHealthDeclaration)
		r.Get("/health-declarations/{id}", apis.GetHealthDeclaration)
		r.Put("/health-declarations/{id}", apis.UpdateHealthDeclaration)

		// Questions
		r.Get("/questions", apis.GetQuestions)
		r.Post("/questions", apis.AddQuestion)
		r.Get("/questions/{id}", apis.GetQuestion)
		r.Put("/questions/{id}", apis.UpdateQuestion)
	})

	// Mount the admin sub-router
	// r.Mount("/admin", adminRouter())

	// Port 8080
	r.Mount("/swagger", httpSwagger.WrapHandler)
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
