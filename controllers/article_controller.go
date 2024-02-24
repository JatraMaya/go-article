package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jatraMaya/go-library/models"
	"gorm.io/gorm"
)

// CreateArticle creates a new article and stores it inside a database
func CreateArticle(c *gin.Context) {
	var requestBodyArticle struct {
		Title   string `json:"Title" binding:"required"`
		Author  string `json:"Author" binding:"required"`
		Content string `json:"Content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestBodyArticle); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Failed to read request body",
			"Error":   err.Error(),
		})
		return
	}

	article := models.Article{Title: requestBodyArticle.Title, Author: requestBodyArticle.Author, Content: requestBodyArticle.Content}
	result := models.DB.Create(&article)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"article": article,
	})
}

// GetArticles retrieves a list of articless from the database
func GetArticles(c *gin.Context) {
	var articles []models.Article

	models.DB.Find(&articles)

	c.JSON(http.StatusOK, gin.H{
		"Articles": articles,
	})
}

// GetArticleByID retrieves an article by ID from the database
func GetArticleById(c *gin.Context) {
	id := c.Param("id")
	var article models.Article

	if err := models.DB.First(&article, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"Message": "Article not found",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"Article": article,
	})
}

// UpdateArticle update the referenced article based on ID
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var article models.Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return

	}

	if models.DB.Model(&article).Where("id = ?", id).Updates(&article).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Failed to update article",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Article is updated",
	})
}

// DeleteArticle delete article with corresponding ID
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	if models.DB.Delete(&models.Article{}, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"Message": "Failed to delete article",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Article succesfully deleted",
	})
}
