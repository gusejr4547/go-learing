package main

func main() {
	var a int
	var f float32 = 11.
	println(a, f)

	a = 10
	f = 12.0
	println(a, f)

	var i, j, k int = 1, 2, 3
	println(i, j, k)

	var v = "Hi"
	println(v)

	const c int = 10
	const s string = "Hi"
	println(c, s)

	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)
	println(Visa, Master, Amex)
	const (
		Apple  = iota //0
		Grape         //1
		Orange        //2
	)
	println(Apple, Grape, Orange)
}
