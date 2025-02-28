package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jcvv0n/min-reader/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	// r.SetFuncMap(template.FuncMap{
	// 	"getStroyCat": service.GetStroyCat,
	// })
	r.LoadHTMLGlob("template/*")
	r.GET("/:namespace/stos", service.Storys)
	r.GET("/:namespace/:storyId/cat", service.Catalog)
	r.GET("/:namespace/:storyId/ft/dl", service.Fulltext)
	r.GET("/:namespace/:storyId/cont", service.Reader)
	return r
}
