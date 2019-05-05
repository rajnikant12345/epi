package ch2

import (
	"testing"
)

func compareintarray(in1, in2 []int) bool {
	for i, v := range in1 {
		if v != in2[i] {
			return false
		}
	}
	return true

}

func TestDutchNationalFlag(t *testing.T) {
	out := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	in := [][]int{[]int{1, 2, 1, 2, 1, 2, 3, 3, 3}, []int{2, 2, 2, 3, 3, 3, 1, 1, 1},
		[]int{3, 3, 3, 2, 2, 2, 1, 1, 1}, []int{1, 2, 3, 1, 2, 3, 1, 2, 3},
		[]int{3, 2, 1, 3, 2, 1, 3, 2, 1}}

	for _, v := range in {
		out1 := DutchNationalFlag(v, 2)
		if !compareintarray(out, out1) {
			t.FailNow()
		}
	}

}

func TestIncreamentLongPrecisionNumber(t *testing.T) {
	in := []string{"12345", "12349", "99999", "899999"}
	out := []string{"12346", "12350", "100000", "900000"}

	for i := 0; i < len(in); i++ {
		tmp := IncreamentLongPrecisionNumber(in[i])
		if tmp != out[i] {
			t.FailNow()
		}
	}
}

func TestIMultiplyLongs(t *testing.T) {
	if "15129" != MultiplyLongs("123", "123") {
		t.FailNow()
	}
	if "998001" != MultiplyLongs("999", "999") {
		t.FailNow()
	}
	if "" != MultiplyLongs("0", "999") {
		t.FailNow()
	}
	if "999" != MultiplyLongs("1", "999") {
		t.FailNow()
	}
	if "121" != MultiplyLongs("11", "11") {
		t.FailNow()
	}
}
