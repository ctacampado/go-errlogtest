package errs

import (
	"encoding/json"
	"fmt"
	"go-errlogtest/internal/arith"
)

// Op is the function name where the error happened
type Op string

// Error is a custom error type that enables us to
// generate rich error messages for easier debugging,
// logging, and querying.
type Error struct {
	Op      Op
	OpStack []Op
	Kind    arith.ERRTYPE
	Err     error
	Args    arith.Args
}

// Print the error
func (e Error) String() string {
	var eType string
	switch e.Kind {
	case arith.DIVERR:
		eType = "DIVERR"
	case arith.MULTERR:
		eType = "MULTERR"
	}

	jStack, _ := json.Marshal(e.OpStack)
	jArgs, _ := json.Marshal(e.Args)

	return fmt.Sprintf(
		"{\"stack\":%v,\"type\":\"%s\",\"err\":\"%s\",\"args\":%v}",
		string(jStack),
		eType,
		e.Err.Error(),
		string(jArgs),
	)
}

// StackTrace contains operation stack
func (e Error) StackTrace() string {
	return fmt.Sprintf("%+v", e.OpStack)
}

// E returns an errs.Error
func E(args ...interface{}) *Error {
	e := &Error{}
	if len(args) == 0 {
		return e
	}
	for _, arg := range args {
		switch a := arg.(type) {
		case Op:
			e.Op = a
			e.OpStack = append(e.OpStack, e.Op)
		case []Op:
			e.OpStack = a
		case arith.ERRTYPE:
			e.Kind = a
		case error:
			e.Err = a
		case arith.Args:
			e.Args = a
		case *Error:
			e = E(e.OpStack, a.Op, a.Kind, a.Err, a.Args)
		default:
			panic("unknow type of arg")
		}
	}
	return e
}
