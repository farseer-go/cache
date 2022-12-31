package cache

import "github.com/farseer-go/collections"

type ICacheManage[TEntity any] interface {
	SetListSource(getListSourceFn func() collections.List[TEntity])  // 集合数据不存在时，则通过getListSourceFn获取
	SetItemSource(getItemSourceFn func(cacheId any) (TEntity, bool)) // 元素不存在时，则通过getItemSourceFn获取
	EnableItemNullToLoadALl()                                        // 元素不存在时，自动读取集合数据源
	Get() collections.List[TEntity]                                  // 获取缓存数据
	Single() TEntity                                                 // 获取单个对象
	GetItem(cacheId any) (TEntity, bool)                             // 从集合中获取指定cacheId的元素
	Set(val ...TEntity)                                              // 缓存整个集合，将覆盖原有集合（如果有数据）
	SaveItem(newVal TEntity)                                         // 更新item数据到集合
	Remove(cacheId string)                                           // 移除集合中的item数据
	Clear()                                                          // 清空数据
	ExistsKey() bool                                                 // 缓存集合是否存在：如果没初始过Key，或者Key缓存已失效，都会返回false
	ExistsItem(cacheId string) bool                                  // 缓存是否存在
	Count() int                                                      // 获取集合内的数量
}
