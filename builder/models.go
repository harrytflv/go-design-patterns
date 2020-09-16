package builder

// VehicleProduct ...
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// BuildProcess ...
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() *VehicleProduct
}
