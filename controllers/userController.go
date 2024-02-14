package controllers

import (
	"final-golang-mygram/database"
	"final-golang-mygram/helpers"
	"final-golang-mygram/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var appJSON = "application/json"

func Register(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": User.ID,
		"email": User.Email,
		"username": User.Username,
	})
}

func Login(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Invalid email/password:",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "Invalid email/password:",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func IndexUser(c *gin.Context) {
	db := database.GetDB()

	Users := []models.User{}

	err := db.Model(&Users).Preload("Photos.Comments").Preload("Comments").Preload("SocialMedias").Order("id").Find(&Users).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	UsersResponse := []models.UserResponse{}
	for _, User := range Users {
		UserResponse := models.UserResponse{}

		UserResponse.ID = User.ID
		UserResponse.Email = User.Email
		UserResponse.Username = User.Username
		UserResponse.Age = User.Age
		UserResponse.Photos = User.Photos
		UserResponse.SocialMedias = User.SocialMedias
		UserResponse.Comments = User.Comments

		UsersResponse = append(UsersResponse, UserResponse)
	}


	c.JSON(http.StatusOK, UsersResponse)
}
