package builder

import (
	"testing"
)

// TestBuilderPattern ...
func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := &ManufacturingDirector{}

	carBuilder := NewCarBuilder()
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("?")
	}

	if car.Structure != "Car" {
		t.Errorf("?")
	}

	if car.Seats != 5 {
		t.Errorf("?")
	}

	bikeBuilder := NewBikeBuilder()

	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	motorbike := bikeBuilder.GetVehicle()
	motorbike.Seats = 1

	if motorbike.Wheels != 2 {
		t.Errorf("?")
	}

	if motorbike.Structure != "Motorbike" {
		t.Errorf("?")
	}
}
