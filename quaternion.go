package main

import (
	"fmt"
	"math"
)

// Quaternion type definition
type Quaternion struct {
	scalarPart float64
	vectorPart [3]float64
	sub        func(other Quaternion) Quaternion
	mul        func(other Quaternion) Quaternion
	div        func(other Quaternion) Quaternion
}

func (q Quaternion) String() string {
	return fmt.Sprintf("Quaternion(%g + [%g, %g, %g])", q.scalarPart, q.vectorPart[0], q.vectorPart[1], q.vectorPart[2])
}

func (q Quaternion) abs() float64 {
	return math.Sqrt(q.scalarPart*q.scalarPart + dotProd(q.vectorPart, q.vectorPart))
}

func (q Quaternion) arg() float64 {
	return math.Acos(q.scalarPart / math.Sqrt(dotProd(q.vectorPart, q.vectorPart)))
}

func (q Quaternion) cnj() Quaternion {
	var newVectorPart = [3]float64{-q.vectorPart[0], -q.vectorPart[1], -q.vectorPart[2]}
	return Quaternion{scalarPart: q.scalarPart, vectorPart: newVectorPart}
}

func (q Quaternion) add(other Quaternion) Quaternion {
	var newScalarPart = q.scalarPart + other.scalarPart
	var newVectorPart = [3]float64{
		q.vectorPart[0] + other.vectorPart[0],
		q.vectorPart[1] + other.vectorPart[1],
		q.vectorPart[2] + other.vectorPart[2],
	}

	return Quaternion{scalarPart: newScalarPart, vectorPart: newVectorPart}
}
