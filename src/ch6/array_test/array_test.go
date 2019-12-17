package array_test

import "testing"

func TestArray(t *testing.T) {
	a := [...]int{1, 2, 3}
	for _, e := range a {
		t.Log(e)
	}

	s := make([]int, 3, 5)
	t.Log(len(s), cap(s))

	var v string = "sss"
	t.Log(v)
}
