package service

import "github.com/dulei0726/gateway/apps/user/dto"

type LoginLogoutService struct {
}

func (service *LoginLogoutService) Login(request *dto.LoginRequest) *dto.LoginResponse {
    return nil
}

func (service *LoginLogoutService) Logout() {

}
