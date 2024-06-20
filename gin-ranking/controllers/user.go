package controllers

import (
	"gin-ranking/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct{}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	//获取参数信息
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}

	user, _ := models.GetUserInfoByUsername(username)
	if user.Id == 0 || user.Password != EncryMd5(password) {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	data := UserApi{Id: user.Id, Username: user.Username}
	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(user.Id), user.Id)
	session.Save()

	ReturnSuccess(c, 0, "success", data, 1)
}

func (u UserController) Register(c *gin.Context) {
	//获取参数信息
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")
	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "请输入正确的信息")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4001, "密码和确认密码不相同")
		return
	}

	user, err := models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "此用户名已存在")
		return
	}
	_, err = models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4002, "注册失败，请重试")
		return
	}

	ReturnSuccess(c, 0, "success", "", 1)
}
