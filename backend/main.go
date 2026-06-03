package main

import (
	"api-go/config"
	"api-go/pkg/product"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)


func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	

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
