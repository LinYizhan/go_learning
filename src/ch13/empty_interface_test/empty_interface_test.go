package empty_interface_test

import (
	"fmt"
	"testing"
)

// func DoSomething(p interface{}) {
// 	if i, ok := p.(int); ok {
// 		fmt.Println("Integer", i)
// 		return
// 	}
//
// 	fmt.Print("unknown type")
// }
func DoSomething(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	default:
		fmt.Print("unknown type")
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
}
