package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var str string

	str = "中"

	c := []rune(str)

	t.Logf("unicode = %x", c[0])
	t.Logf("utf8 = %c", c[0])
}

func TestStringFunc(t *testing.T) {

	s := "A,B,C"
	strs := strings.Split(s, ",")
	for _, str := range strs {
		t.Log(str)
	}
	t.Log(strings.Join(strs, "--"))

	as := strconv.Itoa(10)
	t.Log(as + "sss")
	if i, error := strconv.Atoi("10"); error == nil {
		t.Log(i + 10)
	}
}
