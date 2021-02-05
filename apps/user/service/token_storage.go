package service

type TokenStorage interface {
    Set(info *UserInfo, token *OAuth2Token) error
    Get(accessTokenValue, refreshTokenValue string) (*OAuth2Token, *UserInfo, error)
    Remove(accessTokenValue, refreshTokenValue string)
}

type JWTTokenStorage struct {}

func NewJWTTokenStorage() *JWTTokenStorage {
    return &JWTTokenStorage{}
}

func (storage *JWTTokenStorage) Set(info *UserInfo, token *OAuth2Token) error {
    panic("implement me")
}

func (storage *JWTTokenStorage) Get(accessTokenValue, refreshTokenValue string) (*OAuth2Token, *UserInfo, error) {
    panic("implement me")
}

func (storage *JWTTokenStorage) Remove(accessTokenValue, refreshTokenValue string) {
    panic("implement me")
}



