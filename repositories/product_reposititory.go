package repositories

import (
	"go-ecomrce/models"

	"github.com/jinzhu/gorm"
)

type ProductRepositoryQ interface {
	GetProducts(products *[]models.Product)
	GetProduct(product *models.Product, id int)
}

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (productRepository *ProductRepository) GetProducts(products *[]models.Product) {
	productRepository.DB.Find(products)
}

func (productRepository *ProductRepository) GetProduct(product *models.Product, id int) {
	productRepository.DB.Where("id = ? ", id).Find(product)
}
