package main

import (
	"errors"
	"fmt"
	"go-errlogtest/internal/arith"
	"go-errlogtest/internal/errs"
	"log"
)

func pdivision(a, b, c int) int {
	return a / b / c
}

func intDiv(a, b, c int) (int, *errs.Error) {
	if b == 0 || c == 0 {
		err := errs.New(
			"intDiv",
			"can't divide by zero",
			arith.DIVERR,
			arith.Args{a, b, c},
		)
		return 0, err
	}
	return a / b / c, nil
}

// (a * b) / c
func multdiv(a, b, c int) (arith.Quotient, *errs.Error) {
	op := errs.Op("multdiv")
	ans := arith.Quotient{}

	pargs := arith.Args{a, b}
	//pargs := arith.Args{a}
	p, err := arith.Multiply(pargs)
	if err != nil {
		return ans, errs.E(op, arith.MULTERR, pargs, err)
	}

	qargs := arith.Args{p, c}
	ans, err = arith.Divide(qargs)
	if err != nil {
		return ans, errs.E(op, arith.DIVERR, qargs, err)
	}
	return ans, nil
}

func main() {
	op := errs.Op("main")
	eans, e := intDiv(100, 20, 1)
	if e != nil {
		log.Println(fmt.Errorf("division error: %w", errors.New(e.Err)))
	} else {
		fmt.Printf("[%s] answer: %d\n", op, eans)
	}

	args := arith.Args{100, 50}
	ans, e := arith.Divide(args)
	if e != nil {
		log.Println(errs.E(op, arith.DIVERR, e, args).String())
	} else {
		log.Printf("[%s] ans: %+v\n", op, ans)
	}

	ans, err := multdiv(300, 10, 0)
	if err != nil {
		log.Println(errs.E(op, err).String())
	} else {
		log.Printf("[%s] ans: %+v\n", op, ans)
	}
}
