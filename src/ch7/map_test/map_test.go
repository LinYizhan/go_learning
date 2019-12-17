package map_test

import "testing"

func TestMap(t *testing.T) {
	a := map[int]int{}
	a[2] = 1
	if v, ok := a[2]; ok {
		t.Logf("key is exist, value = %d", v)
	} else {
		t.Log("key is not exist")
	}
}

func TestTravelMap(t *testing.T) {
	a := map[int]int{1: 1, 2: 4, 3: 5}
	for v, k := range a {
		t.Log(v, k)
	}
}
