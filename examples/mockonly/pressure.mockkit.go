package mockonly

// PressureImpl defines a lego block struct which services the methods defined
// by the Pressure interface as assignable fields. These functions are called
// when the respective functions of the struct is called.
type PressureImpl struct {
	StatFunc func() float64
}

// Stat implements the Pressure.Stat method for Pressure interface.
// It calls the PressureImpl.StatFunc field underneath.
func (impl PressureImpl) Stat() float64 {
	ret1 := impl.StatFunc()
	return ret1
}
