package productController

import (
	"net/http"

	"github.com/afiqbomboloni/crud-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	
	var products []models.Product

	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func Show(c * gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil {
		switch err{
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Content Not Found"})
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Network Error"})
		}


	} else {
		c.JSON(http.StatusOK, gin.H{
			"product": product,
		})
	}

	
}

func Create(c * gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			
		})
		return
	}

	

	models.DB.Create(&product)

	if product.Id == 0 {
		c.JSON(http.StatusBadRequest,gin.H{"message": "Value must unique"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"product": product,
	})

}

func Update(c * gin.Context) {

	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			
		})
		return
	}

	if models.DB.Model(&product).Where("id=?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Content Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Updated Successfully"})

}

func Delete(c * gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product

	if models.DB.Where("id=?", id).Delete(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Content Not Found"})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted",
	})
	
}