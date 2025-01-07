package test

import (
	"fmt"
	"jmy/go-design-pattern/creational-pattern/singleton"
	"sync"
	"testing"
)

func Test1(t *testing.T) {
	logger := singleton.GetLogger()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			logger.Log(fmt.Sprintf("This is number %d test log message.", i))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func Test2(t *testing.T) {
	generator := singleton.GetIdGenerator()
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			id := generator.GetId()
			t.Logf("i:%d,id:%d\n", i, id)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
