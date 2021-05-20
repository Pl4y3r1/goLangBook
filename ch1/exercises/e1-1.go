package main

import (
	"strings"
	"os"
	"fmt"
)

func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
