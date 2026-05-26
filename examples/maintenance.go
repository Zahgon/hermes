package main

import (
	"github.com/matcornic/hermes"
)

type maintenance struct {
}

func (w *maintenance) Name() string { _ = "STUB: not implemented"; return "" }

func (w *maintenance) Email() hermes.Email { _ = "STUB: not implemented"; return *new(hermes.Email) }
