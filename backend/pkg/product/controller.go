package product

import (
	"api-go/config"
	"api-go/database/dbmodel"
	model "api-go/pkg/models"
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5"
)

type ProductConfig struct {
	*config.Config
}

func New(configuration *config.Config) *ProductConfig {
	return &ProductConfig{configuration}
}


func (c *ProductConfig) CreateProduct(w http.ResponseWriter, r *http.Request) {
	req := &model.ProductRequest{}
	if err := render.Bind(r, req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	product := dbmodel.Product{
		Name:         req.Name,
		Description:   req.Description,
		ImageUrl:         req.ImageUrl,
		Price:         req.Price,
		
	}

	createdProduct, err := c.ProductEntryRepository.Create(&product)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Error creating product"})
		return
	}

	res := &model.ProductResponse{
		Id:            createdProduct.ID,
		Name:         createdProduct.Name,
		Description:   createdProduct.Description,
		ImageUrl:         createdProduct.ImageUrl,
		Price:         createdProduct.Price,
	
	}
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
}

func (c *ProductConfig) GetProduct(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid id parameter"})
		return
	}

	product, err := c.ProductEntryRepository.FindById(uint(id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "Product not found"})
		return
	}

	res := &model.ProductResponse{
		Id:            product.ID,
		Name:         product.Name,
		Description:   product.Description,
		ImageUrl:         product.ImageUrl,
		Price:         product.Price,
	
	}
	render.JSON(w, r, res)
}


func (c *ProductConfig) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.ProductEntryRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Error fetching products"})
		return
	}

	res := make([]*model.ProductResponse, len(products))
	for i, p := range products {
		res[i] = &model.ProductResponse{
			Id:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			ImageUrl:    p.ImageUrl,
			Price:       p.Price,
		}
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}


func (c *ProductConfig) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid id parameter"})
		return
	}

	product, err := c.ProductEntryRepository.FindById(uint(id))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{"error": "Product not found"})
		return
	}

	req := &model.ProductUpdateRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	if req.Name != nil {
		product.Name = *req.Name
	}
	if req.Description != nil {
		product.Description = *req.Description
	}
	if req.ImageUrl != nil {
		product.ImageUrl = *req.ImageUrl
	}
	if req.Price != nil {
		product.Price = *req.Price
	}

	product, err = c.ProductEntryRepository.Update(product)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Error updating product"})
		return
	}

	res := &model.ProductResponse{
		Id:            product.ID,
		Name:         product.Name,
		Description:   product.Description,
		ImageUrl:         product.ImageUrl,
		Price:         product.Price,
		
	}
	render.JSON(w, r, res)
}


func (c *ProductConfig) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "Invalid id parameter"})
		return
	}

	product := &dbmodel.Product{}
	err = c.ProductEntryRepository.Delete(uint(id), product)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{"error": "Error deleting product"})
		return
	}

	render.Status(r, http.StatusNoContent)
}
