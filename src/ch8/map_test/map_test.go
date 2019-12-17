package map_test

import "testing"

func TestFuncMap(t *testing.T) {
	a := map[int]func(op int) int{}

	a[1] = func(op int) int {
		return op
	}

	a[2] = func(op int) int {
		return op * op
	}

	t.Log(a[1](2), a[2](2))
}
