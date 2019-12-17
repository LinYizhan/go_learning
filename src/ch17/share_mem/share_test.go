package share_men

import (
	"sync"
	"testing"
	"time"
)

func TestShareParam(t *testing.T) {
	count := 0
	for i := 0; i < 5000; i++ {
		go func() {
			count++
		}()
	}

	time.Sleep(time.Millisecond * 50)
	t.Log(count)
}

func TestShareParamWaitGroup(t *testing.T) {
	var mux sync.Mutex
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mux.Unlock()
			}()
			mux.Lock()
			count++
			wg.Done()
		}()
	}

	// time.Sleep(time.Millisecond * 50)
	wg.Wait()
	t.Log(count)
}
