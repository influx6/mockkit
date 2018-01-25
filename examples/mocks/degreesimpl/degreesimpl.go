package degreesimpl

// DegreesImpl defines a lego block struct which services the methods defined
// by the Degrees interface as assignable fields. These functions are called
// when the respective functions of the struct is called.
type DegreesImpl struct {
	StatFunc func() float64
}

// Stat implements the Degrees.Stat method for Degrees interface.
// It calls the DegreesImpl.StatFunc field underneath.
func (impl DegreesImpl) Stat() float64 {
	ret1 := impl.StatFunc()
	return ret1
}
