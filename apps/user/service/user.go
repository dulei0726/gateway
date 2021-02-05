package service

type UserInfo struct {
    UserID      int    // 用户标识
    Username    string // 用户名(唯一)
    Groups      []int  // 用户所属组
    Authorities []int  // 组之外的权限
}

type UserInfoService interface {
    // 通过用户名，密码获取用户信息
    GetUserInfo(username, password string) (*UserInfo, error)
}
