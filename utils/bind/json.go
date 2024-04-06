package bind

import "github.com/gin-gonic/gin"

func BindJson[T any](ctx *gin.Context) (*T, error) {
	param := new(T)
	if err := ctx.BindJSON(param); err != nil {
		return nil, err
	}
	return param, nil
}
