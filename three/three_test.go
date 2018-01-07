package three

import (
	"testing"
)

func TestDistance(t *testing.T) {
	result := Distance(1)
	if result != 0 {
		t.Error("distance for 1 should equal 0, not ", result)
	}
	result = Distance(12)
	if result != 3 {
		t.Error("distance for 12 should equal 3, not ", result)
	}
	result = Distance(23)
	if result != 2 {
		t.Error("distance for 23 should equal 2, not ", result)
	}
	result = Distance(1024)
	if result != 31 {
		t.Error("distance for 1024 should equal 31, not ", result)
	}
}
