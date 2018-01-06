package one

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum("") != 0 {
		t.Error("empty string should equal 0")
	}
	if Sum("1122") != 3 {
		t.Error("1122 != 3")
	}
	if Sum("1111") != 4 {
		t.Error("1111 != 4")
	}
	if Sum("1234") != 0 {
		t.Error("1234 != 0")
	}
	if Sum("91212129") != 9 {
		t.Error("91212129 != 9")
	}
}
