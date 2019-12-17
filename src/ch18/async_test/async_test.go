package async_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "is Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done")
}

func TestSyncService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retChannel := make(chan string, 1)

	go func() {
		ret := service()
		fmt.Println("return response")
		retChannel <- ret
		fmt.Println("service exited")
	}()
	return retChannel
}

func TestAsync(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}
