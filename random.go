/************************************************************
 * Miscellaneous Go functions                               *
 * Time-stamp: <2022-11-06 11:43:16 christophe@pallier.org> *
 ************************************************************/

package goutils

import (
	"pgregory.net/rand"
)

// RandomU returns n random numbers between 0 and 1.
func RandomU(n int, seed uint64) []float32 {
	x := make([]float32, n)
	r := rand.New(seed)
	for i := 0; i < n; i++ {
		x[i] = r.Float32()
	}
	return x
}

// RandomIntn returns n random integer between 0 and max - 1
func RandomIntn(n int, max int, seed uint64) []int {
	x := make([]int, n)
	r := rand.New(seed)
	for i := 0; i < n; i++ {
		x[i] = r.Intn(max)
	}
	return x
}
