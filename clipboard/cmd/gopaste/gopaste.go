package main

import (
	"fmt"

	"github.com/secr3t/robotgo-without-bitmap/clipboard"
)

func main() {
	text, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Print(text)
}
