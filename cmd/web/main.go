package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s/n", err)
	}
}

func run(args []string) error {
	fmt.Println("Hello World")
	return nil
}
