package builder

// NewCarBuilder ...
func NewCarBuilder() BuildProcess {
	return &carBuilder{
		v: &VehicleProduct{},
	}
}

type carBuilder struct {
	v *VehicleProduct
}

func (c *carBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *carBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *carBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *carBuilder) GetVehicle() *VehicleProduct {
	return c.v
}
