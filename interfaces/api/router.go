package api

import (
	"payconfig/interfaces/api/company"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 创建路由组
	CreateRouteGroup(router)
	return router
}

// IRouteRegister 定义路由注册的接口
type IRouteRegister interface {
	RegisterRoutes(rgPublic *gin.RouterGroup, rgPrivate *gin.RouterGroup)
}

// 创建路由组权限
func CreateRouteGroup(r *gin.Engine) {

	// 初始化路由模块
	routeRegisters := initBaseRoutes()

	public := r.Group("/pay/public") // 公共
	auth := r.Group("/pay/auth")     // 私有

	// 注册路由
	for _, routeRegister := range routeRegisters {
		routeRegister.RegisterRoutes(public, auth)
	}
}

// 初始化路由模块
func initBaseRoutes() []IRouteRegister {
	var r []IRouteRegister
	// 初始化路由注册器
	r = append(r, &company.API{})
	return r
}
