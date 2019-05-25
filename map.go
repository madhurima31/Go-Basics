package main 

import "fmt"


func main(){

	fruits := make(map[string]int8)

	fruits["mangoes"] = 2
	fruits["apples"] = 1
	fruits["melon"] = 1

	fmt.Println(fruits)

	delete(fruits, "melon")
	fmt.Println(fruits)

	for k,v := range fruits{
		fmt.Println(k ,":",v)
	}

}