package middleware

import (
    "github.com/dulei0726/gateway/apps/user/service"
    "github.com/dulei0726/gateway/pkg"
    "github.com/dulei0726/gateway/pkg/errcode"
    "github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
    return func(c *gin.Context) {
        accessToken, err := c.Cookie(pkg.AccessTokenCookieName)
        if err != nil {
            pkg.NewResponse(c).ToErrorResponse(errcode.UnauthorizedNoToken)
            c.Abort()
            return
        }
        refreshToken, err := c.Cookie(pkg.RefreshTokenCookieName)
        if err != nil {
            pkg.NewResponse(c).ToErrorResponse(errcode.UnauthorizedNoToken)
            c.Abort()
            return
        }

        tokenProcess := service.NewJWTTokenProcess()
        tokenStorage := service.NewJWTTokenStorage()
        tokenService := service.NewTokenService(tokenProcess, tokenStorage)
        token, info, err := tokenService.ParseTokenValue(accessToken, refreshToken)
        if err != nil {
            pkg.NewResponse(c).ToErrorResponse(err)
            c.Abort()
            return
        }
        c.Set("oauth2Token", token)
        c.Set("userInfo", info)

        c.Next()

        c.SetCookie(pkg.AccessTokenCookieName, token.AccessToken.GetTokenValue(), pkg.TokenCookieMaxAge, pkg.TokenCookiePath,
            pkg.TokenCookieDomain, pkg.TokenCookieSecure, pkg.TokenCookieHttpOnly)
        c.SetCookie(pkg.RefreshTokenCookieName, token.RefreshToken.GetTokenValue(), pkg.TokenCookieMaxAge, pkg.TokenCookiePath,
            pkg.TokenCookieDomain, pkg.TokenCookieSecure, pkg.TokenCookieHttpOnly)
    }
}
