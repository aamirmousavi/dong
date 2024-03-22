package bind

import "github.com/gin-gonic/gin"

func BindHeader[T any](c *gin.Context) (*T, error) {
	param := new(T)
	if err := c.BindHeader(param); err != nil {
		return nil, err
	}
	return param, nil
}
