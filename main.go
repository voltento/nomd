package main

import (
	"fmt"
	"github.com/writeas/go-strip-markdown"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}

	fmt.Printf("\n\n--------------------------------------\n\n")
	fmt.Println(stripmd.Strip(os.Args[1]))
}
