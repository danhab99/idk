// A collection of simple functions for go
package idk

// Does the if err != nil { panic (err) } for a function that returns a result and an error
func Check[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Excatly like Check but one less character to type
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Does the if err != nil { panic (err) } for a function that returns an error
func Check0(err error) {
	if err != nil {
		panic(err)
	}
}

// Exactly like Check0 but one less character to type
func Must0(err error) {
	if err != nil {
		panic(err)
	}
}
