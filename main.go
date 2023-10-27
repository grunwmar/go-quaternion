package main

import "fmt"

func main() {
	var p = Quaternion(1, 1, -1, -1)
	var q = Quaternion(0, -1, 1, 1)
	fmt.Println(p.add(q).vectorPart)
}
