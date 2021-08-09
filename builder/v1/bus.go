package builder

// NewBusBuilder ...
func NewBusBuilder() BuildProcess {
	return &bikeBuilder{
		v: &VehicleProduct{},
	}
}

type busBuilder struct {
	v *VehicleProduct
}

func (b *busBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4 * 2
	return b
}

func (b *busBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

func (b *busBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

func (b *busBuilder) GetVehicle() *VehicleProduct {
	return b.v
}
