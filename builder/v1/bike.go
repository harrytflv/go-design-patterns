package builder

// NewBikeBuilder ...
func NewBikeBuilder() BuildProcess {
	return &bikeBuilder{
		v: &VehicleProduct{},
	}
}

type bikeBuilder struct {
	v *VehicleProduct
}

func (b *bikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *bikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *bikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *bikeBuilder) GetVehicle() *VehicleProduct {
	return b.v
}
