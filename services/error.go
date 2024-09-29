package services

import "errors"

var (
	ErrZeroAmount = errors.New("purchase not be 0")
	ErrRepository = errors.New("repository error")
)
