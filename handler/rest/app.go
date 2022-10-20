package rest

import (
	"mygram/database"
	"mygram/docs"
	"mygram/repository/comment_repository/comment_pg"
	"mygram/repository/photo_repository/photo_pg"
	"mygram/repository/social_media_repository/social_media_pg"
	"mygram/repository/user_repository/user_pg"
	"mygram/service"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var DEFAULT_PORT = ":8080"

func StartApp() {
	database.StartDB()
	db := database.GetDB()

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userRestHandler := NewUserRestHandler(userService)

	photoRepo := photo_pg.NewPhotoPG(db)
	photoService := service.NewPhotoService(photoRepo)
	photoRestHandler := NewPhotoRestHandler(photoService)

	commentRepo := comment_pg.NewCommentPG(db)
	commentService := service.NewCommentService(commentRepo)
	commentRestHandler := NewCommentRestHandler(commentService)

	socialMediaRepo := social_media_pg.NewSocialMediaPG(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepo)
	socialMediaRestHandler := NewSocialMediaRestHandler(socialMediaService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, socialMediaRepo)

	// ! Routing
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	userRoute := router.Group("/users")
	userRoute.POST("/login", userRestHandler.Login)
	userRoute.POST("/register", userRestHandler.Register)
	userRoute.PUT("/", authService.Authentication(), userRestHandler.UpdateUserData)
	userRoute.DELETE("/", authService.Authentication(), userRestHandler.DeleteUser)

	photoRoute := router.Group("/photos")
	photoRoute.Use(authService.Authentication())
	photoRoute.POST("/", photoRestHandler.PostPhoto)
	photoRoute.GET("/", photoRestHandler.GetAllPhotos)
	photoRoute.PUT("/:photoID", authService.PhotoAuthorization(), photoRestHandler.UpdatePhoto)
	photoRoute.DELETE("/:photoID", authService.PhotoAuthorization(), photoRestHandler.DeletePhoto)

	commentRoute := router.Group("/comments")
	commentRoute.Use(authService.Authentication())
	commentRoute.POST("/", commentRestHandler.PostComment)
	commentRoute.GET("/", commentRestHandler.GetAllComments)
	commentRoute.PUT("/:commentID", authService.CommentAuthorization(), commentRestHandler.UpdateComment)
	commentRoute.DELETE("/:commentID", authService.CommentAuthorization(), commentRestHandler.DeleteComment)

	socialMediaRoute := router.Group("/social-medias")
	socialMediaRoute.Use(authService.Authentication())
	socialMediaRoute.POST("/", socialMediaRestHandler.AddSocialMedia)
	socialMediaRoute.GET("/", socialMediaRestHandler.GetAllSocialMedias)
	socialMediaRoute.PUT("/:socialMediaID", authService.SocialMediaAuthorization(), socialMediaRestHandler.EditSocialMediaData)
	socialMediaRoute.DELETE("/:socialMediaID", authService.SocialMediaAuthorization(), socialMediaRestHandler.DeleteSocialMedia)

	docs.SwaggerInfo.Title = "MyGrams API"
	docs.SwaggerInfo.Description = "MyGrams project is a simple REST API for social media application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost" + DEFAULT_PORT
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	router.Run(DEFAULT_PORT)
}
