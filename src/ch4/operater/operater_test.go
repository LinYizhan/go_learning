package operater

import "testing"

func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{2, 2, 3, 4}
	//c := [...]int{0, 1, 2, 3, 4}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c)
	t.Log(a == d)
}
