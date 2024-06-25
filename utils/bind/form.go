package bind

import "github.com/gin-gonic/gin"

func Bind[T any](ctx *gin.Context) (*T, error) {
	param := new(T)
	if err := ctx.Bind(param); err != nil {
		return nil, err
	}
	return param, nil
}
