package main

import (
	"fmt"

	"github.com/secr3t/robotgo-without-bitmap"
)

func colorPicker() {
	m := robotgo.AddEvent("mleft")
	if m {
		x, y := robotgo.GetMousePos()
		fmt.Println("mouse pos: ", x, y)

		clo := robotgo.GetPixelColor(x, y)
		fmt.Println("color: #", clo)

		// clipboard
		robotgo.WriteAll("#" + clo)
	}
}

func main() {
	fmt.Println("color picker...")

	for {
		colorPicker()
	}
}
