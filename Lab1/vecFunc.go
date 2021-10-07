// vecFunc
package main

import (
	"fmt"
)

func PrintVector(a []float64) {
	for j := 0; j < len(a); j++ {
		fmt.Printf("%f ", a[j])
	}
}
func PrintMatrix(a [][]float64) {
	for i := 0; i < len(a); i++ {
		PrintVector(a[i])
		fmt.Println()
	}
}
func CopyVector(src []float64) []float64 {
	result := make([]float64, len(src))
	copy(result, src)
	return result
}
func multipleVector(vec []float64, val float64) {
	for i := 0; i < len(vec); i++ {
		vec[i] *= val
	}
}
func divideVector(vec []float64, val float64) {
	for i := 0; i < len(vec); i++ {
		vec[i] /= val
	}
}
func substractVector(vec1 []float64, vec2 []float64) {
	for i := 0; i < len(vec1); i++ {
		vec1[i] -= vec2[i]
	}
}
