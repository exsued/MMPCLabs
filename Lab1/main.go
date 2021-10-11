// main.go
package main

import ("fmt")

func main() {
	fmt.Println("Gaussian singular division:")

	inputMatrix := [][]float64{
		{4, 2, -3, -3},
		{5, 3, -5, -8},
		{4, 1, 5, 22},
	}
	
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
		{1, -1, 1},
		{2, 2, 3},
		{-2, 4, -1},
	}
	PrintMatrix(inputMatrix)
	fmt.Println("Determinant:")
	fmt.Println(GaussianDeterminant(inputMatrix))
	
	fmt.Println("Inverse matrix:")
	invM := GaussianInverseMatrix(inputMatrix)
	PrintMatrix(invM)

	fmt.Println()
	fmt.Println("Cond m A: ", CondA_M(inputMatrix))
	fmt.Println("Cond l A: ", CondA_L(inputMatrix))
	fmt.Println("Cond k A: ", CondA_K(inputMatrix))
}
