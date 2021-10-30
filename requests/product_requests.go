package requests

import validation "github.com/go-ozzo/ozzo-validation"

type BasicProduct struct {
	Title       string `json:"title,omitempty" validate:"required" example:"Echo"`
	Description string `json:"content" example:"Echo is nice!"`
	Price       uint   `json:"price,omitempty" validate:"required" example:"40000"`
}

func (bp BasicProduct) Validate() error {
	return validation.ValidateStruct(&bp,
		validation.Field(&bp.Title, validation.Required),
		validation.Field(&bp.Description, validation.Required),
	)
}

type CreateProductRequest struct {
	BasicProduct
}

type UpdateProductRequest struct {
	BasicProduct
}
