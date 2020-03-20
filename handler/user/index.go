package user

import (
	"best.me/database"
	"best.me/handler"
	"best.me/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// swagger:parameters userList
type userListQuery struct {
	// in: query
	Keyword string
	Page    int
}

// List 获取用户列表
// swagger:route GET /users user users
//
// Get user list
//     Deprecated: false
// 	   Security:
//       access-authorized:
//     Responses:
//       200: LoginResponse
// 		 default: error
func List(c *gin.Context) {
	query := &userListQuery{}
	if err := c.ShouldBindQuery(query); err != nil {
		c.JSON(http.StatusBadRequest, handler.NewErrResponse("参数错误", nil))
		return
	}

	db := database.MysqlOrm.Order("id DESC")
	var users []models.User
	if len(query.Keyword) > 0 {
		db = db.Where("username like ?", fmt.Sprintf("%s%s%s", "%", query.Keyword, "%"))
	}
	db = db.Find(&users)

	if len(db.GetErrors()) > 0 {
		c.JSON(http.StatusBadRequest, handler.NewErrResponse("数据库异常", nil))
		return
	}

	c.JSON(http.StatusOK, handler.NewResponse(users, nil))
}
