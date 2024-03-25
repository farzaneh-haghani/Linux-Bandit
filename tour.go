package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

// We need var because it is outside a function
var	(
	c,python bool=true,false
	MaxInt uint64     = 1<<64 - 1    // keep 64 bits between 0 to 2^64-1 which is maximum possible value for a unit 64 type
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)


func main() {
	g := 0.867 + 0.5i // complex128
	d:=g   //  d's type is inferred from g
	fmt.Printf("%T \n",d)
	// Short assignment statement only inside a func 
	typeOmitted:=9
	// Convert the type
	var f float64 = math.Sqrt(float64(typeOmitted))
	var i float64   // Without initializing are given 0 / false / ""
	fmt.Println(f,i)

	fmt.Println(time.Now())
	fmt.Println("Random number", rand.Intn(10))
	fmt.Printf("Square root %f \n", math.Sqrt(7))
	fmt.Println(math.Pi)

	a, b := funcTest("Hello",2,2)
	c:=naked(5)
	fmt.Println(a, b,c)

	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)

	// 7/3 => 2
	// 7.0/3.0 => 2.3333333333333335
	fmt.Println(7.0/3.0)  

	var x int
	if _,err:=fmt.Scanf("%d",&x); err!=nil {
		fmt.Println("Error",err)
		return
	}
		// _,err :=fmt.Scanf("%d",&input)
	// if err!=nil{
	// 	fmt.Println("Error",err)
	// 	return
	// }

	if x>10{
		j:=0
		for j<1{
			fmt.Println("passed")
			j++
		}
	}else{
		for i:=0;i<x;i++{
		fmt.Println("failed")
		}
	}

	// var input any="Hi"
	var input interface{}=15
	switch t:=input.(type){    //type inferred
	case string:
		fmt.Println(input.(string))
		fmt.Println(t)
	case int:
		fmt.Println(input.(int) + 10)    // casting type
		fmt.Println(t + 10) 
	default:
		fmt.Println("unknown",t)
	}

	name:="Farzaneh"
	switch name{
	case "Farzaneh":
		fmt.Println(name,"Haghani")
		fallthrough                    // Run next case as well
	case "CYF":
		fmt.Println("Trainee")
	}

	arr:=[5]int{   // Array can not be resized
		1,2,3,    //Array fills 0 for the rest
	}
	fmt.Println(arr)

	// slice:=[]int{   
	// 	1,2,3,
	// }
	slice:=make([]string,0,2)
	slice = append(slice, "a","b","c")
	fmt.Println(slice)
	fmt.Println(cap(slice))      // Capacity will be double if we append more
	fmt.Println(len(slice))

	slice=append(slice[1:],slice...)
	slice=append(slice[:1],slice[1:]...)
	fmt.Println(slice)

}


func funcTest(x string , y,z int) (string,int) {
	return x,y+z
}

func naked(x int)(x2 int){
	x2=x/4
return
}

//printf for using %s,d... for replacing
// println for printing text with new line


// Types:
// bool, string
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte      // alias for uint8
// rune      // alias for int32 that represents a Unicode code point
// float32 float64
// complex64 complex128



// %d  int base 10
// %b  int base 2
// %f  float    6.200000 - 2.645751 
// %s  string
// %g  numbers in correct size  6.2  - 2.6457513110645907
// %T  type
// %v  value