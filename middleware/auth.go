package middleware

import (
	"best.me/config"
	"best.me/handler"
	"best.me/models"
	"best.me/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Auth 登录认证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Access-Token")

		// if len(tokenString) == 0 {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, handler.NewErrResponse("请登录有访问", nil))
		// 	return
		// }

		conf := config.GetJWTConfig()
		sub, err := utils.ParseToken(tokenString, utils.JWTConfig{Key: conf.Key, TTL: conf.TTL})

		if err != nil {
			switch err.(type) {
			case utils.TokenExpiredError:
				c.AbortWithStatusJSON(http.StatusUnauthorized, handler.NewErrResponse(err.Error(), nil))
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, handler.NewErrResponse("token 解析失败", nil))
			}
			return
		}

		user, err := models.FindUserByID(sub)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, handler.NewErrResponse("获取用户信息错误", nil))
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
