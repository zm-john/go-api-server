package routes

import "github.com/gin-gonic/gin"
import handler "best.me/handler/unauthorized"

func unauthorized(route *gin.Engine) {
	route.POST("/login", handler.Login)
}
