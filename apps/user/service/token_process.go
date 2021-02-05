package service

import "time"

type TokenProcessor interface {
    // 授予token
    Grant(expiresTime *time.Time) (*Token, error)
}

// JWTTokenProcess 使用JWT实现token处理
type JWTTokenProcess struct {}

func NewJWTTokenProcess() *JWTTokenProcess {
    return &JWTTokenProcess{}
}

func (grant *JWTTokenProcess) Grant(expiresTime *time.Time) (*Token, error) {
    return nil, nil
}
