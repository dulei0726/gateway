package service

import "github.com/dulei0726/gateway/apps/user/dto"

type LoginLogoutService struct {
}

func NewLoginLogoutService() *LoginLogoutService {
    return &LoginLogoutService{}
}

func (service *LoginLogoutService) Login(request *dto.LoginRequest) *dto.LoginResponse {
    return nil
}

func (service *LoginLogoutService) Logout() {

}
