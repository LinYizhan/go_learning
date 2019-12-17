package pool_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create new obj")
			return 10
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(100)
	v1, ok := pool.Get().(int)
	fmt.Println(v1, ok)
}
