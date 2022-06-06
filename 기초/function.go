package main

func main() {
	msg := "Hello"
	say(msg)

	say_ref(&msg)
	println(msg)

	say_vari("This", "is", "a", "book")
	say_vari("Hi")

	total := sum(1, 7, 3, 5, 9)
	println(total)
	count, total := sum2(1, 7, 3, 5, 9)
	println(count, total)
}

func say(msg string) {
	println(msg)
}

func say_ref(msg *string) {
	println(*msg)
	*msg = "Changed" //메세지 변경
}

func say_vari(msg ...string) {
	for _, s := range msg {
		println(s)
	}
}

func sum(nums ...int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

func sum2(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
