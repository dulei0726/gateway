package main

import (
	"github.com/dulei0726/gateway/apps/user/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	c := &controller.LoginLogoutController{}
	r.POST("/login", c.Login)
	r.Run(":8080")
}
