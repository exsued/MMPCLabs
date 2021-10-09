// LinearEquationSolves
package main

//"fmt"
func GaussianConvertToTriangleView(in [][]float64) [][]float64 {
	//Прямой ход метода Гаусса
	//j - ведущий элемент
	//Копируем, чтобы не менять исходную матрицу
	input := make([][]float64, len(in))
	for i := 0; i < len(in); i++ {
		input[i] = make([]float64, len(in[i]))
		for j := 0; j < len(in[i]); j++ {
			input[i][j] = in[i][j]
		}
	}
	for j := 1; j < len(input); j++ {
		copiedRow := CopyVector(input[j-1])
		fstElement := input[j-1][j-1]
		dividedRow := divideVectorRes(copiedRow, fstElement)
		for i := j; i < len(input); i++ {
			fstEl := input[i][j-1]
			substractVector(input[i], multiplyVectorRes(dividedRow, fstEl))
		}
	}
	return input
}
func GaussianSingularDivision(in [][]float64) []float64 {

	//PrintMatrix(input)
	//Прямой ход метода Гаусса
	//j - ведущий элемент
	//Копируем, чтобы не менять исходную матрицу
	input := make([][]float64, len(in))
	for i := 0; i < len(in); i++ {
		input[i] = make([]float64, len(in[i]))
		for j := 0; j < len(in[i]); j++ {
			input[i][j] = in[i][j]
		}
	}
	for j := 0; j < len(input); j++ {

		//Деление на ведущий элемент
		divideVector(input[j], input[j][j])
		//Исключение j-ой переменной из остальных уравнений
		for i := j + 1; i < len(input); i++ {
			firstRow := CopyVector(input[j])
			multiplyVector(firstRow, input[i][j])
			substractVector(input[i], firstRow)
		}
	}
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
	//fmt.Println()
	//PrintVector(x)
	//fmt.Println()
	return x
}
func GaussianDeterminant(in [][]float64) (result float64) {
	input := GaussianConvertToTriangleView(in)
	result = 1
	for i := 0; i < len(input); i++ {
		result *= input[i][i]
	}
	return result
}
