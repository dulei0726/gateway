package errcode

import "net/http"

// 通用
var (
    Success      = NewError(0, http.StatusOK, "成功")
    ServerError  = NewError(10000001, http.StatusInternalServerError, "服务内部错误")
    UnknownError = NewError(10000002, http.StatusBadRequest, "未知错误")
)

// 认证
var (
    UnauthorizedUserNotExist  = NewError(10001001, http.StatusUnauthorized, "鉴权失败，用户不存在")
    UnauthorizedUserInactive  = NewError(10001002, http.StatusUnauthorized, "鉴权失败，用户不可用")
    UnauthorizedTokenError    = NewError(10001003, http.StatusUnauthorized, "鉴权失败，Token错误")
    UnauthorizedTokenTimeout  = NewError(10001004, http.StatusUnauthorized, "鉴权失败，Token超时")
    UnauthorizedTokenGenerate = NewError(10001005, http.StatusUnauthorized, "鉴权失败，Token生成失败")
    UnauthorizedNoToken       = NewError(10001006, http.StatusUnauthorized, "鉴权失败，未携带Token")
)

// 鉴权
var ()

// 限流
var (
    TooManyRequests = NewError(10003001, http.StatusTooManyRequests, "请求过于频繁")
)

// 参数校验
var (
    InvalidParams = NewError(10004001, http.StatusBadRequest, "入参错误")
)

// 服务发现
var (
    ServiceNotFound = NewError(10005001, http.StatusBadRequest, "服务不存在")
    ServiceEmpty    = NewError(10005002, http.StatusBadRequest, "无可用服务")
)
