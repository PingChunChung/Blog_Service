package routers

import (
	_ "blog-service/docs"
	"blog-service/global"
	"blog-service/internal/middleware"
	"blog-service/internal/routers/api"
	v1 "blog-service/internal/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	tag := v1.NewTag()
	article := v1.NewArticle()
	upload := api.NewUpload()
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	r.POST("/upload/file", upload.UploadFile)
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.GET("/tags", tag.List)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.DELETE("/tags/:id", tag.Delete)

		apiv1.POST("/articles", article.Create)
		apiv1.GET("/articles", article.List)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.DELETE("/articles/:id", article.Delete)
	}

	return r
}
