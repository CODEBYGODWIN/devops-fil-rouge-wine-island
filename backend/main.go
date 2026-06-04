package main

import (
	"api-go/config"
	"api-go/pkg/product"
	"log"
	"net/http"
	"os"
		httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

// @title Smart Coffee API
// @version 1.0
// @description API
// @BasePath /api/v1


func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:5555/swagger/swagger.json"),
	))

	// Serve swagger.json file
	router.Get("/swagger/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, "./docs/swagger.json")
	})

	router.Route("/api/v1", func(r chi.Router) {
		
		r.Mount("/products", product.Routes(configuration))
	})

	return router
}

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Fatal("Failed to initialize configuration:", err)
	}
	godotenv.Load()
	router := Routes(configuration)
	log.Println("Server running on http://localhost:" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
