package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/logging"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logging.Logger().Errorf("System exception interception: %v", r)
				respondWithError(context, handleError(r))
			}
		}()
		context.Next()
	}
}

func handleError(r interface{}) errors.Errors {
	switch e := r.(type) {
	case errors.IllegalArgumentError:
		return e
	case errors.UnauthorizedError:
		return e
	case errors.ConflictError:
		return e
	default:
		// 处理未知的panic
		return errors.NewInternalServerError(fmt.Sprintf("Unknown error: %v", r))
	}
}

func respondWithError(c *gin.Context, err errors.Errors) {
	// 根据错误类型设置HTTP状态码和返回错误信息
	c.AbortWithStatusJSON(err.GetStatus(), gin.H{
		"error": err.GetError(),
	})
}
