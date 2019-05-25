package main

import "fmt"



type fruits struct{
	mangoes int8
	oranges int8
	watermelons int8
}

// value receivers  - copy of the type is passed 
//                  - doesnot modify the actual attributes inside the type struct
func (f fruits) totalFruits() int8{
	return f.mangoes + f.oranges +f.watermelons
}

// pointer receivers   - passing pointer of type
//                     - modifies the actual attributes inside the type struct

func (f *fruits) new_number_of_mangoes(m int8) {
	f.mangoes = m 
}

func main(){

	fruit := fruits{2,3,3}
	fmt.Println(fruit.totalFruits())

	fruit.new_number_of_mangoes(5)
	fmt.Println(fruit.totalFruits())

}