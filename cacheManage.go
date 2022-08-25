package cache

import (
	"github.com/farseer-go/cache/eumCacheStoreType"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/mapper"
	"reflect"
	"time"
)

// value=CacheManage
var cacheConfigure map[string]any

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

type CacheManage[TEntity any] struct {
	CacheKey
	source            func() collections.List[TEntity]
	itemNullToLoadALl bool
}

// GetCacheManage 获取CacheKey
func GetCacheManage[TEntity any](key string) CacheManage[TEntity] {
	cacheKey, exists := cacheConfigure[key]
	if !exists {
		panic(key + "不存在，要使用Cache缓存，需提前初始化")
	}
	return cacheKey.(CacheManage[TEntity])
}

// SetSource 设置数据源
func (receiver *CacheManage[TEntity]) SetSource(getSourceFn func() collections.List[TEntity]) {
	receiver.source = getSourceFn
}

// EnableItemNullToLoadALl 元素不存在时，自动读取数据源
func (receiver *CacheManage[TEntity]) EnableItemNullToLoadALl() {
	receiver.itemNullToLoadALl = true
}

// Get 获取缓存数据
func (receiver CacheManage[TEntity]) Get() collections.List[TEntity] {
	lst := receiver.Cache.Get(receiver.CacheKey)
	// 如果数据为空，则调用数据源
	if lst.IsEmpty() {
		lstSource := receiver.source()
		receiver.Set(lstSource.ToArray()...)
		return lstSource
	}
	return mapper.ToList[TEntity](lst)
}

// Single 获取单个对象
func (receiver CacheManage[TEntity]) Single() TEntity {
	lst := receiver.Cache.Get(receiver.CacheKey)
	return mapper.ToList[TEntity](lst).First()
}

// GetItem 从集合中获取指定cacheId的元素
func (receiver CacheManage[TEntity]) GetItem(cacheId any) (TEntity, bool) {
	item := receiver.Cache.GetItem(receiver.CacheKey, parse.Convert(cacheId, ""))
	if item == nil {
		// 元素不存在时，自动读取数据源
		if receiver.itemNullToLoadALl {
			lstSource := receiver.source()
			receiver.Set(lstSource.ToArray()...)
			// 从列表中读取元素
			for _, source := range lstSource.ToArray() {
				id := receiver.GetUniqueId(source)
				if cacheId == id {
					return source, true
				}
			}
		}
		var entity TEntity
		return entity, false
	}
	return item.(TEntity), true
}

// Set 保存缓存
func (receiver CacheManage[TEntity]) Set(val ...TEntity) {
	lst := collections.NewListAny()
	for _, entity := range val {
		lst.Add(entity)
	}
	receiver.Cache.Set(receiver.CacheKey, lst)
}

// SaveItem 更新缓存
func (receiver CacheManage[TEntity]) SaveItem(newVal TEntity) {
	receiver.Cache.SaveItem(receiver.CacheKey, newVal)
}

// Remove 移除缓存
func (receiver CacheManage[TEntity]) Remove(cacheId string) {
	receiver.Cache.Remove(receiver.CacheKey, cacheId)
}

// Clear 清空缓存
func (receiver CacheManage[TEntity]) Clear() {
	receiver.Cache.Clear(receiver.CacheKey)
}

// Exists 缓存是否存在
func (receiver CacheManage[TEntity]) Exists() bool {
	return receiver.Cache.ExistsKey(receiver.CacheKey)
}

// Count 数据集合的数量
func (receiver CacheManage[TEntity]) Count() int {
	if !receiver.Exists() {
		return 0
	}
	return receiver.Cache.Count(receiver.CacheKey)
}
