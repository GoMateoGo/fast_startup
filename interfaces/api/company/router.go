package company

import (
	"github.com/gin-gonic/gin"
)

type API struct{}

func (b *API) RegisterRoutes(rgPublic *gin.RouterGroup, rgPrivate *gin.RouterGroup) {
	//auth := rgPrivate.Group("company")
	//{
	//	auth.POST("add", func(c *gin.Context) { api.JSON(c, &Add{}) })
	//	auth.GET("get", func(c *gin.Context) { api.FORM(c, &Get{}) })
	//	auth.POST("del", func(c *gin.Context) { api.FORM(c, &Delete{}) })
	//	auth.POST("update", func(c *gin.Context) { api.JSON(c, &Update{}) })
	//}
}
