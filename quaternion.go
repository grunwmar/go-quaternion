package main

import "math"

// QuaternionType type definition
type QuaternionType struct {
	scalarPart float64
	vectorPart [3]float64
	abs        func() float64
	arg        func() float64
	cnj        func() QuaternionType
	add        func(other QuaternionType) QuaternionType
	sub        func(other QuaternionType) QuaternionType
	mul        func(other QuaternionType) QuaternionType
	div        func(other QuaternionType) QuaternionType
}

// Quaternion instance constructor
func Quaternion(s float64, i float64, j float64, k float64) QuaternionType {
	var quaternion = QuaternionType{}
	quaternion.scalarPart = s
	quaternion.vectorPart = [3]float64{i, j, k}
	quaternion.abs = func() float64 { return abs(quaternion) }
	quaternion.arg = func() float64 { return arg(quaternion) }
	quaternion.cnj = func() QuaternionType { return cnj(quaternion) }
	quaternion.add = func(other QuaternionType) QuaternionType { return add(quaternion, other) }
	return quaternion
}

// QuaternionType object Definition of methods
func abs(q QuaternionType) float64 {
	return math.Sqrt(q.scalarPart*q.scalarPart + dotProd(q.vectorPart, q.vectorPart))
}

func arg(q QuaternionType) float64 {
	return math.Acos(q.scalarPart / math.Sqrt(dotProd(q.vectorPart, q.vectorPart)))
}

func cnj(q QuaternionType) QuaternionType {
	return Quaternion(q.scalarPart, -q.vectorPart[0], -q.vectorPart[1], -q.vectorPart[2])
}

func add(q1 QuaternionType, q2 QuaternionType) QuaternionType {
	var newScalarPart = q1.scalarPart + q2.scalarPart
	var newVectorPart = [3]float64{
		q1.vectorPart[0] + q2.vectorPart[0],
		q1.vectorPart[1] + q2.vectorPart[1],
		q1.vectorPart[2] + q2.vectorPart[2],
	}

	return Quaternion(newScalarPart, newVectorPart[0], newVectorPart[1], newVectorPart[2])
}
