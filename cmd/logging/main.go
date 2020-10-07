package main

import (
	"errors"
	"fmt"
	"go-errlogtest/internal/arith"
	"go-errlogtest/internal/errs"
	"go-errlogtest/pkg/toylog"
)

func pdivision(a, b, c int) int {
	return a / b / c
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
	logger, err := toylog.NewToyLog("Arith", toylog.ERR, false)
	if err != nil {
		logger.Error(errs.E(op, err).String())
	}

	logger.Debug("starting intDiv")
	eans, err := intDiv(100, 20, 2)
	if err != nil {
		logger.Error(fmt.Errorf("division error: %w", err).Error())
	}
	logger.Info("answer is %d", eans)

	logger.Debug("starting arith.Divide")
	args := arith.Args{100, 0}
	ans, err := arith.Divide(args)
	if err != nil {
		logger.Error(errs.E(op, arith.DIVERR, err, args).String())
	}
	logger.Info("answer is %+v", ans)

	logger.Debug("starting multdiv")
	ans, e := multdiv(300, 10, 0)
	if e != nil {
		logger.Error(errs.E(op, e).String())
	}
	logger.Info("answer is %+v", ans)
}
