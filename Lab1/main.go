// main.go
package main

import (
	"fmt"
)

func GaussianSingularDivision(input [][]float64) {
	//Прямой ход метода Гаусса
	//j - ведущий элемент
	for j := 0; j < len(input); j++ {

		//Деление на ведущий элемент
		divideVector(input[j], input[j][j])
		//Исключение j-ой переменной из остальных уравнений
		for i := j + 1; i < len(input); i++ {
			firstRow := CopyVector(input[j])
			multipleVector(firstRow, input[i][j])
			substractVector(input[i], firstRow)
		}
	}
	PrintMatrix(input)

	//Обратный ход
	x := make([]float64, len(input))

	rowsSize := len(input)
	lstRowSize := len(input[rowsSize-1])
	x[rowsSize-1] = input[rowsSize-1][lstRowSize-1]

	for i := rowsSize - 1; i >= 0; i-- {
		//Суммирование всех известных элементов в левой части и вынесение за скобку
		var sum float64
		for j := 0; j < len(input[i])-1; j++ {
			if input[i][j] != 0 && input[i][j] != 1 {
				sum += input[i][j] * x[j]
			}
		}
		x[i] = input[i][len(input[i])-1] - sum
	}
	fmt.Println()
	PrintVector(x)
	fmt.Println()
}

func EnterMatrix() (result [][]float64) {
	var rowsCount, columnsCount int
	fmt.Println("Enter rows count:")
	fmt.Scan(&rowsCount)
	fmt.Println("Enter columns count:")
	fmt.Scan(&columnsCount)

	result = make([][]float64, rowsCount)
	for i := range result {
		result[i] = make([]float64, columnsCount)
	}
	for i := 0; i < rowsCount; i++ {
		for j := 0; j < columnsCount; j++ {
			fmt.Scan(&result[i][j])
		}
	}
	return result
}
func main() {
	fmt.Println("Gaussian singular division")

	/*
		inputMatrix := [][]float64{
			{4, 2, -3, -3},
			{5, 3, -5, -8},
			{4, 1, 5, 22},
		}
	*/
	inputMatrix := EnterMatrix()
	GaussianSingularDivision(inputMatrix)
}
