package cache

import (
	"github.com/farseer-go/cache/eumCacheStoreType"
	"github.com/farseer-go/fs/parse"
	"reflect"
	"time"
)

type CacheKey struct {
	// 缓存KEY
	Key string
	// 缓存策略（默认Memory模式）
	CacheStoreType eumCacheStoreType.Enum
	// 设置Redis缓存过期时间
	RedisExpiry time.Duration
	// 设置Memory缓存过期时间
	MemoryExpiry time.Duration
	// hash中的主键（唯一ID的字段名称）
	UniqueField string
	// Redis配置名称
	RedisConfigName string
	// 获取缓存实现
	Cache ICache
	// ItemType
	ItemType reflect.Type
}

// GetUniqueId 获取唯一字段数据
func (receiver CacheKey) GetUniqueId(item any) (T string) {
	val := reflect.ValueOf(item).FieldByName(receiver.UniqueField).Interface()
	return parse.Convert(val, "")
}
