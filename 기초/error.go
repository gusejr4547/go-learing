package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		log.Fatal(err.Error())
		//log.Fatal() 은 메시지를 출력하고 os.Exit(1)을 호출하여 프로그램을 종료한다).
	}
	println(f.Name())

	/*
		_, err := otherFunc()
		switch err.(type) {
		default: // no error
			println("ok")
		case MyError:
			log.Print("Log my error")
		case error:
			log.Fatal(err.Error())
		}
	*/
}
