package main

import (
	"errors"
	"fmt"
	"go-errlogtest/internal/arith"
	"go-errlogtest/internal/errs"
	"go-errlogtest/pkg/toylog"
)

var logger = toylog.NewToyLog("Arith", toylog.ERR, false)

func main() {
	op := errs.Op("main")

	logger.Trace("starting intDiv")
	eans, err := intDiv(100, 20, 2)
	if err != nil {
		logger.Error(fmt.Errorf("division error: %w", err).Error())
	}
	logger.Trace("answer is %d", eans)

	logger.Trace("starting arith.Divide")
	args := arith.Args{100, 0}
	ans, err := arith.Divide(args)
	if err != nil {
		logger.Error(errs.E(op, arith.DIVERR, err, args).String())
	}
	logger.Info("answer is %+v", ans)

	logger.Trace("starting multdiv")
	ans, e := multdiv(300, 10, 0)
	if e != nil {
		logger.Error(errs.E(op, e).String())
	}
	logger.Info("answer is %+v", ans)
}

func intDiv(a, b, c int) (int, error) {
	if b == 0 || c == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return a / b / c, nil
}

// (a * b) / c
func multdiv(a, b, c int) (arith.Quotient, *errs.Error) {
	op := errs.Op("multdiv")
	ans := arith.Quotient{}

	pargs := arith.Args{a, b}
	p, err := arith.Multiply(pargs)
	logger.Debug("args:%+v p: %d", pargs, p)
	if err != nil {
		return ans, errs.E(op, arith.MULTERR, pargs, err)
	}

	qargs := arith.Args{p, c}
	ans, err = arith.Divide(qargs)
	logger.Debug("args:%+v p: %d", qargs, ans)
	if err != nil {
		return ans, errs.E(op, arith.DIVERR, qargs, err)
	}
	return ans, nil
}
