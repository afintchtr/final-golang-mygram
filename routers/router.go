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
		userRouter.GET("/", controllers.IndexUser)
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
	photoRouter := r.Group("/photos") 
	{
		photoRouter.Use(middlewares.Authentication())

		photoRouter.POST("/create", controllers.StorePhoto)
		photoRouter.GET("/", controllers.IndexPhoto)
		photoRouter.GET("/:id", middlewares.PhotoAuthorization(), controllers.ShowPhoto)
		photoRouter.PATCH("/update/:id", middlewares.PhotoAuthorization(), controllers.EditPhoto)
		photoRouter.DELETE("/delete/:id", middlewares.PhotoAuthorization(), controllers.DestroyPhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())

		commentRouter.POST("/create/:photoId", controllers.StoreComment)
		commentRouter.GET("/", controllers.IndexComment)
		commentRouter.GET("/:commentId", middlewares.CommentAuthorization(), controllers.ShowComment)
		commentRouter.PATCH("/update/:commentId", middlewares.CommentAuthorization(), controllers.EditComment)
		commentRouter.DELETE("/delete/:commentId", middlewares.CommentAuthorization(), controllers.DestroyComment)
	}

	return r
}
