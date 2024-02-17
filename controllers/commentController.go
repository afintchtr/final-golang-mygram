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

// StoreComment godoc
// @Summary Create a new comment
// @Description Create a new comment with message
// @Tags comment
// @Accept json
// @Produces json
// @Param models.Comment body models.Comment true "create comment"
// @Success 200 {object} models.Comment
// @Router /comments/{photoId} [post]
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

// IndexComment godoc
// @Summary Get all comments details
// @Description Get details of all comments
// @Tags comment
// @Accept json
// @Produces json
// @Success 200 {object} models.Comment
// @Router /comments/ [get]
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

// ShowComment godoc
// @Summary Get comment details
// @Description Get details of one comment by its id
// @Tags comment
// @Accept json
// @Produces json
// @Param id path int true "ID of the comments to be shown"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [get]
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

// EditComment godoc
// @Summary Update an existing comment
// @Description Update an existing comment with new message
// @Tags comment
// @Accept json
// @Produces json
// @Param id path int true "ID of the comment to be updated"
// @Success 200 {object} models.Comment
// @Router /comments/{commentId} [patch]
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

// DestroyComment godoc
// @Summary Delete an existing comment
// @Description Just delete an existing comment
// @Tags comment
// @Accept json
// @Produces json
// @Param id path int true "ID of the comment to be deleted"
// @Success 204 "No Content"
// @Router /comments/{commentId} [delete]
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