package main

import (
	"fmt"
	"slices"
	"strconv"
)


func IntegerToBinary (n int) string{
	return strconv.FormatInt(int64(n), 2)
}

func main(){

	var s[] string
	fmt.Println("emp:", s, s ==nil, len(s) == 0)
	
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"


	fmt.Println("set:", s)
	fmt.Println("get:", s[2])


	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
	fmt.Println("len:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("scl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2){
		fmt.Println("t == t2")
	}

	twoD := make([][]string, 30)
	for i := 0; i < len(twoD); i++{ //len(twoD) = 3
		innerLen := (i + 1)
		twoD[i] = make([]string, innerLen)
		for j := 0; j < innerLen; j++{
			twoD[i][j] = IntegerToBinary(i) + IntegerToBinary(j)
		}
	}

	fmt.Println("2d: ", twoD)

}