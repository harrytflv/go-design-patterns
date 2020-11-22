package composite

import (
	"fmt"
)

type Athlete struct {
}

func (a *Athlete) Train() {
	fmt.Println("Training!")
}

type CompositeSwimmerA struct {
	MyAthelete Athlete
	MySwim     func()
}

func Swim() {
	fmt.Println("Swimming!")
}

type Animal struct {
}

func (a *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct {
}

func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}
