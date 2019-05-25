package main

import "fmt"


//user defined type
type fruits struct{
	mangoes int8
	apples  int8
	bananas int8
}

func main() {

fruit := fruits{2,3,4}
fmt.Println(fruit.mangoes)

}