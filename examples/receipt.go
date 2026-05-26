package main

import (
	"github.com/matcornic/hermes"
)

type receipt struct {
}

func (r *receipt) Name() string { _ = "STUB: not implemented"; return "" }

func (r *receipt) Email() hermes.Email { _ = "STUB: not implemented"; return *new(hermes.Email) }
