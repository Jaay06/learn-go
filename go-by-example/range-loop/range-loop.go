package main

import "fmt"

//test recursion
func test(n int) int{

	if n ==0 {
		return 1
	}
	return n * test(n - 1) //call function again (recursion)

}

func main() {

	nums := []int{2, 3, 4}

	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana", "c": "car"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	testSum := test(20)

	fmt.Println(testSum)

}