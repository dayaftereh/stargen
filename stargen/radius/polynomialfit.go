package radius

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func vandermonde(a []float64, degree int) *mat.Dense {
	x := mat.NewDense(len(a), degree+1, nil)
	for i := range a {
		for j, p := 0, 1.; j <= degree; i, p = j+1, p*a[i] {
			x.Set(i, j, p)
		}
	}
	return x
}

func polynomialfit(degree int, x []float64, y []float64) []float64 {
	a := vandermonde(x, degree)
	b := mat.NewDense(len(y), 1, y)
	c := mat.NewDense(degree+1, 1, nil)

	qr := new(mat.QR)
	qr.Factorize(a)

	err := qr.SolveTo(c, false, b)
	if err != nil {
		fmt.Println(err)
	}

	result := make([]float64, degree)
	for i := 0; i < degree; i++ {
		result[i] = c.At(i, 0)
	}
	return result
}
