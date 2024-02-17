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

// StoreSocialMedia godoc
// @Summary Create a new social media
// @Description Create a new social media with name and url
// @Tags social-media
// @Accept json
// @Produces json
// @Param models.SocialMedia body models.SocialMedia true "create social media"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/ [post]
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
	c.JSON(http.StatusCreated, SocialMedia)
}

// IndexSocialMedia godoc
// @Summary Get all social media details
// @Description Get details of all social medias
// @Tags social-media
// @Accept json
// @Produces json
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/ [get]
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
	c.JSON(http.StatusOK, SocialMedias)
}

// ShowSocialMedia godoc
// @Summary Get social media details
// @Description Get details of one social media by its id
// @Tags social-media
// @Accept json
// @Produces json
// @Param id path int true "ID of the social media to be shown"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{id} [get]
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
	c.JSON(http.StatusOK, SocialMedia)
}

// EditSocialMedia godoc
// @Summary Update an existing social media
// @Description Update an existing social media with new name and url
// @Tags social-media
// @Accept json
// @Produces json
// @Param id path int true "ID of the social media to be updated"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{id} [patch]
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
		"id": SocialMedia.ID,
		"user_id": SocialMedia.UserID,
		"updated_name": SocialMedia.Name,
		"updated_social_media_url": SocialMedia.SocialMediaUrl,
		"updated_at": SocialMedia.UpdatedAt,
	})
}

// DestroySocialMedia godoc
// @Summary Delete an existing social media
// @Description Just delete an existing social media
// @Tags social-media
// @Accept json
// @Produces json
// @Param id path int true "ID of the social media to be deleted"
// @Success 204 "No Content"
// @Router /social-medias/{id} [delete]
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