package degreesmock

import (
	"errors"
	"time"

	"runtime"
)

// MethodCallForStat defines a type which holds meta-details about the giving calls associated
// with the Degrees.Stat() method.
type MethodCallForStat struct {
	When  time.Time
	Start time.Time
	End   time.Time

	// Details of panic if such occurs.
	PanicStack []byte
	PanicError interface{}

	// Argument values.

	// Return values.

	Ret1 float64
}

// MatchArguments returns true/false if provider other(MethodCallForStat) arguments
// values match this existing argument values.
func (me MethodCallForStat) MatchArguments(other MethodCallForStat) bool {

	return true
}

// DegreesMock defines a type which implements a struct with the
// methods for the Degrees as fields which allows you provide implementations of
// these functions to provide flexible testing.
type DegreesMock struct {
	StatMethodCalls []MethodCallForStat
}

// Stat implements the Degrees.Stat() method for the Degrees.
func (impl *DegreesMock) Stat() (MethodCallForStat, error) {
	var caller MethodCallForStat

	caller.When = time.Now()
	caller.Start = caller.When

	var found bool
	for _, possibleCall := range impl.StatMethodCalls {
		if possibleCall.MatchArguments(caller) {
			found = true

			caller.Ret1 = possibleCall.Ret1

			caller.PanicError = possibleCall.PanicError
			caller.PanicStack = possibleCall.PanicStack
			break
		}
	}

	caller.End = time.Now()
	if found {
		return caller, nil
	}

	return caller, errors.New("no matching response found")
}

// DegreesSnitch defines a type which implements a struct with the
// methods for the Degrees as fields which allows you provide implementations of
// these functions to provide flexible testing.
type DegreesSnitch struct {
	StatMethodCalls []MethodCallForStat
	StatFunc        func() float64
}

// Stat implements the Degrees.Stat() method for the Degrees.
func (impl *DegreesSnitch) Stat() float64 {
	var caller MethodCallForStat

	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1000)
			trace = trace[:runtime.Stack(trace, true)]

			caller.PanicError = err
			caller.PanicStack = trace
		}

		caller.End = time.Now()
		impl.StatMethodCalls = append(impl.StatMethodCalls, caller)
	}()

	caller.When = time.Now()
	caller.Start = caller.When

	ret1 := impl.StatFunc()

	caller.Ret1 = ret1

	return ret1
}
