package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hipeday/upay/internal/errors"
	"github.com/hipeday/upay/internal/logging"
	"github.com/hipeday/upay/pkg/util"
	"strings"
	"sync"
	"time"
)

var TokenStoreCache = TokenStore{
	tokens: sync.Map{},
}

// TokenStore 用于存储Token
type TokenStore struct {
	tokens sync.Map
}

// SetToken 用于设置Token 使用token作为key并且将用户的secret作为value
func (store *TokenStore) SetToken(token, secret string) {
	store.tokens.Store(token, secret)
	logging.Logger().Debugf("Token stored in cache: %s", token)
}

// RemoveToken 用于删除Token
func (store *TokenStore) RemoveToken(token string) {
	store.tokens.Delete(token)
	logging.Logger().Debugf("Token Remove in cache: %s", token)
}

// RemoveUsedToken 用于删除Token
func (store *TokenStore) RemoveUsedToken(accountId int64) {
	store.tokens.Range(func(key, value interface{}) bool {
		claims, err := util.ValidateToken(key.(string), value.(string))
		if err != nil {
			store.tokens.Delete(key)
		}
		if claims.AccountId == accountId {
			store.tokens.Delete(key)
		}
		return true
	})
}

// IsValidToken 用于验证Token是否有效
func (store *TokenStore) IsValidToken(token string) (*util.Claims, bool) {
	secret, ok := store.tokens.Load(token)
	if !ok {
		logging.Logger().Debugf("Token not found in cache: %s", token)
		return nil, false
	}
	claims, err := util.ValidateToken(token, secret.(string))
	if err != nil {
		store.tokens.Delete(token)
		return nil, false
	}

	expiresAt := claims.ExpiresAt
	if time.Now().After(time.Unix(expiresAt, 0)) {
		store.tokens.Delete(token)
		return nil, false
	}

	return claims, true
}

// ExpireTokens 清理过期的Token，可以定期运行
func (store *TokenStore) ExpireTokens() {
	store.tokens.Range(func(key, value interface{}) bool {
		_, ok := store.tokens.Load(key)
		if !ok {
			return false
		}
		claims, err := util.ValidateToken(key.(string), value.(string))
		if err != nil {
			store.tokens.Delete(key)
		}
		expiresAt := claims.ExpiresAt
		if time.Now().After(time.Unix(expiresAt, 0)) {
			store.tokens.Delete(key)
		}
		return true
	})
}

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

		claims, validToken := TokenStoreCache.IsValidToken(token)
		logging.Logger().Debugf("Valid token result: %t", validToken)
		if !validToken {
			panic(errors.NewUnauthorizedError("Unauthorized"))
		}

		context.Set("account_id", claims.AccountId)
		context.Set("token", token)

		logging.Logger().Debugf("Verification authentication passed. Add attribute account_id: %d token: %s to context", claims.AccountId, token)

		// 放行
		context.Next()
	}
}
