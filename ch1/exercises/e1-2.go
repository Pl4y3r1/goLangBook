package main

import (
	"fmt"
	"os"
)

func main() {
	s := ""
	for i := 1; i < len(os.Args); i++{
		s = ": " + os.Args[i] + "\n"
		fmt.Printf("%v" + s, i)
	}
}
