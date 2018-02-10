package behavior

import (
	"fmt"
)

type Flyer interface {
	Fly()
}

type Wings struct{}

func (w *Wings) Fly() {
	fmt.Println("I'm flying!!")
}

type NoFly struct{}

func (n *NoFly) Fly() {
	fmt.Println("I can't fly!")
}

type Rocket struct{}

func (n *Rocket) Fly() {
	fmt.Println("I'm flying with a rocket!")
}
