package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 17

	if number%2 == 0 {
		fmt.Println("Even number")
	} else if number%2 == 1 {
		fmt.Println("Odd number")
	}
	str := strconv.Itoa(number)
	if len(str) > 1 {
		fmt.Println("The number has more than 1 digit")
	} else if len(str) < 2 {
		fmt.Println("The number has 1 digit")
	}
	//Insert your code here
	//Hint if a:= ?? ; condition {  }
}
