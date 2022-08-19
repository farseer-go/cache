package cache

import (
	"github.com/farseer-go/fs/modules"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetProfilesInMemory(t *testing.T) {
	modules.StartModules(Module{})
	type po struct {
		Name string
		Age  int
	}
	SetProfilesInMemory[po]("test", "Name", 0)
	profiles, exists := cacheConfigure["test"]
	assert.True(t, exists)

	cacheMange := profiles.(CacheManage[po])

	assert.Equal(t, cacheMange.Key, "test")
	assert.Equal(t, cacheMange.UniqueField, "Name")
}
