package main

import (
	"github.com/matcornic/hermes"
)

type inviteCode struct {
}

func (w *inviteCode) Name() string { _ = "STUB: not implemented"; return "" }

func (w *inviteCode) Email() hermes.Email { _ = "STUB: not implemented"; return *new(hermes.Email) }
