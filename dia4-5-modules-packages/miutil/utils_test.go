// miutil/utils_test.go
package miutil

import "testing"

func TestSumInt(t *testing.T) {
	if SumInt(2, 3) != 5 {
		t.Errorf("SumInt(2,3)=%d; want 5", SumInt(2, 3))
	}
}
