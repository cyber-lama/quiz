package exceptions

import (
	"fmt"
)

type BaseErr struct {
	Msg string
	Err error
}

func (e *BaseErr) Error() string {
	mess := fmt.Sprintf("%s: %s", e.Msg, e.Err)
	return mess
}
