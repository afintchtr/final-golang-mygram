package routers

import (
	"final-golang-mygram/controllers"
	"final-golang-mygram/middlewares"

	"github.com/gin-gonic/gin"

	_ "final-golang-mygram/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API (Afin)
// @version 1.0
// @description This is what ive been practiced in Golang class Bootcamp BRI Golang for Final project
// @termsOfService http://swagger.io/terms/
// @contact.name Afin Tachtiar
// @contact.email afintchtr@gmail.com
// @license.name AfinT
// @license.url https://github.com/afintchtr
// @host localhost:8085
// @BasePath /
func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users") 
	{
		userRouter.POST("/register", controllers.Register)
		userRouter.POST("/login", controllers.Login)
		userRouter.GET("/", controllers.IndexUser)
	}
	socialMediaRouter := r.Group("/social-medias") 
	{
		socialMediaRouter.Use(middlewares.Authentication())

		socialMediaRouter.POST("/create", controllers.StoreSocialMedia)
		socialMediaRouter.GET("/", controllers.IndexSocialMedia)
		socialMediaRouter.GET("/:id", middlewares.SocialMediaAuthorization(), controllers.ShowSocialMedia)
		socialMediaRouter.PATCH("/:id", middlewares.SocialMediaAuthorization(), controllers.EditSocialMedia)
		socialMediaRouter.DELETE("/:id", middlewares.SocialMediaAuthorization(), controllers.DestroySocialMedia)
	}
	photoRouter := r.Group("/photos") 
	{
		photoRouter.Use(middlewares.Authentication())

		photoRouter.POST("/create", controllers.StorePhoto)
		photoRouter.GET("/", controllers.IndexPhoto)
		photoRouter.GET("/:id", middlewares.PhotoAuthorization(), controllers.ShowPhoto)
		photoRouter.PATCH("/:id", middlewares.PhotoAuthorization(), controllers.EditPhoto)
		photoRouter.DELETE("/:id", middlewares.PhotoAuthorization(), controllers.DestroyPhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())

		commentRouter.POST("/create/:photoId", controllers.StoreComment)
		commentRouter.GET("/", controllers.IndexComment)
		commentRouter.GET("/:commentId", middlewares.CommentAuthorization(), controllers.ShowComment)
		commentRouter.PATCH("/:commentId", middlewares.CommentAuthorization(), controllers.EditComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DestroyComment)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
