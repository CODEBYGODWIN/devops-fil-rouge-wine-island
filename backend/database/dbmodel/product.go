package dbmodel

import "gorm.io/gorm"

type Product struct {
    gorm.Model
    Name string `gorm:"not null;"`
	Description string `gorm:"not null;"`
	ImageUrl string 
	Price float64 `gorm:"not null;"`
	
}

type ProductRepository interface {
	Create(product *Product) (*Product, error)
	FindAll() ([]*Product, error)
	FindById(id uint) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(id uint, product *Product) error
}

type productRepository struct {
	db *gorm.DB
}	

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll() ([]*Product, error) {
	var products []*Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) FindById(id uint) (*Product, error) {
	var product Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *Product) (*Product, error) {
	if err := r.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Update(product *Product) (*Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *productRepository) Delete(id uint, product *Product) error {
	return r.db.Delete(product, id).Error
}
