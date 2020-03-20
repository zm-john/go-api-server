package models

import (
	"best.me/config"
	"best.me/database"
	"best.me/utils"
	"fmt"
	"log"
	"time"
)

var jwtConfig utils.JWTConfig

// User 用户
type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"Column:username;type:varchar(24);UNIQUE"`
	Password  string `gorm:"Column:password;NOT NULL" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// NewToken create token from user
func (u *User) NewToken() (*utils.Token, error) {
	sub := fmt.Sprintf("%d", u.ID)
	return utils.NewToken(sub, jwtConfig)
}

// FindUserByID 通过 ID 获取 User
func FindUserByID(id string) (*User, error) {
	user := new(User)
	db := database.MysqlOrm.Where("id = ?", id).First(user)

	errs := db.GetErrors()
	if len(errs) > 0 {
		log.Printf("[warning] FindUserByID %v", errs)
		return nil, errs[0]
	}

	return user, nil
}

// NewUserFromToken parse user from token
func NewUserFromToken(token string) (*User, error) {
	sub, err := utils.ParseToken(token, jwtConfig)
	if err != nil {
		return nil, err
	}

	user := new(User)
	database.MysqlOrm.Where("id = ?", sub).First(user)
	return user, nil
}

func init() {
	conf := config.GetJWTConfig()
	jwtConfig = utils.JWTConfig{Key: conf.Key, TTL: conf.TTL}
}
