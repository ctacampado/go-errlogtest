package arith

import (
	"go-errlogtest/internal/errs"
	"log"
)

// Arg type
type Arg int

// Args type
type Args []int

// Len returns the number of arguments
func (a Args) Len() int {
	return len(a)
}

// Quotient represents an answer to a custom division operation
type Quotient struct {
	Q int
	R int
}

const (
	// DIVERR division error
	DIVERR errs.ErrType = "DIVERR"
	// MULTERR multiplication error
	MULTERR errs.ErrType = "MULTERR"
)

// Print prints out all values inside args
func (a *Args) Print() {
	for _, v := range *a {
		log.Println(v)
	}
}

// Divide func
func Divide(args Args) (Quotient, *errs.Error) {
	if len(args) <= 1 {
		return Quotient{}, errs.New(
			"arith.Divide",
			"needs more than 1 argument",
			DIVERR,
			args,
		)
	}

	var q Quotient
	for i, v := range args {
		if i+1 >= len(args) {
			break
		}
		if args[i+1] == 0 {
			return Quotient{}, errs.New(
				"arith.Divide",
				"cannot divide by zero",
				DIVERR,
				args,
			)
		}
		q.Q = v
		q.Q /= args[i+1]
		q.R %= args[i+1]
	}

	return q, nil
}

// Multiply func
func Multiply(args Args) (int, *errs.Error) {
	if len(args) <= 1 {
		return 0, errs.New(
			"arith.Multiply",
			"needs more than 1 argument",
			MULTERR,
			args,
		)
	}

	p := 1
	for _, v := range args {
		p *= v
	}

	return p, nil
}
