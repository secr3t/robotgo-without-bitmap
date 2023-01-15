// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package main

import (
	"fmt"
	"github.com/secr3t/robotgo-without-bitmap"
)

func aRobotgo() {
	abool := robotgo.ShowAlert("test", "robotgo")
	if abool {
		fmt.Println("ok@@@", "ok")
	}

	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)

	robotgo.MoveMouse(x, y)
	robotgo.MoveMouse(100, 200)

	robotgo.MouseToggle("up")

	for i := 0; i < 1080; i += 1000 {
		fmt.Println(i)
		robotgo.MoveMouse(800, i)
	}

	robotgo.TypeStr("Hello World")
	// robotgo.KeyTap("a", "control")
	robotgo.KeyTap("f1", "control")
	// robotgo.KeyTap("enter")
	// robotgo.KeyToggle("enter", "down")
	robotgo.TypeStr("en")

	// robotgo.MouseClick()
	robotgo.ScrollMouse(10, "up")
}

func main() {
	aRobotgo()
}
