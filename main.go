package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Category struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type Product struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
}

func main() {
	// Connect to the database
	var err error
	db, err = gorm.Open("mysql", "admin:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Auto migrate the database
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Product{})

	// Create a new Gin router
	router := gin.Default()

	// Define the API endpoints
	router.GET("/categories", getAllCategories)
	router.GET("/categories/:id", getCategoryById)
	router.POST("/categories", createCategory)
	router.PUT("/categories/:id", updateCategory)
	router.DELETE("/categories/:id", deleteCategory)

	router.GET("/products", getAllProducts)
	router.GET("/products/:id", getProductById)
	router.POST("/products", createProduct)
	router.PUT("/products/:id", updateProduct)
	router.DELETE("/products/:id", deleteProduct)

	// Start the server
	router.Run(":8080")
}

func getAllCategories(c *gin.Context) {
	var categories []Category
	db.Find(&categories)
	c.JSON(http.StatusOK, categories)

}

func getCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var category Category
	db.First(&category, id)
	if category.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, category)
}

func createCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	db.Create(&category)
	c.JSON(http.StatusCreated, category)

}

func updateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var category Category
	db.First(&category, id)
	if category.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := c.BindJSON(&category); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db.Save(&category)
	c.JSON(http.StatusOK, category)

}

func deleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var category Category
	db.First(&category, id)
	if category.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&category)
	c.Status(http.StatusNoContent)
}

func getAllProducts(c *gin.Context) {
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)

}

func getProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var product Product
	db.First(&product, id)
	if product.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, product)

}

func createProduct(c *gin.Context) {
	var product Product
	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	db.Create(&product)
	c.JSON(http.StatusCreated, product)

}
func updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var product Product
	db.First(&product, id)
	if product.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err := c.BindJSON(&product); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	db.Save(&product)
	c.JSON(http.StatusOK, product)

}

func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var product Product
	db.First(&product, id)
	if product.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	db.Delete(&product)
	c.Status(http.StatusNoContent)

}
