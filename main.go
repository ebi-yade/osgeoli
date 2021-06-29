package main

import (
	"fmt"
	"os"
)

const project = "osgeoli"

func main() {
	fmt.Fprintf(os.Stderr, "hello, %s!\n", project)
}
