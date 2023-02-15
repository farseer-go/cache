package cache

import (
	"github.com/farseer-go/collections"
)

// ICache 缓存操作
type ICache interface {
	// Get 从本地内存中获取
	Get() collections.ListAny
	// GetItem 获取集合中的item
	GetItem(cacheId any) any
	// Set 整个数据保存到缓存
	Set(val collections.ListAny)
	// SaveItem 添加或修改集合的Item
	SaveItem(newVal any)
	// Remove 移除缓存
	Remove(cacheId any)
	// Clear 清空缓存
	Clear()
	// Count 获取集合数量
	Count() int
	// ExistsItem 指定数据是否存在
	ExistsItem(cacheId any) bool
	// ExistsKey 缓存是否存在
	ExistsKey() bool
	// GetUniqueId 获取唯一字段数据
	GetUniqueId(item any) string
}
