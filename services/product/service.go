package product

import (
	"go-ecomrce/models"
	"go-ecomrce/requests"

	"github.com/jinzhu/gorm"
)

type ServiceWrapper interface {
	Create(product *models.Product)
	Delete(product *models.Product)
	Update(product *models.Product, updateProductRequest *requests.UpdateProductRequest)
}

type Service struct {
	DB *gorm.DB
}

func NewProductService(db *gorm.DB) *Service {
	return &Service{DB: db}
}
