package main

import "fmt"

func intSeq() func() int{
	i := 0
	return func() int {
		i++

		return i
	}
}

func letter() func() string{
	someLetter :=  "A"
	fmt.Println(someLetter)
	return func() string{
		someLetter ="B"	
	return someLetter
	}
}

func main() {
nextInt := intSeq()

fmt.Println(nextInt())
fmt.Println(nextInt())
fmt.Println(nextInt())

newInts := intSeq()

var somefunc = letter()

fmt.Println(newInts())

fmt.Println(somefunc())

}