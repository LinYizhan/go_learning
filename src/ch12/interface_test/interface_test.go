package interface_test

import "testing"

type Programer interface {
	WriteHelloWorld() string
}

type GoProgramer struct {
}

func (g *GoProgramer) WriteHelloWorld() string {
	return "Hello World"
}

func TestInterface(t *testing.T) {
	var p Programer
	p = new(GoProgramer)
	t.Log(p.WriteHelloWorld())
}
