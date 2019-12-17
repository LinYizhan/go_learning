package first_response_test

import (
	"fmt"
	"testing"
	"time"
)

func RunTask(i int) string {
	time.Sleep(time.Millisecond * 100)
	return fmt.Sprintf("the result is from id %d", i)
}

func FirstResponse() string {
	number := 10
	ch := make(chan string, number)
	for i := 0; i < number; i++ {
		go func(i int) {
			ret := RunTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {

	t.Log(FirstResponse())
}
