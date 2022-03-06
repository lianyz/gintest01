package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lianyz/gintest01/global"
	"strconv"
)

//post struct
type TestForPost struct {
	Username string `form:"username" binding:"required,min=6"`
	Password string `form:"password" binding:"required,min=6,max=12"`
	//Number int  `form:"number" binding:"required"`
}

type UserController struct{}

func NewUserController() UserController {
	return UserController{}
}

// Get接口
func (g *UserController) GetUser(c *gin.Context) {
	res := global.NewResult(c)
	uid := c.DefaultQuery("uid", "0")
	uidint, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		res.Error(1, "用户id需要为整数")
		return
	}
	if uidint <= 0 {
		res.Error(2, "用户id<=0")
		return
	}
	res.Success("success")
	return
}

// Post接口
func (g *UserController) PostUser(c *gin.Context) {
	res := global.NewResult(c)
	req := &TestForPost{}
	if err := c.ShouldBindWith(req, binding.Form); err != nil {
		//fmt.Println("bind error",err)
		res.Error(1, "参数不匹配")
		return
	}
	res.Success("success")
}
