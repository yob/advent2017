package two

import (
	"testing"
)

func TestChecksum(t *testing.T) {
	if Checksum("") != 0 {
		t.Error("checksum for empty string should equal 0")
	}
	result := Checksum("5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8")
	if result != 18 {
		t.Error("checksum should equal 18, not: ", result)
	}
}

func TestChecksumDivisible(t *testing.T) {
	if ChecksumDivisible("") != 0 {
		t.Error("checksum for empty string should equal 0")
	}
	result := ChecksumDivisible("5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5")
	if result != 9 {
		t.Error("checksum should equal 9, not: ", result)
	}
}
