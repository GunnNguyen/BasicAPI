package errors

type Error struct {
	Code   Code
	Detail string
}

func (e Error) Error() string {
	msg := e.Code.String()
	if len(e.Detail) > 0 {
		msg += ": " + e.Detail
	}
	return msg
}

func ErrorIs(err error, target Code) bool {
	if e, ok := err.(Error); ok {
		return e.Code == target
	}
	return false
}
