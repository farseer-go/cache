package cache

import "github.com/farseer-go/collections"

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
	// Set 缓存整个集合，将覆盖原有集合（如果有数据）
	Set(val ...TEntity)
	// SaveItem 更新item数据到集合
	SaveItem(newVal TEntity)
	// Remove 移除集合中的item数据
	Remove(cacheId string)
	// Clear 清空数据
	Clear()
	// ExistsKey 缓存集合是否存在：如果没初始过Key，或者Key缓存已失效，都会返回false
	ExistsKey() bool
	// ExistsItem 缓存是否存在
	ExistsItem(cacheId string) bool
	// Count 获取集合内的数量
	Count() int
}
