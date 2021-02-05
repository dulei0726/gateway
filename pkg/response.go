package pkg

import (
	"github.com/dulei0726/gateway/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Ctx *gin.Context
}

type responseData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}) {
	r.Ctx.JSON(errcode.Success.StatusCode(), responseData{
		Code: errcode.Success.Code(),
		Msg:  errcode.Success.Msg(),
		Data: data,
	})
}

func (r *Response) ToResponseList(list interface{}, totalRows int) {
	respData := responseData{
		Code: errcode.Success.Code(),
		Msg:  errcode.Success.Msg(),
		Data: map[string]interface{}{
			"rows": list,
			"pager": Pager{
				Page:      GetPage(r.Ctx),
				PageSize:  GetPageSize(r.Ctx),
				TotalRows: totalRows,
			},
		},
	}
	r.Ctx.JSON(errcode.Success.StatusCode(), respData)
}

func (r *Response) ToErrorResponse(err error) {
	e, ok := err.(*errcode.Error)
	if !ok {
		e = errcode.UnknownError.WithDetails(err.Error())
	}
	r.Ctx.JSON(e.StatusCode(), responseData{
		Code: e.Code(),
		Msg:  e.Msg(),
		Data: e.Details(),
	})
}
