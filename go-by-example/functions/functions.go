package main

import "fmt"

func plus(a int, b int)int{
return a + b
}

func plusplus(a,b,c int)int{

	return a + b + c
}

//multiple return value function
func vals() (int, string){

	return 3, "some string"

}

func main(){


	addVal := plus(1, 2)
	fmt.Println("add: ",addVal)

	addVal = plusplus(1, 2, 3)
	fmt.Println("1 + 2 + 3:", addVal)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

}