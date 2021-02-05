package service

import (
    "github.com/dulei0726/gateway/pkg"
    "github.com/dulei0726/gateway/pkg/errcode"
    "time"
)

type TokenType string

const (
    JWTToken    TokenType = "JWT"
    BearerToken TokenType = "Bearer"
    BasicToken  TokenType = "Basic"
    MACToken    TokenType = "MAC"
)

type Token struct {
    tokenType   TokenType  // token类型
    tokenValue  string     // token值
    expiresTime *time.Time // 过期时间
}

type OAuth2Token struct {
    AccessToken  *Token
    RefreshToken *Token
}

func (t *Token) GetTokenType() TokenType {
    return t.tokenType
}

func (t *Token) GetTokenValue() string {
    return t.tokenValue
}

func (t *Token) IsExpired() bool {
    return t.expiresTime != nil && t.expiresTime.Before(time.Now())
}

type TokenService interface {
    CreateOAuth2Token(accessExpiresTime, refreshExpiresTime *time.Time, userInfo *UserInfo) (*OAuth2Token, error)
    ParseTokenValue(accessTokenValue, refreshTokenValue string) (*OAuth2Token, *UserInfo, error)
}

type DefaultTokenService struct {
    tokenProcess TokenProcessor
    tokenStorage TokenStorage
}

func NewTokenService(tokenProcess TokenProcessor, tokenStorage TokenStorage) *DefaultTokenService {
    return &DefaultTokenService{
        tokenProcess: tokenProcess,
        tokenStorage: tokenStorage,
    }
}

func (svc *DefaultTokenService) CreateOAuth2Token(userInfo *UserInfo) (*OAuth2Token, error) {
    var (
        now                = time.Now()
        accessExpiresTime  = now.Add(pkg.AccessTokenDuration)
        refreshExpiresTime = now.Add(pkg.RefreshTokenDuration)
    )
    accessToken, err := svc.tokenProcess.Grant(&accessExpiresTime)
    if err != nil {
        return nil, errcode.UnauthorizedTokenGenerate.WithDetails(err.Error())
    }
    refreshToken, err := svc.tokenProcess.Grant(&refreshExpiresTime)
    if err != nil {
        return nil, errcode.UnauthorizedTokenGenerate.WithDetails(err.Error())
    }
    oauth2Token := &OAuth2Token{AccessToken: accessToken, RefreshToken: refreshToken}
    err = svc.tokenStorage.Set(userInfo, oauth2Token)
    if err != nil {
        return nil, errcode.UnauthorizedTokenGenerate.WithDetails(err.Error())
    }
    return oauth2Token, nil
}

func (svc *DefaultTokenService) ParseTokenValue(accessTokenValue, refreshTokenValue string) (*OAuth2Token, *UserInfo, error) {
    token, info, err := svc.tokenStorage.Get(accessTokenValue, refreshTokenValue)
    if err != nil {
        return nil, nil, errcode.UnauthorizedTokenError.WithDetails(err.Error())
    }
    if !token.AccessToken.IsExpired() {
        return token, info, nil
    }
    if token.RefreshToken.IsExpired() {
        return nil, nil, errcode.UnauthorizedTokenTimeout
    }
    // AccessToken过期但RefreshToken未过期, 自动更新token
    svc.tokenStorage.Remove(accessTokenValue, refreshTokenValue)
    newToken, err := svc.CreateOAuth2Token(info)
    if err != nil {
        return token, info, nil
    }
    return newToken, info, nil
}
