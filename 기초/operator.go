package main

func main() {
	var a, b = 12, 5
	c := (a + b) / 5
	println(c)

	println(true && false)

	println((a & b) << 5)

	a = 100
	println(a)
	a *= 10
	println(a)
	a >>= 2
	println(a)
	a |= 1
	println(a)

	var k int = 10
	var p = &k
	println(p)
	println(*p)
	k += 5
	println(*p)
}
