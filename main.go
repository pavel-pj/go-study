package main

import "strconv"
import "fmt"

func main () {
	fmt.Println("Hellow world")
	fmt.Println(string(66))

	count := 66
	fmt.Println("New version: " + strconv.Itoa(count))

	pi := 3.14
	msg := "Pi is approximately " + strconv.FormatFloat(pi, 'f', 9, 32)

	fmt.Println(msg)

	name := "Иваныч"
	fmt.Println(fmt.Sprintf("Имя : %s", name  ))

	fmt.Println(len("go"))
	fmt.Println(len("го"))

}