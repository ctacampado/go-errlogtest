package arith

import (
	"errors"
	"log"
)

// Arg type
type Arg int

// Args type
type Args []int

// Quotient represents an answer to a custom division operation
type Quotient struct {
	Q int
	R int
}

// ERRTYPE type
type ERRTYPE int

const (
	// DIVERR division error
	DIVERR ERRTYPE = iota
	// MULTERR multiplication error
	MULTERR
)

// Print prints out all values inside args
func (a *Args) Print() {
	for _, v := range *a {
		log.Println(v)
	}
}

// Divide func
func Divide(args Args) (Quotient, error) {
	if len(args) <= 1 {
		return Quotient{}, errors.New("needs more than 1 argument")
	}

	var q Quotient
	for i, v := range args {
		if i+1 >= len(args) {
			break
		}
		if args[i+1] == 0 {
			return Quotient{}, errors.New("cannot divide by 0")
		}
		q.Q = v
		q.Q /= args[i+1]
		q.R %= args[i+1]
	}

	return q, nil
}

// Multiply func
func Multiply(args Args) (int, error) {
	if len(args) <= 1 {
		return 0, errors.New("needs more than 1 argument")
	}

	p := 1
	for _, v := range args {
		p *= v
	}

	return p, nil
}
