package main

import "fmt"

func main() {
	// Raw String Literal. 복수라인.
	rawLiteral := `아리랑\n
아리랑\n
	아라리요`

	interLiteral := "아리랑아리랑\n아라리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)

	interLiteral = "안바뀌나?"
	fmt.Println(interLiteral)

	var i int = 100
	var u uint = uint(i)
	var f float32 = float32(i)
	println(u, f)

	str := "ABC"
	bytes := []byte(str)
	str2 := string(bytes)
	println(bytes, str2)
}
