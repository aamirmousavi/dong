package bind

import "github.com/gin-gonic/gin"

func BindJson[T any](c *gin.Context) (*T, error) {
	param := new(T)
	if err := c.BindJSON(param); err != nil {
		return nil, err
	}
	return param, nil
}
