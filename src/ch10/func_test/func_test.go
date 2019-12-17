package func_test

import (
	"fmt"
	"testing"
	"time"
)

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Print("time spent: ", time.Since(start).Seconds(), start)
		return ret
	}
}

func slowTime(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	tsf := timeSpent(slowTime)
	t.Log(tsf(10))
}

func clear() {
	fmt.Print("----")
}

func TestDefer(t *testing.T) {
	defer clear()
	fmt.Print("11111")
	panic("error")
}
