package main

import (
	//"github.com/liuhongdi/digv16/global"
	"github.com/gin-gonic/gin"
	"github.com/lianyz/gintest01/router"
	//"log"
)

var globalrouter *gin.Engine

func main() {
	//引入路由
	globalrouter = router.Router()
	//run
	globalrouter.Run(":8080")
}
