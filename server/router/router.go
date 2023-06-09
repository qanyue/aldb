package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qanyue/aldb/server/api/v1"
	docs "github.com/qanyue/aldb/server/docs"
	"github.com/qanyue/aldb/server/middleware"
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
	r.POST("uploadexample", v1.Uploadexamle)
	// api
	api := r.Group("api")
	//api.Use(middleware.JWTAuthMiddleware())
	{
		algae := api.Group("alga")
		{
			algae.GET("anno", v1.GetAnnotationByAlga)
			algae.POST("search", v1.SearchAlga)
			algae.POST("add", v1.AddAlga)
			algae.POST("addMore", v1.AddAlgas)
			algae.GET("all", v1.GetAllAlgas)
		}
		anno := api.Group("anno")
		{
			anno.POST("delete", v1.DeleteAnnotation)
			anno.POST("add", v1.AddAnnotation)
			//anno.POST("update", v1.UpdateAnnotation)
		}
		user := api.Group("user")
		{
			user.GET("info", v1.GetUser)
			//user.GET("anno", v1.GetAnnotationByUser)
			user.GET("all", v1.GetUsers)
			user.GET("delete", v1.DeleteUser)
			user.POST("pwd", v1.ChangePassword)
			user.POST("update", v1.UpdateUser)
		}
		river := api.Group("river")
		{
			river.GET("info", v1.GetRiverInfo)
			river.POST("all", v1.GetRivers)
			river.POST("add", v1.AddRiver)
			river.POST("search", v1.SearchRiver)
			river.POST("share", v1.ShareRiver)
			//TODO：删除数据集
		}
		tag := api.Group("tag")
		{
			tag.GET("all", v1.GetTags)
			tag.POST("add", v1.AddTag)
			tag.POST("delete", v1.DeleteTag)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
	return r
}
