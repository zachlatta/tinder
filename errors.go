package tinder

import "errors"

var (
	RecsExhausted = errors.New("recs exhausted")
	RecsTimeout   = errors.New("recs timeout")
)
