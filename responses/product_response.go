package responses

import (
	"go-ecomrce/models"
)

type ProductResponse struct {
	Title       string `json:"title" example:"Echo"`
	Description string `json:"desc" example:"Echo is nice and delicious"`
	Price       uint   `json:"price" example:"100000"`
	ImageUrl    string `json:"image_url"`
	Username    string `json:"username" example:"John Doe"`
	ID          uint   `json:"id" example:"1"`
}

func NewProductResponse(products []models.Product) *[]ProductResponse {
	productResponse := make([]ProductResponse, 0)

	for i := range products {
		productResponse = append(productResponse, ProductResponse{
			Title:       products[i].Title,
			Description: products[i].Description,
			Price:       products[i].Price,
			ImageUrl:    products[i].ImageUrl,
			Username:    products[i].User.Name,
			ID:          products[i].ID,
		})
	}

	return &productResponse
}
