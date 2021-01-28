package pkg

import (
    "github.com/dulei0726/gateway/pkg/errcode"
    "github.com/gin-gonic/gin"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    "github.com/pkg/errors"
    "strings"
)

type ValidError struct {
    Key     string `json:"key"`
    Message string `json:"message"`
}

func (e *ValidError) Error() string {
    return e.Message
}

type ValidErrors []*ValidError

func (es ValidErrors) Error() string {
    return strings.Join(es.Errors(), ";")
}

func (es ValidErrors) Errors() []string {
    var errs = make([]string, 0, len(es))
    for _, e := range es {
        errs = append(errs, e.Error())
    }
    return errs
}

func BindAndValid(c *gin.Context, v interface{}) error {
    var es ValidErrors
    if err := c.ShouldBind(v); err != nil {
        trans, ok := c.Value("trans").(ut.Translator)
        if !ok {
            return errors.WithStack(err)
        }
        verrs, ok := err.(validator.ValidationErrors)
        if !ok {
            return errors.WithStack(err)
        }

        for key, value := range verrs.Translate(trans) {
            es = append(es, &ValidError{Key: key, Message: value})
        }
        return errcode.InvalidParams.WithDetails(es.Errors()...)
    }
    return nil
}
