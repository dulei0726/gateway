package service

import (
    "github.com/dulei0726/gateway/apps/user/dao"
    "github.com/dulei0726/gateway/apps/user/dto"
)

type LoginLogoutService struct {
    dao.Dao
}

func NewLoginLogoutService() *LoginLogoutService {
    return &LoginLogoutService{
        //dao.New(),
    }
}

func (service *LoginLogoutService) Login(request *dto.LoginRequest) (*dto.LoginResponse, error) {
    return nil, nil
}

func (service *LoginLogoutService) Logout() {

}
