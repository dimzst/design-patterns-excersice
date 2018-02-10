package duck

import (
	"fmt"

	"github.com/dimzst/design-pattern-exercise/duck/behavior"
)

type Displayer interface {
	Display()
}

type Duck struct {
	Displayer
	quacker behavior.Quacker
	flyer   behavior.Flyer
}

func (d *Duck) PerformFLy() {
	d.flyer.Fly()
}

func (d *Duck) PerformQuack() {
	d.quacker.Quack()
}

func (d *Duck) SetFlyer(f behavior.Flyer) {
	d.flyer = f
}

func (d *Duck) SetQuacker(q behavior.Quacker) {
	d.quacker = q
}

func (d *Duck) Swim() {
	fmt.Println("All ducks float, even decoys!")
}

type MallardDuck struct {
	Duck
}

func (m *MallardDuck) Display() {
	fmt.Println("I'm a real Mallard Duck")
}

func NewMallardDuck() MallardDuck {
	return MallardDuck{
		Duck{
			quacker: &behavior.Quack{},
			flyer:   &behavior.Wings{},
		},
	}
}

type ModelDuck struct {
	Duck
}

func (m *ModelDuck) Display() {
	fmt.Println("I'm a model duck")
}

func NewModelDuck() ModelDuck {
	return ModelDuck{
		Duck{
			quacker: &behavior.Quack{},
			flyer:   &behavior.NoFly{},
		},
	}
}
