// A collection of simple functions for go
package main

import "os"

// Does the if err != nil { panic (err) } for a function that returns a result and an error
func Check[T any](v T, err error) T {
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

// Accumulates the content of a channel into an array
func Accumulate[T any](c <-chan T) chan []T {
	out := make(chan []T)
	go func() {
		defer close(out)
		var a []T
		for e := range c {
			a = append(a, e)
		}
		out <- a
	}()
	return out
}

// Generic interface that accepts any number
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Min function that takes any number
func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max function that takes any number
func Max[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}
