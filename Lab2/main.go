package main

import ("fmt" 
	"math")

func ConvergenceCondition (a [][]float64) (res bool){
	for i:= 0; i < len(a); i++ {
		sum := sumElementsVectorAbs(a[i]) - math.Abs(a[i][i]) -
		math.Abs(a[i][len(a[i]) - 1])
		if math.Abs(a[i][i]) <= sum {
			fmt.Println("Нет сходимости")
			return false
		}
	}
	return true
}

func MaxElement(vec []float64) (res float64){
	for i:=0; i < len(vec); i++ {
		if res < vec[i] {
			res = vec[i]
		}
	}
	return
}
func CalculateErr(x []float64, prx []float64) (res float64) {
	errVec := make([]float64, len(x))
	for i:=0; i < len(x); i++ {
		errVec[i] = math.Abs(x[i]-prx[i])
	}
	res = MaxElement(errVec)
	return
}
func SimpleIter(a [][]float64, accuracy float64) {
	prx := make([]float64, len(a))
	iters := 0
	var e float64
	for {
		x := make([]float64, len(a))
		for i:=0; i < len(a); i++ {
			for j:=0; j < len(a); j++ {
				if j != i {
					//prevx := x[i]
					x[i] += -a[i][j]*prx[j]
					//fmt.Println(prevx," + ", -a[i][j], " * ", prx[j], " = ", x[i])
				}
			}
			x[i] += a[i][len(a[i]) - 1]
			x[i] /= a[i][i]
		}
		e = CalculateErr(x, prx)
		prx = CopyVector(x)
		iters++
		if e <= accuracy {
			break
		}
	}
	fmt.Println(prx, " ", e, " ", iters)
}
func Seidel(a [][]float64, accuracy float64) {
	prx := make([]float64, len(a))
	iters := 0
	var e float64
	for {
		x := make([]float64, len(a))
		for i:=0; i < len(a); i++ {
			for j:=0; j < len(a); j++ {
				if j < i {
					x[i] += -a[i][j]*x[j]
				}else if j > i {
					x[i] += -a[i][j]*prx[j]
				}
			}
			x[i] += a[i][len(a[i]) - 1]
			x[i] /= a[i][i]
		}
		e = CalculateErr(x, prx)
		prx = CopyVector(x)
		iters++
		if e <= accuracy {
			break
		}
	}
	fmt.Println(prx, " ", e, " ", iters)
}
func main(){
	//Вариант 10
	a := [][]float64{
		{5,   -1, 3, 5},
		{1, -4, 2, 20},
		{2, -1, 5, 10},
	}
	/*
	a := [][]float64{
		{12.1,   0.0528, 0.0639, 0.0749, 14.8310},
		{0.0365, 11.2,   0.0586, 0.0697, 15.9430},
		{0.0312, 0.0423, 10.3,   0.0644, 15.6926},
		{0.026,  0.037,  0.0481, 9.4,    17.08},
	}
	*/
	if(ConvergenceCondition(a)) {
		SimpleIter(a, 0.001)
		Seidel(a, 0.001)
	}
	//PrintMatrix(a)
}

