package controller

import (
    "fmt"
    "github.com/dulei0726/gateway/apps/user/dto"
    "github.com/dulei0726/gateway/apps/user/service"
    "github.com/gin-gonic/gin"
)

type LoginLogoutController struct {
}

func (controller *LoginLogoutController) Login(c *gin.Context) {
    loginRequest := &dto.LoginRequest{}
    err := c.ShouldBind(loginRequest)
    if err != nil {
        fmt.Printf("ShouldBind error: %#v", err.Error())
    }
    svc := service.NewLoginLogoutService()
    respData := svc.Login(loginRequest)
    c.JSON(200, gin.H{"data": respData})
    return
}

func (controller *LoginLogoutController) Logout(c *gin.Context) {

}
