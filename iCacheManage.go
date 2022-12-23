package cache

import "github.com/farseer-go/collections"

type ICacheManage[TEntity any] interface {
	SetListSource(getListSourceFn func() collections.List[TEntity])
	SetItemSource(getItemSourceFn func(cacheId any) (TEntity, bool))
	EnableItemNullToLoadALl()
	Get() collections.List[TEntity]
	Single() TEntity
	GetItem(cacheId any) (TEntity, bool)
	Set(val ...TEntity)
	SaveItem(newVal TEntity)
	Remove(cacheId string)
	Clear()
	ExistsKey() bool
	ExistsItem(cacheId string) bool
	Count() int
	Key() CacheKey
}
