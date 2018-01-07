package two

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Checksum("") != 0 {
		t.Error("checksum for empty string should equal 0")
	}
	result := Checksum("5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8")
	if result != 18 {
		t.Error("checksum should equal 18, not: ", result)
	}
}
