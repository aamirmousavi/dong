package bind

import "github.com/gin-gonic/gin"

func BindHeader[T any](ctx *gin.Context) (*T, error) {
	param := new(T)
	if err := ctx.BindHeader(param); err != nil {
		return nil, err
	}
	return param, nil
}
