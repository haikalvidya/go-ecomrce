package handlers

import (
	"go-ecomrce/models"
	"go-ecomrce/repositories"
	"go-ecomrce/requests"
	"go-ecomrce/responses"
	s "go-ecomrce/server"
	productservice "go-ecomrce/services/product"
	"go-ecomrce/services/token"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type ProductHandlers struct {
	server *s.Server
}

func NewProductHandlers(server *s.Server) *ProductHandlers {
	return &ProductHandlers{server: server}
}

// CreateProduct godoc
// @Summary Create product
// @Description Create product
// @ID products-create
// @Tags Products Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateProductRequest true "Product title, description, price, and image"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Security ApiKeyAuth
// @Router /api/v1/products [post]
func (p *ProductHandlers) CreateProduct(c echo.Context) error {
	createProductRequest := new(requests.CreateProductRequest)

	if err := c.Bind(createProductRequest); err != nil {
		return err
	}

	if err := createProductRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*token.JwtCustomClaims)
	id := claims.ID

	resultImageUrl := ""
	product := models.Product{
		Title:       createProductRequest.Title,
		Description: createProductRequest.Description,
		Price:       createProductRequest.Price,
		ImageUrl:    resultImageUrl,
		UserID:      id,
	}
	productService := productservice.NewProductService(p.server.DB)
	productService.Create(&product)

	return responses.MessageResponse(c, http.StatusCreated, "Product successfully created")
}

// UploadImageProduct godoc
// @Summary Upload Image product
// @Description Upload Image product
// @ID products-upload-image
// @Tags Products Actions
// @Accept multipart/form-data
// @Produce json
// @Param id path int true "Product ID"
// @Param file formData file fale  "this is a test file"
// @Success 201 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Security ApiKeyAuth
// @Router /api/v1/products/{id}/image [post]
func (p *ProductHandlers) UploadImageProduct(c echo.Context) error {
	updateProductRequest := new(requests.UpdateProductRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	src, err := file.Open()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	t := strconv.FormatInt(time.Now().Unix(), 10)
	dst, err := os.Create("uploads/" + t + "-" + file.Filename)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	product := models.Product{}

	productRepository := repositories.NewProductRepository(p.server.DB)
	productRepository.GetProduct(&product, id)

	if product.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Product not found")
	}

	if product.ImageUrl != "" {
		os.Remove(product.ImageUrl)
	}

	ImageUrl := dst.Name()
	productService := productservice.NewProductService(p.server.DB)
	productService.Update(&product, updateProductRequest, ImageUrl)

	return responses.MessageResponse(c, http.StatusOK, "Product image uploaded successfully")
}

// DeleteProduct godoc
// @Summary Delete product
// @Description Delete product
// @ID products-delete
// @Tags Products Actions
// @Param id path int true "Product ID"
// @Success 204 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /api/v1/products/{id} [delete]
func (p *ProductHandlers) DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product := models.Product{}

	productRepository := repositories.NewProductRepository(p.server.DB)
	productRepository.GetProduct(&product, id)

	if product.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Product not found")
	}

	if product.ImageUrl != "" {
		os.Remove(product.ImageUrl)
	}

	productService := productservice.NewProductService(p.server.DB)
	productService.Delete(&product)

	return responses.MessageResponse(c, http.StatusNoContent, "Product deleted successfully")
}

// GetProducts godoc
// @Summary Get products
// @Description Get the list of all products
// @ID products-get
// @Tags Products Actions
// @Produce json
// @Success 200 {array} responses.ProductResponse
// @Security ApiKeyAuth
// @Router /api/v1/products [get]
func (p *ProductHandlers) GetProducts(c echo.Context) error {
	var products []models.Product

	productRepository := repositories.NewProductRepository(p.server.DB)
	productRepository.GetProducts(&products)

	for i := 0; i < len(products); i++ {
		p.server.DB.Model(&products[i]).Related(&products[i].User)
	}

	response := responses.NewProductResponse(products)
	return responses.Response(c, http.StatusOK, response)
}

// GetProduct godoc
// @Summary Get product
// @Description Get the products by id
// @ID products-get-one
// @Tags Products Actions
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {array} responses.ProductResponse
// @Security ApiKeyAuth
// @Router /api/v1/products/{id} [get]
func (p *ProductHandlers) GetAProduct(c echo.Context) error {
	var product models.Product
	id, _ := strconv.Atoi(c.Param("id"))

	productRepository := repositories.NewProductRepository(p.server.DB)
	productRepository.GetProduct(&product, id)

	return responses.Response(c, http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary Update product
// @Description Update product
// @ID products-update
// @Tags Products Actions
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param params body requests.UpdateProductRequest true "Product title, description, price, and image"
// @Success 200 {object} responses.Data
// @Failure 400 {object} responses.Error
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /api/v1/products/{id} [put]
func (p *ProductHandlers) UpdateProduct(c echo.Context) error {
	updateProductRequest := new(requests.UpdateProductRequest)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(updateProductRequest); err != nil {
		return err
	}

	if err := updateProductRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	product := models.Product{}

	productRepository := repositories.NewProductRepository(p.server.DB)
	productRepository.GetProduct(&product, id)

	if product.ID == 0 {
		return responses.ErrorResponse(c, http.StatusNotFound, "Product not found")
	}

	var ImageUrl string
	if product.ImageUrl == "" {
		ImageUrl = ""
	} else {
		ImageUrl = product.ImageUrl
	}
	productService := productservice.NewProductService(p.server.DB)
	productService.Update(&product, updateProductRequest, ImageUrl)

	return responses.MessageResponse(c, http.StatusOK, "Product successfully updated")
}
