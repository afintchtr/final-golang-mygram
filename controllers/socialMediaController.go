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

func StoreSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"social_media": SocialMedia,
	})
}

func IndexSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	SocialMedias := []models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	err := db.Where("user_id = ?", userID).Order("id").Find(&SocialMedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"social_medias": SocialMedias,
	})
}

func ShowSocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("id"))

	err := db.First(&SocialMedia, socialMediaId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Social media not found",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"social_media": SocialMedia,
	})
}

func EditSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("id"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID
	SocialMedia.ID = uint(socialMediaId)

	err := db.Model(&SocialMedia).Where("id = ?", socialMediaId).Updates(models.SocialMedia{
		Name:       		SocialMedia.Name,
		SocialMediaUrl: SocialMedia.SocialMediaUrl,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"updated_social_media": SocialMedia,
	})
}

func DestroySocialMedia(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}

	socialMediaId, _ := strconv.Atoi(c.Param("id"))

	err := db.First(&SocialMedia, socialMediaId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Social media not found",
				"message": err.Error(),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Delete(&SocialMedia, socialMediaId).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Social media successfully deleted",
	})
}