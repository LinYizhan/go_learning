package series_test

import (
	"testing"

	"ch14/error_test"
)

func TestSerise(t *testing.T) {
	t.Log(error_test.GetFibonacci(10))
}
