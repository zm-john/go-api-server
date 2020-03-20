package routes

import "github.com/gin-gonic/gin"

// AddRoutes 添加路由
func AddRoutes(router *gin.Engine) {
	unauthorized(router)
	authorized(router)
}
