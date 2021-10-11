package main

func GaussianInverseMatrix(in [][]float64) (res [][]float64) {
	//Создание расширенной матрицы - копии входной + единичная в правой части
	input := make([][]float64, len(in))
	for i := 0; i < len(in); i++ {
		input[i] = make([]float64, len(in[i])*2)
		for j := 0; j < len(in[i]); j++ {
			input[i][j] = in[i][j]
		}
	}
	for i := 0; i < len(input); i++ {
		input[i][len(input[i])/2+i] = 1
	}
	//PrintMatrix(input)
	M1 := GaussianConvertToUpTriangleView(input)
	M2 := GaussianConvertToDownTriangleView(M1)
	
	final := make([][]float64, len(M2))
	for i := 0; i < len(M2); i++ {
		final[i] = make([]float64, len(M2[i]))
		for j := 0; j < len(M2[i]); j++ {
			final[i][j] = M2[i][j]
		}
	}
	
	for i := 0; i < len(final); i++ {
		divideVector(final[i], final[i][i])
	}

	res = make([][]float64, len(final))

	for i := 0; i < len(final); i++ {
		rLen := len(final[i])/2

		res[i] = make([]float64, rLen)
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = final[i][j + rLen]
		}
	}
	return
}
func GaussianConvertToDownTriangleView(in [][]float64) [][]float64 {
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
func GaussianConvertToUpTriangleView(in [][]float64) [][]float64 {
	//Копируем, чтобы не менять исходную матрицу
	input := make([][]float64, len(in))
	for i := 0; i < len(in); i++ {
		input[i] = make([]float64, len(in[i]))
		for j := 0; j < len(in[i]); j++ {
			input[i][j] = in[i][j]
		}
	}
	for j:= len(input) - 1; j > 0; j-- {
		lstI := j
		copiedRow := CopyVector(input[lstI])
		fstElement := input[lstI][lstI]
		dividedRow := divideVectorRes(copiedRow, fstElement)
		for i:= lstI - 1; i >= 0; i-- {
			fstEl := input[i][lstI]
			substractVector(input[i], multiplyVectorRes(dividedRow, fstEl))
		}
	}
	return input
}
func GaussianDeterminant(in [][]float64) (result float64) {
	input := GaussianConvertToDownTriangleView(in)
	result = 1
	for i := 0; i < len(input); i++ {
		result *= input[i][i]
	}
	return result
}