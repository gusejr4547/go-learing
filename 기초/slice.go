package main

import "fmt"

func main() {
	var a []int
	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println(a)

	s := make([]int, 5, 10)
	println(len(s), cap(s)) // len 5, cap 10

	fmt.Println(s)

	var s1 []int

	if s1 == nil {
		println("Nil Slice")
	}
	println(len(s), cap(s)) // 모두 0

	s2 := []int{0, 1, 2, 3, 4, 5}
	s2 = s2[2:5]
	fmt.Println(s2)

	s3 := []int{0, 1, 2, 3, 4, 5}
	s3 = s3[2:5]
	s3 = s3[1:]
	fmt.Println(s3)

	s4 := []int{0, 1}
	s4 = append(s4, 2)
	s4 = append(s4, 3, 4, 5)
	fmt.Println(s4)

	// len=0, cap=3 인 슬라이스
	sliceA := make([]int, 0, 3)

	// 계속 한 요소씩 추가
	for i := 1; i <= 15; i++ {
		sliceA = append(sliceA, i)
		// 슬라이스 길이와 용량 확인
		fmt.Println(len(sliceA), cap(sliceA))
	}

	fmt.Println(sliceA) // 1 부터 15 까지 숫자 출력

	sliceB := []int{1, 2, 3}
	sliceC := []int{4, 5, 6}

	sliceB = append(sliceB, sliceC...)
	//sliceA = append(sliceA, 4, 5, 6)

	fmt.Println(sliceB) // [1 2 3 4 5 6] 출력

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)               // [0 1 2 ] 출력
	println(len(target), cap(target)) // 3, 6 출력
}
