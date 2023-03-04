package api

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"github.com/gin-gonic/gin"
	"github.com/primasio/wormhole/db"
	"github.com/primasio/wormhole/http/middlewares"
	"github.com/primasio/wormhole/http/token"
	"github.com/primasio/wormhole/models"
)

type UserController struct{}

type LoginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Remember string `form:"remember" json:"remember"`
}

type RegisterForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Nickname string `form:"nickname" json:"nickname" binding:"required"`
}

func (ctrl *UserController) Create(c *gin.Context) {

	var form RegisterForm

	if err := c.ShouldBind(&form); err != nil {
		Error(err.Error(), c)
	} else {
		dbi := db.GetDb()

		// Check username uniqueness

		exist := &models.User{}
		exist.Username = form.Username

		dbi.Where(&exist).First(&exist)

		if exist.ID != 0 {
			Error("Username exists", c)
			return
		}

		// Save user to db

		user := &models.User{}
		user.Username = form.Username
		user.Password = form.Password
		user.Nickname = form.Nickname

		if err := user.SetUniqueID(dbi); err != nil {
			ErrorServer(err, c)
		}

		dbi2 := dbi.Create(&user)

		if dbi2.Error != nil {
			ErrorServer(dbi2.Error, c)
		}

		Success(user, c)
	}
}

func (ctrl *UserController) Get(c *gin.Context) {

	userId, _ := c.Get(middlewares.AuthorizedUserId)

	user := &models.User{}
	user.ID = userId.(uint)

	dbi := db.GetDb()
	dbi.First(&user)

	if user.CreatedAt == 0 {
		Error("User not found", c)
		return
	}

	Success(user, c)
}

func (ctrl *UserController) Auth(c *gin.Context) {

	var login LoginForm

	if err := c.ShouldBind(&login); err != nil {
		Error(err.Error(), c)
	} else {

		user := &models.User{Username: login.Username}

		dbi := db.GetDb()
		dbi.Where("username = ?", user.Username).First(&user)

		if user.ID == 0 {
			ErrorUnauthorized("User not found", c)
			return
		}

		if !user.VerifyPassword(login.Password) {
			ErrorUnauthorized("Incorrect password", c)
		} else {

			// Login success, generate token
			err, accessToken := token.IssueToken(user.ID, login.Remember == "")

			if err != nil {
				ErrorServer(err, c)
				return
			}

			Success(accessToken, c)
		}
	}
}
