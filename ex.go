package ex

import "github.com/u6du/zerolog/warn"

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func Warn(err error) {
	if err != nil {
		warn.Err(err)
	}
}

func Call(f func() error) {
	Warn(f())
}

type close interface {
	Close() error
}

func Close(c close) {
	Call(c.Close)
}
