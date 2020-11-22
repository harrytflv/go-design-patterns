package composite

import "testing"

func TestComposite(t *testing.T) {
	swimmer := CompositeSwimmerA{
		MySwim: Swim,
	}

	swimmer.MyAthelete.Train()
	swimmer.MySwim()

	fish := Shark{
		Swim: Swim,
	}

	fish.Eat()
	fish.Swim()

	swimmer2 := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	swimmer2.Train()
	swimmer2.Swim()
}
