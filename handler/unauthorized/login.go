package unauthorized

import (
	"best.me/database"
	"best.me/handler"
	"best.me/models"
	"best.me/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	// "log"
	"net/http"
)

// LoginForm 登录表单
// swagger:parameters login
type LoginForm struct {
	// in: body
	Body LoginFormBody
}

// LoginFormBody 登录表单参数
type LoginFormBody struct {
	// Required: true
	Username string `json:"username" validate:"required"`
	// Required: true
	Password string `json:"password" validate:"required"`
}

// LoginResponse 登录响应
// swagger:response LoginResponse
type LoginResponse struct {
	// in: body
	Body LoginResponseBody
}

// LoginResponseBody 登录响应内容
type LoginResponseBody struct {
	// 用户信息
	Data *models.User `json:"data"`
	// Token 信息
	Meta *utils.Token `json:"meta"`
}

// Login 登录
// swagger:route POST /login user login
//
// 用户登录
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Deprecated: false
//
//     Responses:
//       200: LoginResponse
//       422: error
func Login(c *gin.Context) {
	params := LoginFormBody{}
	if err := c.ShouldBind(&params); err != nil {
		// 数据异常
		c.JSON(http.StatusBadRequest, handler.NewErrResponse(err.Error(), []handler.Error{}))
		return
	}

	if err := handler.Validate.Struct(params); err != nil {
		errs := handler.ValidateErrorToErrResponse(err)
		c.JSON(http.StatusUnprocessableEntity, handler.NewErrResponse("参数错误", errs))
		return
	}

	user := new(models.User)
	database.MysqlOrm.Where("username = ?", params.Username).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, handler.NewErrResponse("用户不存在", nil))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		// 密码错误
		c.JSON(http.StatusUnauthorized, handler.NewErrResponse("账号或密码错误", []handler.Error{}))
		return
	}

	// 生成 token
	token, err := user.NewToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, handler.NewErrResponse("生成 token 错误", []handler.Error{}))
		return
	}

	c.JSON(http.StatusOK, LoginResponseBody{user, token})
}
