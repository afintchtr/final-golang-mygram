package controllers

import (
	"errors"
	"final-golang-mygram/database"
	"final-golang-mygram/helpers"
	"final-golang-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StoreComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	Photo := models.Photo{}

	userID := uint(userData["id"].(float64))

	photoId, _ := (strconv.Atoi(c.Param("photoId")))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.PhotoID = uint(photoId)
	Comment.UserID = userID

	err := db.Select("id").First(&Photo, photoId).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Affiliated photo does not exist",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Comment)
}

func IndexComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Comments := []models.Comment{}

	userID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", userID).Order("id").Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comments)
}

func ShowComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.First(&Comment, commentId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Requested comment not found",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Comment)
}

func EditComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err := db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{
		Message:    Comment.Message,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": Comment.ID,
		"user_id": Comment.UserID,
		"updated_message": Comment.Message,
		"updated_at": Comment.UpdatedAt,
	})
}

func DestroyComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	commentId, _ := strconv.Atoi(c.Param("commentId"))

	err := db.First(&Comment, commentId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Comment not found",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Delete(&Comment, commentId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Comment successfully deleted",
	})
}