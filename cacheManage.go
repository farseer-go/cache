package cache

import (
	"github.com/farseer-go/cache/eumCacheStoreType"
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs/flog"
	"github.com/farseer-go/fs/parse"
	"github.com/farseer-go/fs/stopwatch"
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
	// 缓存key
	CacheKey
	// 集合的数据的来源
	listSourceFn func() collections.List[TEntity]
	// item的数据来源
	// bool: isExists
	itemSourceFn func(cacheId any) (TEntity, bool)
	// item项为nil时，是否重新加载整个集合
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

// SetListSource 集合数据不存在时，则通过getListSourceFn获取
func (receiver *CacheManage[TEntity]) SetListSource(getListSourceFn func() collections.List[TEntity]) {
	receiver.listSourceFn = getListSourceFn
}

// SetItemSource 元素不存在时，则通过getItemSourceFn获取
func (receiver *CacheManage[TEntity]) SetItemSource(getItemSourceFn func(cacheId any) (TEntity, bool)) {
	receiver.itemSourceFn = getItemSourceFn
}

// EnableItemNullToLoadALl 元素不存在时，自动读取集合数据源
func (receiver *CacheManage[TEntity]) EnableItemNullToLoadALl() {
	receiver.itemNullToLoadALl = true
}

// Get 获取缓存数据
func (receiver CacheManage[TEntity]) Get() collections.List[TEntity] {
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("cacheManage", ".Get：%s，耗时：%s", receiver.Key, sw.GetMillisecondsText())
	}()

	lst := receiver.Cache.Get(receiver.CacheKey)
	// 如果数据为空，则调用数据源
	if lst.IsEmpty() && receiver.listSourceFn != nil {
		lstSource := receiver.listSourceFn()
		receiver.Set(lstSource.ToArray()...)
		return lstSource
	}
	return mapper.ToList[TEntity](lst)
}

// Single 获取单个对象
func (receiver CacheManage[TEntity]) Single() TEntity {
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("cacheManage", ".Single：%s，耗时：%s", receiver.Key, sw.GetMillisecondsText())
	}()

	lst := receiver.Get()
	return mapper.ToList[TEntity](lst).First()
}

// GetItem 从集合中获取指定cacheId的元素
func (receiver CacheManage[TEntity]) GetItem(cacheId any) (TEntity, bool) {
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("cacheManage", ".GetItem：%s.%v，耗时：%s", receiver.Key, cacheId, sw.GetMillisecondsText())
	}()

	item := receiver.Cache.GetItem(receiver.CacheKey, parse.Convert(cacheId, ""))
	if item == nil {
		// 设置了单独的数据源时，则只读item数据源
		if receiver.itemSourceFn != nil {
			dbItem, isExists := receiver.itemSourceFn(cacheId)
			if isExists {
				receiver.SaveItem(dbItem)
				return dbItem, true
			}
		}
		// 元素不存在时，自动读取数据源
		if receiver.itemNullToLoadALl && receiver.listSourceFn != nil {
			lstSource := receiver.listSourceFn()
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
	if len(val) == 0 {
		return
	}
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("cacheManage", ".Set：%s，耗时：%s", receiver.Key, sw.GetMillisecondsText())
	}()

	lst := collections.NewListAny()
	for _, entity := range val {
		lst.Add(entity)
	}
	receiver.Cache.Set(receiver.CacheKey, lst)
}

// SaveItem 更新缓存
func (receiver CacheManage[TEntity]) SaveItem(newVal TEntity) {
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("cacheManage", ".SaveItem：%s，耗时：%s", receiver.Key, sw.GetMillisecondsText())
	}()

	receiver.Cache.SaveItem(receiver.CacheKey, newVal)
}

// Remove 移除缓存
func (receiver CacheManage[TEntity]) Remove(cacheId string) {
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("[cacheManage].Remove：%s.%v，耗时：%s", receiver.Key, cacheId, sw.GetMillisecondsText())
	}()

	receiver.Cache.Remove(receiver.CacheKey, cacheId)
}

// Clear 清空缓存
func (receiver CacheManage[TEntity]) Clear() {
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("cacheManage", ".Clear：%s，耗时：%s", receiver.Key, sw.GetMillisecondsText())
	}()

	receiver.Cache.Clear(receiver.CacheKey)
}

// ExistsKey 缓存是否存在
func (receiver CacheManage[TEntity]) ExistsKey() bool {
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("cacheManage", ".ExistsKey：%s，耗时：%s", receiver.Key, sw.GetMillisecondsText())
	}()

	return receiver.Cache.ExistsKey(receiver.CacheKey)
}

// ExistsItem 缓存是否存在
func (receiver CacheManage[TEntity]) ExistsItem(cacheId string) bool {
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("cacheManage", ".ExistsItem：%s，耗时：%s", receiver.Key, sw.GetMillisecondsText())
	}()

	return receiver.Cache.ExistsItem(receiver.CacheKey, cacheId)
}

// Count 数据集合的数量
func (receiver CacheManage[TEntity]) Count() int {
	sw := stopwatch.StartNew()
	defer func() {
		flog.AppInfof("cacheManage", ".Count：%s，耗时：%s", receiver.Key, sw.GetMillisecondsText())
	}()

	if !receiver.ExistsKey() {
		return 0
	}
	return receiver.Cache.Count(receiver.CacheKey)
}
