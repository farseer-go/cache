package test

import (
	"github.com/farseer-go/cache"
	"github.com/farseer-go/cache/eumCacheStoreType"
	"github.com/farseer-go/fs"
	"github.com/farseer-go/fs/container"
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
	fs.Initialize[cache.Module]("unit test")
	assert.Panics(t, func() {
		cache.SetProfilesInMemory[po]("test", "", 5)
	})
	assert.Panics(t, func() {
		cache.SetProfilesInMemory[po]("test", "ClientName", 5)
	})

	cache.SetProfilesInMemory[po]("test", "Name", 5)

	cacheManage := container.Resolve[cache.ICacheManage[po]]("test")
	assert.Equal(t, cacheManage.Key().Key, "test")
	assert.Equal(t, cacheManage.Key().UniqueField, "Name")
	assert.Nil(t, cacheManage.Key().Cache)
	assert.Equal(t, cacheManage.Key().MemoryExpiry, time.Duration(5))
	assert.Equal(t, cacheManage.Key().CacheStoreType, eumCacheStoreType.Memory)
	assert.Equal(t, cacheManage.Key().ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheManage.Key().RedisConfigName, "")
	assert.Equal(t, cacheManage.Key().RedisExpiry, time.Duration(0))

}

func TestSetProfilesInRedis(t *testing.T) {
	fs.Initialize[cache.Module]("unit test")

	assert.Panics(t, func() {
		cache.SetProfilesInRedis[po]("test1", "default", "", 5)
	})
	assert.Panics(t, func() {
		cache.SetProfilesInRedis[po]("test1", "default", "ClientName", 5)
	})

	cache.SetProfilesInRedis[po]("test1", "default", "Name", 5)
	cacheManage := container.Resolve[cache.ICacheManage[po]]("test1")

	assert.Equal(t, cacheManage.Key().Key, "test1")
	assert.Equal(t, cacheManage.Key().UniqueField, "Name")
	assert.Nil(t, cacheManage.Key().Cache)
	assert.Equal(t, cacheManage.Key().MemoryExpiry, time.Duration(0))
	assert.Equal(t, cacheManage.Key().CacheStoreType, eumCacheStoreType.Redis)
	assert.Equal(t, cacheManage.Key().ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheManage.Key().RedisConfigName, "default")
	assert.Equal(t, cacheManage.Key().RedisExpiry, time.Duration(5))
}

func TestSetProfilesInMemoryAndRedis(t *testing.T) {
	fs.Initialize[cache.Module]("unit test")

	assert.Panics(t, func() {
		cache.SetProfilesInMemoryAndRedis[po]("test2", "default", "", 5, 6)
	})
	assert.Panics(t, func() {
		cache.SetProfilesInMemoryAndRedis[po]("test2", "default", "ClientName", 5, 6)
	})

	cache.SetProfilesInMemoryAndRedis[po]("test2", "default", "Name", 5, 6)
	cacheManage := container.Resolve[cache.ICacheManage[po]]("test2")

	assert.Equal(t, cacheManage.Key().Key, "test2")
	assert.Equal(t, cacheManage.Key().UniqueField, "Name")
	assert.Nil(t, cacheManage.Key().Cache)
	assert.Equal(t, cacheManage.Key().MemoryExpiry, time.Duration(6))
	assert.Equal(t, cacheManage.Key().CacheStoreType, eumCacheStoreType.MemoryAndRedis)
	assert.Equal(t, cacheManage.Key().ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheManage.Key().RedisConfigName, "default")
	assert.Equal(t, cacheManage.Key().RedisExpiry, time.Duration(5))
}

/*
func TestSetSingleProfilesInMemory(t *testing.T) {
	fs.Initialize[Module]("unit test")
	SetSingleProfilesInMemory[po]("test3", 5)
	cacheManage := container.Resolve[cache.ICacheManage[po]]("test3")

	assert.Equal(t, cacheManage.Key, "test3")
	assert.Equal(t, cacheManage.UniqueField, "")
	assert.Nil(t, cacheManage.Cache)
	assert.Equal(t, cacheManage.MemoryExpiry, time.Duration(5))
	assert.Equal(t, cacheManage.CacheStoreType, eumCacheStoreType.Memory)
	assert.Equal(t, cacheManage.ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheManage.RedisConfigName, "")
	assert.Equal(t, cacheManage.RedisExpiry, time.Duration(0))
}

func TestSetSingleProfilesInRedis(t *testing.T) {
	fs.Initialize[Module]("unit test")
	SetSingleProfilesInRedis[po]("test4", "default", 5)
	cacheManage := container.Resolve[cache.ICacheManage[po]]("test4")

	assert.Equal(t, cacheManage.Key, "test4")
	assert.Equal(t, cacheManage.UniqueField, "")
	assert.Nil(t, cacheManage.Cache)
	assert.Equal(t, cacheManage.MemoryExpiry, time.Duration(0))
	assert.Equal(t, cacheManage.CacheStoreType, eumCacheStoreType.Redis)
	assert.Equal(t, cacheManage.ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheManage.RedisConfigName, "default")
	assert.Equal(t, cacheManage.RedisExpiry, time.Duration(5))
}

func TestSetSingleProfilesInMemoryAndRedis(t *testing.T) {
	fs.Initialize[Module]("unit test")
	SetSingleProfilesInMemoryAndRedis[po]("test5", "default", 6, 7)
	cacheManage := container.Resolve[cache.ICacheManage[po]]("test5")

	assert.Equal(t, cacheManage.Key, "test5")
	assert.Equal(t, cacheManage.UniqueField, "")
	assert.Nil(t, cacheManage.Cache)
	assert.Equal(t, cacheManage.MemoryExpiry, time.Duration(7))
	assert.Equal(t, cacheManage.CacheStoreType, eumCacheStoreType.MemoryAndRedis)
	assert.Equal(t, cacheManage.ItemType, reflect.TypeOf(po{}))
	assert.Equal(t, cacheManage.RedisConfigName, "default")
	assert.Equal(t, cacheManage.RedisExpiry, time.Duration(6))
}
*/
