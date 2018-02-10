package main

import (
	"github.com/dimzst/design-pattern-exercise/duck"
	"github.com/dimzst/design-pattern-exercise/duck/behavior"
)

func main() {
	maDuck := duck.NewMallardDuck()
	maDuck.Display()
	maDuck.PerformFLy()
	maDuck.PerformQuack()

	moDuck := duck.NewModelDuck()
	moDuck.Display()
	moDuck.PerformFLy()
	moDuck.SetFlyer(&behavior.Rocket{})
	moDuck.PerformFLy()
}
