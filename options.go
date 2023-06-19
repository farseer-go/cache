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
