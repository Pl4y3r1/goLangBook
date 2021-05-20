package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

func main() {
	s, sep := "", ""

	fmt.Println("Testing Echo 1")
	start := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.8fs elapsed \n", time.Since(start).Seconds())

	fmt.Println("Testing Echo 2")
	s = ""
	sep = ""
	start = time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%.8fs elapsed \n", time.Since(start).Seconds())

	fmt.Println("Testing Echo 3")
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%.8fs elapsed \n", time.Since(start).Seconds())
}
