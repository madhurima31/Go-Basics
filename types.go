package main

import ("fmt")

// func name(params type) (return type) {}

func testFunc(x1,x2 string) (string,string){
	return x1,x2
}

func add(a1,b1 float64) float64{
	return a1+b1
}

func main(){

	var m float64 = 2.1
	var n float64 = 3.1

	fmt.Println(add(m,n))

	x1,x2 := "Hey","Beautiful !"
	fmt.Println(testFunc(x1,x2))
}