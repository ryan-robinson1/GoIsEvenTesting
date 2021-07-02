/*
   Ryan Robinson, 2021

   Simple testing and benchmarking for isEven function
*/

package main

import (
	"math"
	"strconv"
	"testing"
)

//Tests parity of function from -100 to 100. Zero is assumed to be even
func TestIsEven_Range100(t *testing.T) {
	evenFlag := true
	for i := -100; i <= 100; i++ {
		if evenFlag {
			even := IsEven(i)
			if even != true {
				t.Errorf("Returned parity was incorrect for %v, got: %t, want: %t.", i, even, true)
			}
			evenFlag = false
		} else {
			even := IsEven(i)
			if even != false {
				t.Errorf("Returned parity was incorrect for %v, got: %t, want: %t.", i, even, false)
			}
			evenFlag = true
		}
	}
}

//Tests parity of function on arbituary large 8 digit positive numbers
func TestIsEven_LargeNum(t *testing.T) {
	evenFlag := false
	for i := 24303443; i <= 24303463; i++ {
		if evenFlag {
			even := IsEven(i)
			if even != true {
				t.Errorf("Returned parity was incorrect for %v, got: %t, want: %t.", i, even, true)
			}
			evenFlag = false
		} else {
			even := IsEven(i)
			if even != false {
				t.Errorf("Returned parity was incorrect for %v, got: %t, want: %t.", i, even, false)
			}
			evenFlag = true
		}
	}
}

//Tests parity of function on arbituary large 8 digit negative numbers
func TestIsEven_LargeNegNum(t *testing.T) {
	evenFlag := false
	for i := -21372345; i <= -21372326; i++ {
		if evenFlag {
			even := IsEven(i)
			if even != true {
				t.Errorf("Returned parity was incorrect for %v, got: %t, want: %t.", i, even, true)
			}
			evenFlag = false
		} else {
			even := IsEven(i)
			if even != false {
				t.Errorf("Returned parity was incorrect for %v, got: %t, want: %t.", i, even, false)
			}
			evenFlag = true
		}
	}
}

//Tests parity of function on the upper bound of int32, which is known to be odd
func TestIsEven_MaxUpper32(t *testing.T) {
	in := math.MaxInt32
	exOut := false
	even := IsEven(in)

	if even != exOut {
		t.Errorf("Returned parity was incorrect, got: %t, want: %t.", even, exOut)
	}
}

//Tests parity of function on the lower bound of int32, which is known to be even
func TestIsEven_MaxLower32(t *testing.T) {
	in := math.MinInt32
	exOut := true
	even := IsEven(in)

	if even != exOut {
		t.Errorf("Returned parity was incorrect, got: %t, want: %t.", even, exOut)
	}
}

//Tests parity of function on the upper bound of int64, which is known to be odd, if the machine's int type is 64 bits
func TestIsEven_MaxUpper64(t *testing.T) {
	if strconv.IntSize == 64 {
		in := math.MaxInt64
		exOut := false
		even := IsEven(in)

		if even != exOut {
			t.Errorf("Returned parity was incorrect, got: %t, want: %t.", even, exOut)
		}
	}
}

//Tests parity of function on the lower bound of int64, which is known to be even, if the machine's int type is 64 bits
func TestIsEven_MaxLower64(t *testing.T) {
	if strconv.IntSize == 64 {
		in := math.MinInt64
		exOut := true
		even := IsEven(in)

		if even != exOut {
			t.Errorf("Returned parity was incorrect, got: %t, want: %t.", even, exOut)
		}
	}
}

func BenchmarkIsEven(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsEven(i)
	}
}
