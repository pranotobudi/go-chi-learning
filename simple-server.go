package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	// Creating a New Router
	apiRouter := chi.NewRouter()
	apiRouter.Get("/articles/{date}-{slug}", getArticle)

	// Mounting the new Sub Router on the main router
	r.Mount("/api", apiRouter)

	http.ListenAndServe(":8080", r)
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	dateParam := chi.URLParam(r, "date")
	slugParam := chi.URLParam(r, "slug")
	// article, err := database.GetArticle(date, slug)
	article := "dateParam: " + dateParam + " slugParam: " + slugParam

	// if err != nil {
	//   w.WriteHeader(422)
	//   w.Write([]byte(fmt.Sprintf("error fetching article %s-%s: %v", dateParam, slugParam, err)))
	//   return
	// }

	// if article == nil {
	//   w.WriteHeader(404)
	//   w.Write([]byte("article not found"))
	//   return
	// }
	// w.Write([]byte(article.Text()))
	w.Write([]byte(article))
}
