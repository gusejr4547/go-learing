package main

import "fmt"

func main() {
	var idMap map[int]string //[key]value
	idMap = make(map[int]string)
	fmt.Println(idMap)

	//리터럴을 사용한 초기화
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "FaceBook",
	}
	fmt.Println(tickers)

	var m map[int]string

	m = make(map[int]string)
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	fmt.Println(str)

	noData := m[999]
	fmt.Println(noData)

	delete(m, 777)
	fmt.Println(m)

	val, exists := tickers["MSFT"]
	if !exists {
		fmt.Println("No MSFT ticker")
	} else {
		fmt.Println(val)
	}

	myMap := map[string]string{
		"A": "Apple",
		"B": "Banana",
		"C": "Charlie",
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range myMap {
		fmt.Println(key, val)
	}
}
