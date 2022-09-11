## cache
缓存操作

## What are the functions?
* cache
  * func
      * GetCacheManage （获取缓存对象）
      * SetProfilesInMemory （设置内存缓存（集合））
      * SetProfilesInRedis （设置Redis缓存（集合））
      * SetProfilesInMemoryAndRedis （设置内存、Redis缓存（集合））
  * struct
    * CacheManage（缓存对象）
      * Get（获取缓存集合）
      * Single（获取单个缓存对象）
      * GetItem（从集合中获取指定cacheId的元素）
      * Set（保存缓存）
      * SaveItem（更新缓存）
      * Remove（移除缓存）
      * Clear（清空缓存）
      * ExistsKey（缓存是否存在）
      * Count（数据集合的数量）
      * SetListSource（设置集合数据源）
      * SetItemSource（设置Item数据源）
      * EnableItemNullToLoadALl（元素不存在时，自动读取集合数据源）

## Getting Started
```go
type po struct {
    Name string
    Age  int
}
```

### 本地缓存使用
**注意：你需要在`startupModule`中依赖`cacheMemory.Module`模块**
```go
// 配置缓存
// arg1: Key name
// arg2: UniqueField
// arg3：MemoryExpiry
cache.SetProfilesInMemory[po]("test", "Name", 60*time.Second)
// 获取key=test的缓存对象
cacheManage := cache.GetCacheManage[po]("test")

// 设置：集合数据不存在时，则通过func获取
cacheManage.SetListSource(func() collections.List[taskGroup.DomainObject] {
    var lst collections.List[taskGroup.DomainObject]
    repository.TaskGroup.ToList().MapToList(&lst)
    return lst
})

// 设置：元素不存在时，则通过func获取
cacheManage.SetItemSource(func(cacheId any) (taskGroup.DomainObject, bool) {
    po := repository.TaskGroup.Where("Id = ?", cacheId).ToEntity()
    if po.Id > 0 {
        return mapper.Single[taskGroup.DomainObject](&po), true
    }
    var do taskGroup.DomainObject
      return do, false
})

// 元素不存在时，自动读取集合数据源（需调用了SetListSource方法）
cacheManage.EnableItemNullToLoadALl()

// 设置缓存
cacheManage.Set(po{Name: "steden", Age: 18}, po{Name: "steden2", Age: 19})
// 获取整个集合
cacheManage.Get()
// 获取Name=steden 的缓存item
cacheManage.GetItem("steden")
// 更新Name=steden的缓存item
cacheManage.SaveItem(po{Name: "steden", Age: 99})
// 移除Name=steden缓存item
cacheManage.Remove("steden")
// 清除缓存
cacheManage.Clear()
// 检查缓存是否存在
cacheManage.ExistsKey()
// 获取缓存的集合数量
cacheManage.Count()
```