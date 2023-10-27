package main

// Definition of auxiliary methods

func dotProd(u [3]float64, v [3]float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(u); i++ {
		sum += u[i] * v[i]
	}

	return sum
}

func vectorProd(u [3]float64, v [3]float64) [3]float64 {
	var w0 = u[1]*v[2] - u[2]*v[1]
	var w1 = u[2]*v[0] - u[0]*v[2]
	var w2 = u[0]*v[1] - u[1]*v[0]

	return [3]float64{w0, w1, w2}
}

func vec(x float64, y float64, z float64) [3]float64 {
	return [3]float64{x, y, z}
}
