package cache

import (
	"github.com/farseer-go/cache/eumExpiryType"
	"time"
)

// Option 缓存选项
type Option func(*Op)

// Op 缓存选项
type Op struct {
	ExpiryType eumExpiryType.Enum // 过期策略
	Expiry     time.Duration      // 缓存失效时间
}

// AbsoluteExpiration 绝对时间，到期自动移除
func (receiver *Op) AbsoluteExpiration(expiry time.Duration) {
	receiver.ExpiryType = eumExpiryType.AbsoluteExpiration
	receiver.Expiry = expiry
}

// SlidingExpiration 任意一次访问都会重置过期时间
func (receiver *Op) SlidingExpiration(expiry time.Duration) {
	receiver.ExpiryType = eumExpiryType.SlidingExpiration
	receiver.Expiry = expiry
}
