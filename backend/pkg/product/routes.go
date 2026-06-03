package product

import (
	"api-go/config"

	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) chi.Router {
	productConfig := New(configuration)
	router := chi.NewRouter()

	
	router.Group(func(r chi.Router) {
		r.Get("/", productConfig.GetAllProducts)
		r.Get("/{id}", productConfig.GetProduct)
		r.Post("/", productConfig.CreateProduct)
		r.Put("/{id}", productConfig.UpdateProduct)
		r.Delete("/{id}", productConfig.DeleteProduct)
	})

	return router
}
