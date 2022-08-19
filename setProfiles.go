package cache

import (
	"github.com/farseer-go/cache/eumCacheStoreType"
	"github.com/farseer-go/fs/container"
	"github.com/farseer-go/fs/exception"
	"reflect"
	"time"
)

// SetProfilesInMemory 设置内存缓存（集合）
func SetProfilesInMemory[TEntity any](key string, uniqueField string, memoryExpiry time.Duration) {
	if uniqueField == "" {
		exception.ThrowRefuseException("缓存集合数据时，需要设置UniqueField字段")
	}
	var entity TEntity
	entityType := reflect.TypeOf(entity)
	_, isExists := entityType.FieldByName(uniqueField)
	if !isExists {
		exception.ThrowRefuseException(uniqueField + "字段，在缓存集合中不存在")
	}

	cacheConfigure[key] = CacheManage[TEntity]{
		CacheKey: CacheKey{
			Key:            key,
			CacheStoreType: eumCacheStoreType.Memory,
			MemoryExpiry:   memoryExpiry,
			UniqueField:    uniqueField,
			ItemType:       entityType,
			Cache:          container.ResolveName[ICache]("memory"),
		},
	}
}

// SetProfilesInRedis 设置Redis缓存（集合）
func SetProfilesInRedis[TEntity any](key string, redisConfigName string, uniqueField string, redisExpiry time.Duration) {
	if uniqueField == "" {
		exception.ThrowRefuseException("缓存集合数据时，需要设置UniqueField字段")
	}
	var entity TEntity
	entityType := reflect.TypeOf(entity)
	_, isExists := entityType.FieldByName(uniqueField)
	if !isExists {
		exception.ThrowRefuseException(uniqueField + "字段，在缓存集合中不存在")
	}

	cacheConfigure[key] = CacheManage[TEntity]{
		CacheKey: CacheKey{
			Key:             key,
			CacheStoreType:  eumCacheStoreType.Redis,
			RedisExpiry:     redisExpiry,
			ItemType:        entityType,
			UniqueField:     uniqueField,
			RedisConfigName: redisConfigName,
			Cache:           container.ResolveName[ICache]("redis"),
		},
	}
}

// SetProfilesInMemoryAndRedis 设置内存-Redis缓存（集合）
func SetProfilesInMemoryAndRedis[TEntity any](key string, redisConfigName string, uniqueField string, redisExpiry time.Duration, memoryExpiry time.Duration) {
	if uniqueField == "" {
		exception.ThrowRefuseException("缓存集合数据时，需要设置UniqueField字段")
	}
	var entity TEntity
	entityType := reflect.TypeOf(entity)
	_, isExists := entityType.FieldByName(uniqueField)
	if !isExists {
		exception.ThrowRefuseException(uniqueField + "字段，在缓存集合中不存在")
	}

	cacheConfigure[key] = CacheManage[TEntity]{
		CacheKey: CacheKey{
			Key:             key,
			CacheStoreType:  eumCacheStoreType.MemoryAndRedis,
			RedisExpiry:     redisExpiry,
			MemoryExpiry:    memoryExpiry,
			UniqueField:     uniqueField,
			ItemType:        entityType,
			RedisConfigName: redisConfigName,
			Cache:           container.ResolveName[ICache]("memoryAndRedis"),
		},
	}
}

// SetSingleProfilesInMemory 设置内存缓存（缓存单个对象）
func SetSingleProfilesInMemory[TEntity any](key string, memoryExpiry time.Duration) {
	var entity TEntity
	entityType := reflect.TypeOf(entity)
	cacheConfigure[key] = CacheManage[TEntity]{
		CacheKey: CacheKey{
			Key:            key,
			CacheStoreType: eumCacheStoreType.Memory,
			MemoryExpiry:   memoryExpiry,
			ItemType:       entityType,
			Cache:          container.ResolveName[ICache]("memory"),
		},
	}
}

// SetSingleProfilesInRedis 设置Redis缓存（缓存单个对象）
func SetSingleProfilesInRedis[TEntity any](key string, redisConfigName string, redisExpiry time.Duration) {
	var entity TEntity
	entityType := reflect.TypeOf(entity)
	cacheConfigure[key] = CacheManage[TEntity]{
		CacheKey: CacheKey{
			Key:             key,
			CacheStoreType:  eumCacheStoreType.Redis,
			RedisExpiry:     redisExpiry,
			ItemType:        entityType,
			RedisConfigName: redisConfigName,
			Cache:           container.ResolveName[ICache]("redis"),
		},
	}
}

// SetSingleProfilesInMemoryAndRedis 设置内存-Redis缓存（缓存单个对象）
func SetSingleProfilesInMemoryAndRedis[TEntity any](key string, redisConfigName string, redisExpiry time.Duration, memoryExpiry time.Duration) {
	var entity TEntity
	entityType := reflect.TypeOf(entity)
	cacheConfigure[key] = CacheManage[TEntity]{
		CacheKey: CacheKey{
			Key:             key,
			CacheStoreType:  eumCacheStoreType.MemoryAndRedis,
			RedisExpiry:     redisExpiry,
			MemoryExpiry:    memoryExpiry,
			ItemType:        entityType,
			RedisConfigName: redisConfigName,
			Cache:           container.ResolveName[ICache]("memoryAndRedis"),
		},
	}
}
