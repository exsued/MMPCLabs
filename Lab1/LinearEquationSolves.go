// LinearEquationSolves
package main

//"fmt"

func GaussianSingularDivision(in [][]float64) []float64 {
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
			multipleVector(firstRow, input[i][j])
			substractVector(input[i], firstRow)
		}
	}
	//PrintMatrix(input)

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
