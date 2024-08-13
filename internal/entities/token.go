package entities

import "time"

type TokenType string

const (
	// AccountTokenType 管理员账户
	AccountTokenType = "account"
	// MerchantsTokenType 商户token账户
	MerchantsTokenType = "merchants"
)

type Token struct {
	Entity
	// 目标id 取决 Type 类型决定该数据值的指向id
	TargetId int64 `db:"target_id"`
	// token目标类型
	Type TokenType `db:"type"`
	// 访问token
	AccessToken string `db:"access_token"`
	// 刷新token
	RefreshToken string `db:"refresh_token"`
	// token国企时间，如果为 nil 则永不过期
	ExpiresAt *time.Time `db:"expires_at"`
}
