package bind

import "github.com/gin-gonic/gin"

func Bind[T any](c *gin.Context) (*T, error) {
	param := new(T)
	if err := c.Bind(param); err != nil {
		return nil, err
	}
	return param, nil
}
