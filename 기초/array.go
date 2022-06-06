package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	var a1 = [3]int{1, 2, 3}
	var a2 = [...]int{1, 2, 3}
	fmt.Println(a1, a2)

	var multiArray [3][4][5]int
	multiArray[0][1][2] = 10
	fmt.Println(multiArray)

	var a3 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(a3)
}
