package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println("Random number", rand.Intn(10))
	fmt.Printf("Square root %g", math.Sqrt(7))
	fmt.Println(math.Pi)
	a, b := funcTest("Hello",2,2)
	fmt.Println(a, b)
}


func funcTest(x string , y,z int) (string,int) {
	return x,y+z
}


