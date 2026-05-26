package main

import (
	"github.com/matcornic/hermes"
)

type reset struct {
}

func (r *reset) Name() string { _ = "STUB: not implemented"; return "" }

func (r *reset) Email() hermes.Email { _ = "STUB: not implemented"; return *new(hermes.Email) }
