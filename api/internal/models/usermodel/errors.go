package usermodel

import "errors"

var (
	ErrCreateUser         = errors.New("не получилось создать пользователя")
	ErrInvalidAccessToken = errors.New("не верный токен валидации")
)
