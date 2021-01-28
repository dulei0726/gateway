package controller

import (
    "github.com/dulei0726/gateway/apps/user/dto"
    "github.com/dulei0726/gateway/apps/user/service"
    "github.com/dulei0726/gateway/pkg"
    "github.com/gin-gonic/gin"
)

type LoginLogoutController struct {
}

// @Summary 登录
// @Accept json
// @Produce json
// @Tags user
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.TagSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (controller *LoginLogoutController) Login(c *gin.Context) {
    loginRequest := &dto.LoginRequest{}
    response := pkg.NewResponse(c)
    err := pkg.BindAndValid(c, loginRequest)
    if err != nil {
        response.ToErrorResponse(err)
        return
    }

    svc := service.NewLoginLogoutService()
    respData, err := svc.Login(loginRequest)
    if err != nil {
        response.ToErrorResponse(err)
        return
    }
    response.ToResponse(respData)
    return
}

func (controller *LoginLogoutController) Logout(c *gin.Context) {

}
