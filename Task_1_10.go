package main

import (
	"fmt"
	"math"
)
func main() {
	fmt.Println("Enter a:")
	var a float64
	fmt.Scanln(&a)

	z1 := (math.Sin(math.Pi/2 + 3*a))/(1-math.Sin(3*a-math.Pi))
	z2 := math.Cos(5/4*math.Pi+3/2*a)/math.Sin(5/4*math.Pi+3/2*a)
	fmt.Println(z1, z2)
}
