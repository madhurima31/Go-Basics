package main
import "fmt"

func main(){

	x:=2
	y:=&x
	fmt.Println(x,y)

	z:= *y

	//prints 2
	fmt.Println(z)

	*y=10
	//prints 10
	fmt.Println(x)
}