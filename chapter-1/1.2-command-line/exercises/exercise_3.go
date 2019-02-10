// Experiment to measure the difference in run-time between version 2 version 3 (of the book versions)
package exercises

import (
	"fmt"
	"os"
	"strings"
)

func exercise3_1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = ""
	}
	fmt.Println(s)
}

func exercise3_2() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
