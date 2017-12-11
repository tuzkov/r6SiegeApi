package r6

import (
	"github.com/pkg/errors"
)

var (
	// ErrLoginIncorrect ...
	ErrLoginIncorrect = errors.New("неверный логин/пароль")
)