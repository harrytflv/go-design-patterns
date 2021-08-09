package builder

// ManufacturingDirector ...
type ManufacturingDirector struct {
	builder BuildProcess
}

// Construct ...
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder ...
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}
