package singleton

import (
	"sync/atomic"
)

// 全局唯一类
type IdGenerator struct {
	id int64
}

var (
	idGeneratorInstance *IdGenerator
)

func GetIdGenerator() *IdGenerator {
	once.Do(func() {
		idGeneratorInstance = &IdGenerator{id: 0}
	})
	return idGeneratorInstance
}

func (g *IdGenerator) GetId() int64 {
	return atomic.AddInt64(&g.id, 1)
}
