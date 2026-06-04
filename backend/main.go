package main

import (
	"api-go/config"
	_ "api-go/docs"
	"api-go/pkg/product"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/swaggo/swag"
)

// @title Smart Coffee API
// @version 1.0
// @description API
// @BasePath /api/v1

// corsMiddleware handles CORS headers for frontend communication
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Routes(configuration *config.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(corsMiddleware)

	router.Get("/", healthHandler)
	router.Get("/health", healthHandler)

	if configuration == nil {
		return router
	}

	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:"+os.Getenv("PORT")+"/swagger/swagger.json"),
	))

	router.Get("/swagger/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		doc, _ := swag.ReadDoc()
		w.Write([]byte(doc))
	})

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/products", product.Routes(configuration))
	})

	return router
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func hasDatabaseConfig() bool {
	requiredVariables := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT", "DB_SSLMODE"}
	for _, variable := range requiredVariables {
		if os.Getenv(variable) == "" {
			return false
		}
	}

	return true
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	var configuration *config.Config
	if hasDatabaseConfig() {
		var err error
		configuration, err = config.New()
		if err != nil {
			log.Println("Failed to initialize database configuration, product routes disabled:", err)
			configuration = nil
		}
	} else {
		log.Println("Database environment is incomplete, product routes disabled")
	}

	router := Routes(configuration)
	log.Println("Server running on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
