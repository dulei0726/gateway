package pkg

import (
    "github.com/gin-gonic/gin"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    "strings"
)

type ValidError struct {
    Key     string `json:"key"`
    Message string `json:"message"`
}

func (ve *ValidError) Error() string {
    return ve.Message
}

type ValidErrors []*ValidError

func (ves ValidErrors) Error() string {
    return strings.Join(ves.Errors(), ";")
}

func (ves ValidErrors) Errors() []string {
    var errs = make([]string, 0, len(ves))
    for _, ve := range ves {
        errs = append(errs, ve.Error())
    }
    return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
    var ves ValidErrors
    if err := c.ShouldBind(v); err != nil {
        trans, _ := c.Value("trans").(ut.Translator)
        verrs, ok := err.(validator.ValidationErrors)
        if !ok {
            return false, ves
        }

        for key, value := range verrs.Translate(trans) {
            ves = append(ves, &ValidError{Key: key, Message: value})
        }
        return false, ves
    }
    return true, nil
}
