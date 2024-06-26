package controllers

import (
	"net/http"

	"app.myriadflow.com/db"
	"app.myriadflow.com/models"
	"github.com/gin-gonic/gin"
)

func CreateBrand(c *gin.Context) {
	var brand models.Brand
	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brand)
}

func GetBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand
	if err := db.DB.First(&brand, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	c.JSON(http.StatusOK, brand)
}

// get all brands api
func GetAllBrands(c *gin.Context) {
	var brands []models.Brand
	if err := db.DB.Find(&brands).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, brands)
}

func UpdateBrand(c *gin.Context) {
	id := c.Param("id")
	var brand models.Brand
	if err := db.DB.First(&brand, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Brand not found"})
		return
	}

	if err := c.ShouldBindJSON(&brand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&brand).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, brand)
}

func DeleteBrand(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Brand{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Brand deleted successfully"})
}
