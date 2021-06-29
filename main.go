package main

import (
	"fmt"
	"os"
)

const project = "go-template"

func main() {
	fmt.Fprintf(os.Stderr, "hello, %s!\n", project)
}
