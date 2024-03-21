package main

import (
	"fmt"
	"os"

	"github.com/aamirmousavi/dong/cmd"
)

func main() {
	// Todo: add cmd App name
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
