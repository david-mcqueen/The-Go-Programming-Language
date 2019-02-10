package exercises

import (
	"fmt"
	"os"
)

func main() {
	exercise1()
	exercise2()
}

// Modify the echo program to also print os.Args[0]
func exercise1() {
	fmt.Println("Exercise 1: Modify the echo program to also print os.Args[0]")
	fmt.Println("Called by: ", os.Args[0:1])
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// Modify the echo program to print the inder and value of each of its arguments, one per line
func exercise2() {
	fmt.Println("Exercise 2: Modify the echo program to print the index and value of each of its arguments, one per line")

	var s string
	for i, arg := range os.Args[1:] {
		fmt.Println(i, " ", arg)
	}
	fmt.Println(s)
}
