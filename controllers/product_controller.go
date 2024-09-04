package controllers

import (
	"awesomeProject/config"
	"awesomeProject/dtos"
	"awesomeProject/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// Get all the products
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	var productResponse []dtos.ProductDTO

	for _, product := range products {
		productResponse = append(productResponse, dtos.ProductDTO{
			ID:    product.ID,
			Name:  product.Name,
			Price: product.Price,
		})
	}

	c.JSON(http.StatusOK, productResponse)
}

// Get a product by id
func GetProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
		return
	}

	productResponse := dtos.ProductDTO{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}

	c.JSON(http.StatusOK, productResponse)
}

// Create a product
func CreateProduct(c *gin.Context) {
	var newProduct models.Product
	var input map[string]interface{}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct.Name = input["name"].(string)
	priceStr := input["price"].(string)
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct.Price = price

	result := config.DB.Create(&newProduct)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, newProduct)
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
