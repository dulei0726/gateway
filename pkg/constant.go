package pkg

import "time"

const (
    AccessTokenCookieName  string = "access_token"        // AccessToken Cookie名称
    RefreshTokenCookieName string = "refresh_token"       // RefreshToken Cookie名称
    TokenCookieMaxAge      int    = 60 * 60 * 24 * 7 * 52 // Cookie最大有效时间(秒)
    TokenCookiePath        string = "/"                   // Cookie路径
    TokenCookieDomain      string = ""                    // Cookie域名
    TokenCookieSecure      bool   = false                 // Cookie是否只在https时传递
    TokenCookieHttpOnly    bool   = true                  // Cookie是否限制js访问

    AccessTokenDuration  time.Duration = time.Second * 60 * 60 * 24     // AccessToken有效期
    RefreshTokenDuration time.Duration = time.Second * 60 * 60 * 24 * 7 // RefreshToken有效期
)

const (
    DefaultPageSize = 50  // 默认分页大小
    MaxPageSize     = 200 // 最大分页
)

const (
    ServiceDiscoveryAddress string = "127.0.0.1:8500" // 服务发现地址
)
