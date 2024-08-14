package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/constants/http"
	"github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/logging"
	"github.com/hipeday/upay/internal/repository"
	"github.com/hipeday/upay/internal/service"
	"github.com/hipeday/upay/pkg/config"
	"strings"
)

// BearerAuthorizationMiddleware 验证 Bearer Token
// 该中间件用于验证请求头中的Authorization字段是否为Bearer Token
// 如果不是Bearer Token则返回401 Unauthorized
// 如果是Bearer Token但是验证失败则返回401 Unauthorized
// 如果是Bearer Token且验证成功则放行
// 该中间件适用于需要验证Token的接口
// 用例: curl -H "Authorization: Bearer testing" http://127.0.0.1:5266/use/account/1
func BearerAuthorizationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		logging.Logger().Debugf("Get Authorization Header: %s", authHeader)

		if authHeader == "" {
			panic(errors.NewUnauthorizedError("Unauthorized"))
		}

		// 获取值结果应该是 `Bearer ${token}`
		parts := strings.Split(authHeader, " ")

		logging.Logger().Debugf("Parse Authorization Header: %s", parts)

		if len(parts) != 2 || parts[0] != "Bearer" {
			panic(errors.NewUnauthorizedError("Unauthorized"))
		}

		token := parts[1]

		cfg := config.GetCfg()
		db, err := repository.InitMySQL(cfg)
		if err != nil {
			logging.Logger().Errorf("create mysql instance err: %v", err)
			panic(errors.NewInternalServerError("Server Error"))
		}

		tokenService := service.GetTokenServiceInstance(db)

		tokenEntity, validToken, err := tokenService.IsValidToken(token)

		logging.Logger().Debugf("Valid token result: %t token entity: %v", validToken, tokenEntity)
		if err != nil {
			logging.Logger().Errorf("check token err: %v", err)
			panic(errors.NewInternalServerError("Server Error"))
		}

		if !validToken {
			panic(errors.NewUnauthorizedError("Unauthorized"))
		}

		context.Set(http.AccountIdContext, tokenEntity.TargetId)
		context.Set(http.AccountTokenContext, token)

		logging.Logger().Debugf("Verification authentication passed. Add attribute account_id: %d token: %s to context", tokenEntity.ID, token)

		// 放行
		context.Next()
	}
}
