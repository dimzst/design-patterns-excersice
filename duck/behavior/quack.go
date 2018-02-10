package behavior

import (
	"fmt"
)

type Quacker interface {
	Quack()
}

type Quack struct{}

func (q *Quack) Quack() {
	fmt.Println("Quack")
}

type Mute struct{}

func (m *Mute) Quack() {
	fmt.Println("<<Silence>>")
}

type Squeak struct{}

func (s *Squeak) Quack() {
	fmt.Println("Squeak")
}
