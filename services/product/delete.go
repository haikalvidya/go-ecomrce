package product

import "go-ecomrce/models"

func (productService *Service) Delete(product *models.Product) {
	productService.DB.Delete(product)
}
