package product

import "go-ecomrce/models"

func (productService *Service) Create(product *models.Product) {
	productService.DB.Create(product)
}
