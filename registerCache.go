package cache

import (
	"reflect"

	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/exception"
)

// RegisterCacheModule 注册缓存
func RegisterCacheModule[TEntity any](key string, cacheStoreType string, uniqueField string, cache ICache) ICacheManage[TEntity] {
	if uniqueField == "" {
		exception.ThrowRefuseException("缓存集合数据时，需要设置UniqueField字段")
	}
	var entity TEntity
	entityType := reflect.TypeOf(entity)
	_, isExists := entityType.FieldByName(uniqueField)
	if !isExists {
		exception.ThrowRefuseException(uniqueField + "字段，在缓存集合中不存在")
	}

	var cacheManage ICacheManage[TEntity] = &cacheManage[TEntity]{
		key:   key,
		cache: cache,
	}
	if !container.IsRegister[ICacheManage[TEntity]](key) {
		container.RegisterInstance(cacheManage, key)
	}
	return cacheManage
}
