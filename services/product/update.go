package product

import (
	"go-ecomrce/models"
	"go-ecomrce/requests"
)

func (productService *Service) Update(product *models.Product, updateProductRequest *requests.UpdateProductRequest, imageUrl string) {
	product.Description = updateProductRequest.Description
	product.Price = updateProductRequest.Price
	product.ImageUrl = imageUrl
	product.Title = updateProductRequest.Title
	productService.DB.Save(product)
}
