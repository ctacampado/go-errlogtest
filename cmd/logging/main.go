package main

import (
	"go-errlogtest/internal/arith"
	"go-errlogtest/internal/errs"
	"go-errlogtest/pkg/toylog"
)

var l = toylog.NewToyLog("Arith", toylog.ERR, false)

func main() {
	op := errs.Op("main")
	eans, err := intDiv(100, 20, 0)
	if err != nil {
		l.Error(errs.E(op, err).String())
	}
	l.Info("%d", eans)

	args := arith.Args{100, 0}
	ans, err := arith.Divide(args)
	if err != nil {
		newe := errs.E(op, err)
		l.Error(newe.String())
	}
	l.Info("%+v", ans)

	ans, err = multdiv(300, 10, 0)
	if err != nil {
		l.Error(errs.E(op, err).String())
	}
	l.Info("%+v", ans)
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
	p, err := arith.Multiply(pargs)
	l.Debug("args:%+v p: %d", pargs, p)
	if err != nil {
		return ans, errs.E(op, err)
	}

	qargs := arith.Args{p, c}
	ans, err = arith.Divide(qargs)
	l.Debug("args:%+v p: %d", qargs, ans)
	if err != nil {
		return ans, errs.E(op, err)
	}
	return ans, nil
}
