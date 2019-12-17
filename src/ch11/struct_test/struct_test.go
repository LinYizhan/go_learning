package struct_test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

// func (e Employee) String() string {
// 	return fmt.Sprintf("ID=%s-NAME=%s-AGE=%d", e.Id, e.Name, e.Age)
// }

func (e *Employee) String() string {
	return fmt.Sprintf("ID=%s-NAME=%s-AGE=%d", e.Id, e.Name, e.Age)
}

func TestStruct(t *testing.T) {
	e := &Employee{"0", "Bo", 14}
	t.Log(e.String())
}
