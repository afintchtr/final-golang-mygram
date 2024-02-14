package routers

import (
	"final-golang-mygram/controllers"
	"final-golang-mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users") 
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
	}
	socialMediaRouter := r.Group("/social-medias") 
	{
		socialMediaRouter.Use(middlewares.Authentication())

		socialMediaRouter.POST("/create", controllers.StoreSocialMedia)
		socialMediaRouter.GET("/", controllers.IndexSocialMedia)
		socialMediaRouter.GET("/:id", middlewares.SocialMediaAuthorization(), controllers.ShowSocialMedia)
		socialMediaRouter.PATCH("/update/:id", middlewares.SocialMediaAuthorization(), controllers.EditSocialMedia)
		socialMediaRouter.DELETE("/delete/:id", middlewares.SocialMediaAuthorization(), controllers.DestroySocialMedia)
	}


	return r
}