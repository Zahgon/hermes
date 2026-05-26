package main

import (
	"github.com/matcornic/hermes"
)

type welcome struct {
}

func (w *welcome) Name() string { _ = "STUB: not implemented"; return "" }

func (w *welcome) Email() hermes.Email { _ = "STUB: not implemented"; return *new(hermes.Email) }
