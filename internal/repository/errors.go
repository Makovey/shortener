package repository

import "errors"

// Ошибки, которые могут возникать в репозитории
var (
	ErrURLNotFound        = errors.New("url is not existed yet")
	ErrURLIsAlreadyExists = errors.New("url is already existed")
)
