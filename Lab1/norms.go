package main

import(
	   "math")

func vecNormSum(vec []float64) (res float64) {
	for i:= 0; i < len(vec); i++ {
		res += math.Abs(vec[i])
	}
	return
}
func vecNormChebishev(vec []float64) (res float64) {
	for i:= 0; i < len(vec); i++ {
		if res < vec[i] {
			res = math.Abs(vec[i])
		}
	}
	return
}
func vecNormEuclidian(vec []float64) (res float64) {
	for i:= 0; i < len(vec); i++ {
		res += math.Pow(vec[i], 2)
	}
	res = math.Sqrt(res)
	return
}
func matrNormM (m [][]float64) (res float64) {
	for i:=0; i < len(m); i++ {
		var sum float64
		for j:=0; j < len(m[i]); j++ {
			sum += math.Abs(m[i][j])
		}
		res = math.Max(res, sum)
	}
	return
}
func matrNormL (m [][]float64) (res float64) {
	for i:=0; i < len(m[0]); i++ {
		var sum float64
		for j:=0; j < len(m); j++ {
			sum += math.Abs(m[j][i])
		}
		res = math.Max(res, sum)
	}
	return res
}
func matrNormK (m [][]float64) (res float64) {
	for i:=0; i < len(m); i++ {
		for j:=0; j < len(m[i]); j++ {
			res += math.Pow(m[i][j], 2)
		}
	}
	res = math.Sqrt(res)
	return
}

func CondA_M (m [][]float64) (res float64) {
	invM := GaussianInverseMatrix(m)
	mNorm := matrNormM(m)
	imNorm := matrNormM(invM)
	res = mNorm * imNorm
	return
}
func CondA_L (m [][]float64) (res float64) {
	invM := GaussianInverseMatrix(m)
	mNorm := matrNormL(m)
	imNorm := matrNormL(invM)
	res = mNorm * imNorm
	return
}
func CondA_K (m [][]float64) (res float64) {
	invM := GaussianInverseMatrix(m)
	mNorm := matrNormK(m)
	imNorm := matrNormK(invM)
	res = mNorm * imNorm
	return
}