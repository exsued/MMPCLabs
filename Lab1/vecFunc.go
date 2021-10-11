// vecFunc
package main

import (
	"fmt"
)

func EnterMatrix() (result [][]float64) {
	var rowsCount, columnsCount int
	fmt.Println("Enter rows:")
	fmt.Scanln(&rowsCount)
	fmt.Println("Enter colums:")
	fmt.Scanln(&columnsCount)
	result = make([][]float64, rowsCount)
	for i := range result {
		result[i] = make([]float64, columnsCount)
	}
	for i := 0; i < rowsCount; i++ {
		for j := 0; j < columnsCount; j++ {
			fmt.Printf("Enter the element for Matrix1 %d %d :", i+1, j+1)
			fmt.Scanln(&result[i][j])
		}
	}
	return result
}
func PrintVector(a []float64) {
	for j := 0; j < len(a); j++ {
		fmt.Printf("%f ", a[j])
	}
	fmt.Println()
}
func PrintMatrix(a [][]float64) {
	for i := 0; i < len(a); i++ {
		PrintVector(a[i])
	}
	fmt.Println()
}
func CopyVector(src []float64) []float64 {
	result := make([]float64, len(src))
	copy(result, src)
	return result
}
func multiplyVector(vec []float64, val float64) {

	for i := 0; i < len(vec); i++ {
		vec[i] *= val
	}
}
func multiplyVectorRes(vec []float64, val float64) (res []float64) {
	res = make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		res[i] = vec[i] * val
	}
	return
}
func divideVector(vec []float64, val float64) {
	for i := 0; i < len(vec); i++ {
		vec[i] /= val
	}
}
func divideVectorRes(vec []float64, val float64) (res []float64) {
	res = make([]float64, len(vec))
	for i := 0; i < len(vec); i++ {
		res[i] = vec[i] / val
	}
	return
}
func sumVector(vec1 []float64, vec2 []float64) {
	for i := 0; i < len(vec1); i++ {
		vec1[i] += vec2[i]
	}
}
func substractVector(vec1 []float64, vec2 []float64) {
	for i := 0; i < len(vec1); i++ {
		vec1[i] -= vec2[i]
	}
}
func substractVectorRes(vec1 []float64, vec2 []float64) (res []float64) {
	res = make([]float64, len(vec1))
	for i := 0; i < len(vec1); i++ {
		res[i] = vec1[i] - vec2[i]
	}
	return res
}