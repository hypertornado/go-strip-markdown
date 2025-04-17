package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	fmt.Print(stripmd.Strip(os.Args[1]))
}
