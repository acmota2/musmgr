package middleware

import (
	"github.com/gin-gonic/gin"
)

func GetContextValue[T any](c *gin.Context, key string) (val T, ok bool) {
	anyVal, ok := c.Get(key)
	if !ok {
		return val, false
	}
	val, ok = anyVal.(T)
	if !ok {
		return val, false
	}
	return val, true
}
