// package main

// import (
// 	"fmt"
// 	"math"
// )

// func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
// 	return func(time float64) float64 {
// 		return 0.5*a*math.Pow(time, 2) + vo*time + so
// 	}
// }

// func main() {
// 	var acceleration, initialVelocity, initialDisplacement, time float64

// 	fmt.Println("Enter acceleration (float value):")
// 	fmt.Scanln(&acceleration)

// 	fmt.Println("Enter initial velocity (float value):")
// 	fmt.Scanln(&initialVelocity)

// 	fmt.Println("Enter initial displacement (float value):")
// 	fmt.Scanln(&initialDisplacement)

// 	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

// 	fmt.Println("Enter time (in seconds):")
// 	fmt.Scanln(&time)

// 	displacement := fn(time)
// 	fmt.Println("Displacement after", time, "seconds:", displacement)
// }
