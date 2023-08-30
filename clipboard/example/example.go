package main

import (
	"log"

	"github.com/secr3t/robotgo-without-bitmap/clipboard"
)

func main() {
	clipboard.WriteAll("日本語")
	text, err := clipboard.ReadAll()
	if err != nil {
		log.Println("clipboard read all error: ", err)
	} else {
		if text != "" {
			log.Println("text is: ", text)
			// Output: 日本語
		}
	}
}
