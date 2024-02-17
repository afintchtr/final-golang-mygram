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

// StorePhoto godoc
// @Summary Create a new photo
// @Description Create a new photo with title, caption and url
// @Tags photo
// @Accept json
// @Produces json
// @Param models.Photo body models.Photo true "create photo"
// @Success 200 {object} models.Photo
// @Router /photos/ [post]
func StorePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, Photo)
}

// IndexPhoto godoc
// @Summary Get all photo details
// @Description Get details of all photos
// @Tags photo
// @Accept json
// @Produces json
// @Success 200 {object} models.Photo
// @Router /photos/ [get]
func IndexPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photos := []models.Photo{}

	userID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", userID).Preload("Comments").Order("id").Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Photos)
}

// ShowPhoto godoc
// @Summary Get photo details
// @Description Get details sof one photo by its id
// @Tags photo
// @Accept json
// @Produces json
// @Param id path int true "ID of the photo to be shown"
// @Success 200 {object} models.Photo
func ShowPhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("id"))

	err := db.Preload("Comments").First(&Photo, photoId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Requested photo not found",
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
	c.JSON(http.StatusOK, Photo)
}

// EditPhoto godoc
// @Summary Update an existing photo
// @Description Update an existing photo with new title, caption, and url
// @Tags photo
// @Accept json
// @Produces json
// @Param id path int true "ID of the photo to be updated"
// @Success 200 {object} models.Photo
// @Router /photos/{id} [patch]
func EditPhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("id"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err := db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{
		Title:    Photo.Title,
		Caption: 	Photo.Caption,
		PhotoUrl: Photo.PhotoUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": Photo.ID,
		"user_id": Photo.UserID,
		"updated_title": Photo.Title,
		"updated_caption": Photo.Caption,
		"updated_photo_url": Photo.PhotoUrl,
		"updated_at": Photo.UpdatedAt,
	})
}

// DestroyPhoto godoc
// @Summary Delete an existing photo
// @Description Just delete an existing photo
// @Tags photo
// @Accept json
// @Produces json
// @Param id path int true "ID of the photo to be deleted"
// @Success 204 "No Content"
// @Router /photos/{id} [delete]
func DestroyPhoto(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}

	photoId, _ := strconv.Atoi(c.Param("id"))

	err := db.First(&Photo, photoId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error":   "Photo not found",
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

	err = db.Model(&models.Comment{}).Where("photo_id = ?", photoId).Delete(&models.Comment{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Can't delete comments",
			"error_detail": err,
		})
		return
	}

	err = db.Delete(&Photo, photoId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Photo successfully deleted",
	})
}
