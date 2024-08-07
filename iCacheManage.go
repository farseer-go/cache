package cache

import (
	"github.com/farseer-go/collections"
	"time"
)

type ICacheManage[TEntity any] interface {
	// SetListSource 集合数据不存在时，则通过getListSourceFn获取
	SetListSource(getListSourceFn func() collections.List[TEntity])
	// SetItemSource 元素不存在时，则通过getItemSourceFn获取
	SetItemSource(getItemSourceFn func(cacheId any) (TEntity, bool))
	// EnableItemNullToLoadAll 元素不存在时，自动读取集合数据源
	EnableItemNullToLoadAll()
	// Get 获取缓存数据
	Get() collections.List[TEntity]
	// Single 获取单个对象
	Single() TEntity
	// GetItem 从集合中获取指定cacheId的元素
	GetItem(cacheId any) (TEntity, bool)
	// GetItems 从集合中获取多个cacheId的元素
	GetItems(cacheIds ...any) collections.List[TEntity]
	// Set 缓存整个集合，将覆盖原有集合（如果有数据）
	Set(val ...TEntity)
	// SaveItem 更新item数据到集合
	SaveItem(newVal TEntity)
	// Remove 移除集合中的item数据
	Remove(cacheId any)
	// Clear 清空数据
	Clear()
	// ExistsKey 缓存集合是否存在：如果没初始过Key，或者Key缓存已失效，都会返回false
	ExistsKey() bool
	// ExistsItem 缓存是否存在
	ExistsItem(cacheId any) bool
	// Count 获取集合内的数量
	Count() int
	// SetSyncSource 设置定义将缓存的数据同步到你需要的位置，比如同步到数据库
	SetSyncSource(duration time.Duration, f func(val TEntity))
	// SetClearSource 设置定义清理缓存中的数据
	SetClearSource(duration time.Duration, f func(val TEntity) bool)
}
