package cache

import (
	"github.com/farseer-go/cache/eumCacheStoreType"
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

type po struct {
	Name string
	Age  int
}

func TestSetProfilesInMemory(t *testing.T) {
	modules.StartModules(Module{})
	SetProfilesInMemory[po]("test", "Name", 5)
	cacheMange := GetCacheManage[po]("test")

	assert.Equal(t, cacheMange.Key, "test")
	assert.Equal(t, cacheMange.UniqueField, "Name")
	assert.Nil(t, cacheMange.Cache)
	assert.Equal(t, cacheMange.MemoryExpiry, time.Duration(5))
	assert.Equal(t, cacheMange.CacheStoreType, eumCacheStoreType.Memory)
	assert.Equal(t, cacheMange.ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheMange.RedisConfigName, "")
	assert.Equal(t, cacheMange.RedisExpiry, time.Duration(0))
}

func TestSetProfilesInRedis(t *testing.T) {
	modules.StartModules(Module{})
	SetProfilesInRedis[po]("test1", "default", "Name", 5)
	cacheMange := GetCacheManage[po]("test1")

	assert.Equal(t, cacheMange.Key, "test1")
	assert.Equal(t, cacheMange.UniqueField, "Name")
	assert.Nil(t, cacheMange.Cache)
	assert.Equal(t, cacheMange.MemoryExpiry, time.Duration(0))
	assert.Equal(t, cacheMange.CacheStoreType, eumCacheStoreType.Redis)
	assert.Equal(t, cacheMange.ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheMange.RedisConfigName, "default")
	assert.Equal(t, cacheMange.RedisExpiry, time.Duration(5))
}

func TestSetProfilesInMemoryAndRedis(t *testing.T) {
	modules.StartModules(Module{})
	SetProfilesInMemoryAndRedis[po]("test2", "default", "Name", 5, 6)
	cacheMange := GetCacheManage[po]("test2")

	assert.Equal(t, cacheMange.Key, "test2")
	assert.Equal(t, cacheMange.UniqueField, "Name")
	assert.Nil(t, cacheMange.Cache)
	assert.Equal(t, cacheMange.MemoryExpiry, time.Duration(6))
	assert.Equal(t, cacheMange.CacheStoreType, eumCacheStoreType.MemoryAndRedis)
	assert.Equal(t, cacheMange.ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheMange.RedisConfigName, "default")
	assert.Equal(t, cacheMange.RedisExpiry, time.Duration(5))
}

func TestSetSingleProfilesInMemory(t *testing.T) {
	modules.StartModules(Module{})
	SetSingleProfilesInMemory[po]("test3", 5)
	cacheMange := GetCacheManage[po]("test3")

	assert.Equal(t, cacheMange.Key, "test3")
	assert.Equal(t, cacheMange.UniqueField, "")
	assert.Nil(t, cacheMange.Cache)
	assert.Equal(t, cacheMange.MemoryExpiry, time.Duration(5))
	assert.Equal(t, cacheMange.CacheStoreType, eumCacheStoreType.Memory)
	assert.Equal(t, cacheMange.ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheMange.RedisConfigName, "")
	assert.Equal(t, cacheMange.RedisExpiry, time.Duration(0))
}

func TestSetSingleProfilesInRedis(t *testing.T) {
	modules.StartModules(Module{})
	SetSingleProfilesInRedis[po]("test4", "default", 5)
	cacheMange := GetCacheManage[po]("test4")

	assert.Equal(t, cacheMange.Key, "test4")
	assert.Equal(t, cacheMange.UniqueField, "")
	assert.Nil(t, cacheMange.Cache)
	assert.Equal(t, cacheMange.MemoryExpiry, time.Duration(0))
	assert.Equal(t, cacheMange.CacheStoreType, eumCacheStoreType.Redis)
	assert.Equal(t, cacheMange.ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheMange.RedisConfigName, "default")
	assert.Equal(t, cacheMange.RedisExpiry, time.Duration(5))
}

func TestSetSingleProfilesInMemoryAndRedis(t *testing.T) {
	modules.StartModules(Module{})
	SetSingleProfilesInMemoryAndRedis[po]("test5", "default", 6, 7)
	cacheMange := GetCacheManage[po]("test5")

	assert.Equal(t, cacheMange.Key, "test5")
	assert.Equal(t, cacheMange.UniqueField, "")
	assert.Nil(t, cacheMange.Cache)
	assert.Equal(t, cacheMange.MemoryExpiry, time.Duration(7))
	assert.Equal(t, cacheMange.CacheStoreType, eumCacheStoreType.MemoryAndRedis)
	assert.Equal(t, cacheMange.ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheMange.RedisConfigName, "default")
	assert.Equal(t, cacheMange.RedisExpiry, time.Duration(6))
}
