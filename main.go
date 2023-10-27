package main

import "fmt"

func main() {
	var q1 = Quaternion{scalarPart: 1, vectorPart: vec(1, 2, 3)}
	var q2 = Quaternion{scalarPart: 1, vectorPart: vec(-1, -2, -3)}

	fmt.Println(q1, q2, q1.add(q2))
}
