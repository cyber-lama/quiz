package exceptions

type ConnectionDBErr struct {
	*BaseErr
}

func NewConnectionErr(err error) ConnectionDBErr {
	b := &BaseErr{Msg: "Failed to connect to database", Err: err}
	return ConnectionDBErr{
		b,
	}
}
