package exceptions

type ConnectionDBErr struct {
	*BaseErr
}

func NewConnectionErr(err error) ConnectionDBErr {
	b := &BaseErr{Msg: "Не удалось подключиться к базе данных", Err: err}
	return ConnectionDBErr{
		b,
	}
}
