package router

import (
	"github.com/SukiEva/aldb/server/api/v1"
	docs "github.com/SukiEva/aldb/server/docs"
	"github.com/SukiEva/aldb/server/middleware"
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api"
	r := gin.New()
	r.Use(middleware.ZapLogger(), middleware.ZapRecovery(true))
	r.Use(middleware.Cors())
	// Auth
	r.GET("captcha", v1.GetCaptcha)
	r.POST("auth", v1.GetAuth)
	r.POST("register", v1.AddUser)
	r.POST("upload", v1.Upload)
	// api
	api := r.Group("api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		api.GET("data", v1.GetData)
		algae := api.Group("alga")
		{
			algae.GET("anno", v1.GetAnnotationByAlga)
			algae.GET("search", v1.SearchAlga)
			algae.POST("add", v1.AddAlga)
		}
		anno := api.Group("anno")
		{
			anno.GET("delete", v1.DeleteAnnotation)
			anno.POST("add", v1.AddAnnotation)
			anno.POST("update", v1.UpdateAnnotation)
		}
		user := api.Group("user")
		{
			user.GET("info", v1.GetUser)
			user.GET("anno", v1.GetAnnotationByUser)
			user.GET("all", v1.GetUsers)
			user.GET("delete", v1.DeleteUser)
			user.POST("pwd", v1.ChangePassword)
			user.POST("update", v1.UpdateUser)
		}
		river := api.Group("river")
		{
			river.GET("all", v1.GetRivers)
			river.POST("add", v1.AddRiver)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	return r
}
