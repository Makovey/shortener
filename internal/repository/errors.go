package repository

import "errors"

var ErrURLNotFound = errors.New("url is not existed yet")
var ErrURLIsAlreadyExists = errors.New("url is already existed")
