package cache

import (
	"github.com/farseer-go/collections"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/mapper"
	"runtime"
	"time"
)

type cacheManage[TEntity any] struct {
	// 缓存KEY
	key string
	// 获取缓存实现
	cache ICache
	// 集合的数据的来源
	listSourceFn func() collections.List[TEntity]
	// item的数据来源
	// bool: isExists
	itemSourceFn func(cacheId any) (TEntity, bool)
	// item项为nil时，是否重新加载整个集合
	itemNullToLoadALl bool
}

// SetListSource 集合数据不存在时，则通过getListSourceFn获取
func (receiver *cacheManage[TEntity]) SetListSource(getListSourceFn func() collections.List[TEntity]) {
	receiver.listSourceFn = getListSourceFn
}

// SetItemSource 元素不存在时，则通过getItemSourceFn获取
func (receiver *cacheManage[TEntity]) SetItemSource(getItemSourceFn func(cacheId any) (TEntity, bool)) {
	receiver.itemSourceFn = getItemSourceFn
}

// EnableItemNullToLoadAll 元素不存在时，自动读取集合数据源
func (receiver *cacheManage[TEntity]) EnableItemNullToLoadAll() {
	receiver.itemNullToLoadALl = true
}

// Get 获取缓存数据
func (receiver *cacheManage[TEntity]) Get() collections.List[TEntity] {
	lst := receiver.cache.Get()
	// 如果数据为空，则调用数据源
	if lst.IsEmpty() && receiver.listSourceFn != nil {
		lstSource := receiver.listSourceFn()
		receiver.Set(lstSource.ToArray()...)
		return lstSource
	}
	return mapper.ToList[TEntity](lst)
}

// Single 获取单个对象
func (receiver *cacheManage[TEntity]) Single() TEntity {
	lst := receiver.Get()
	return mapper.ToList[TEntity](lst).First()
}

// GetItem 从集合中获取指定cacheId的元素
func (receiver *cacheManage[TEntity]) GetItem(cacheId any) (TEntity, bool) {
	item := receiver.cache.GetItem(cacheId)
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
				id := receiver.cache.GetUniqueId(source)
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

// GetItems 从集合中获取多个cacheId的元素
func (receiver *cacheManage[TEntity]) GetItems(cacheIds ...any) collections.List[TEntity] {
	items := receiver.cache.GetItems(cacheIds)
	lst := collections.NewList[TEntity]()
	for _, item := range items.ToArray() {
		lst.Add(*item.(*TEntity))
	}
	return lst
}

// Set 缓存整个集合，将覆盖原有集合（如果有数据）
func (receiver *cacheManage[TEntity]) Set(val ...TEntity) {
	if len(val) == 0 {
		return
	}

	lst := collections.NewListAny()
	for _, entity := range val {
		lst.Add(entity)
	}
	receiver.cache.Set(lst)
}

// SaveItem 更新item数据到集合
func (receiver *cacheManage[TEntity]) SaveItem(newVal TEntity) {
	receiver.cache.SaveItem(newVal)
}

// Remove 移除集合中的item数据
func (receiver *cacheManage[TEntity]) Remove(cacheId any) {
	receiver.cache.Remove(cacheId)
}

// Clear 清空数据
func (receiver *cacheManage[TEntity]) Clear() {
	receiver.cache.Clear()
}

// ExistsKey 缓存集合是否存在：如果没初始过Key，或者Key缓存已失效，都会返回false
func (receiver *cacheManage[TEntity]) ExistsKey() bool {
	return receiver.cache.ExistsKey()
}

// ExistsItem 缓存是否存在
func (receiver *cacheManage[TEntity]) ExistsItem(cacheId any) bool {
	return receiver.cache.ExistsItem(cacheId)
}

// Count 获取集合内的数量
func (receiver *cacheManage[TEntity]) Count() int {
	count := receiver.cache.Count()
	if count == 0 && !receiver.ExistsKey() {
		// key不存在，需要重新读一遍Get
		return receiver.Get().Count()
	}
	return count
}

// SetSyncSource 设置将缓存的数据同步到你需要的位置，比如同步到数据库
func (receiver *cacheManage[TEntity]) SetSyncSource(duration time.Duration, f func(val TEntity)) {
	go func() {
		timer := time.NewTimer(duration)
		for {
			timer.Reset(duration)
			select {
			case <-timer.C:
				lst := receiver.Get()
				for i := 0; i < lst.Count(); i++ {
					f(lst.Index(i))
					runtime.Gosched()
				}
			case <-fs.Context.Done():
				return
			}
		}
	}()
}

// SetClearSource 设置清理缓存中的数据
func (receiver *cacheManage[TEntity]) SetClearSource(duration time.Duration, f func(val TEntity) bool) {
	go func() {
		timer := time.NewTimer(duration)
		for {
			timer.Reset(duration)
			select {
			case <-timer.C:
				lst := receiver.Get()
				for i := 0; i < lst.Count(); i++ {
					if f(lst.Index(i)) {
						receiver.Remove(receiver.cache.GetUniqueId(lst.Index(i)))
					}
					runtime.Gosched()
				}
			case <-fs.Context.Done():
				return
			}
		}
	}()
}
