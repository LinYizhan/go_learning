package error_test

import (
	"errors"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	if n < 1 || n > 100 {
		return nil, errors.New("n should be in [2, 100]")
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}

	return fibList, nil
}

func TestError(t *testing.T) {
	if v, err := GetFibonacci(10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}
}
