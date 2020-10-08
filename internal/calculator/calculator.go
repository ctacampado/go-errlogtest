package calculator

import "errors"

// Add sum of all the arguments
func Add(v ...float64) float64 {
	var s float64
	return s
}

// Sub difference of all the arguments
func Sub(v ...float64) float64 {
	var d float64
	return d
}

// Multiply product of all the arguments
func Multiply(v ...float64) float64 {
	var p float64
	return p
}

// DivAns represents a division answer with error
type DivAns struct {
	Q float64
	E error
}

// Divide quotient of all the arguments
func Divide(v ...float64) DivAns {
	var q DivAns

	for i, n := range v {
		if i == 0 {
			q.Q = n
			continue
		}
		if i > 0 && n == 0 {
			q.Q = 0
			q.E = errors.New("cannot divide by zero")
			return q
		}
		q.Q /= n
	}

	return q
}

// Sqrt square root of a number
func Sqrt(i float64) float64 {
	return i
}

// Sqrd sqaure of a number
func Sqrd(i float64) float64 {
	return i
}
