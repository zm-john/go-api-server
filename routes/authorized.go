package routes

import (
	handler "best.me/handler/user"
	"best.me/middleware"
	"github.com/gin-gonic/gin"
)

func authorized(route *gin.Engine) {
	authorized := route.Group("/")
	authorized.Use(middleware.Auth())
	authorized.GET("/users", handler.List)
}
