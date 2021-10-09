// main.go
package main

import (
	"fmt"
)

func GetDiscrepancyVec(m [][]float64, x []float64) []float64 {
	//Вычисление AX
	AX := make([]float64, len(x))
	B := make([]float64, len(m))
	for i := 0; i < len(m); i++ {
		B[i] = m[i][len(m[i])-1]
	}
	for i := 0; i < len(m); i++ {
		leftSide := m[i][:len(m[i])-1]
		for j := 0; j < len(leftSide); j++ {
			AX[i] += leftSide[j] * x[j]
		}
	}
	//PrintVector(AX)
	return substractVectorRes(B, AX)
}
func main() {
	fmt.Println("Gaussian singular division:")

	inputMatrix := [][]float64{
		{4, 2, -3, -3},
		{5, 3, -5, -8},
		{4, 1, 5, 22},
	}

	/*
		inputMatrix := [][]float64{
			{18, 12, -43, -32},
			{15, 23, -75, -98},
			{74, 61, 55, 22},
		}
	*/
	PrintMatrix(inputMatrix)
	solve := GaussianSingularDivision(inputMatrix)
	fmt.Println("Solve vector:")
	PrintVector(solve)
	fmt.Println()
	disperancy := GetDiscrepancyVec(inputMatrix, solve)
	fmt.Println("Disperancy vector:")
	PrintVector(disperancy)
	fmt.Println()

	fmt.Println("Determinant calculation:")
	inputMatrix = [][]float64{
		{4, 2, -3},
		{5, 3, -5},
		{4, 1, 5},
	}
	PrintMatrix(inputMatrix)
	fmt.Println("Determinant:")
	fmt.Println(GaussianDeterminant(inputMatrix))
}
