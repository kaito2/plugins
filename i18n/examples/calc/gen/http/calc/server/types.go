// Code generated by goa v3.0.2, DO NOT EDIT.
//
// calc HTTP server types
//
// Command:
// $ goa gen goa.design/plugins/v3/i18n/examples/calc/design

package server

import (
	calc "goa.design/plugins/v3/i18n/examples/calc/gen/calc"
)

// NewAddPayload builds a calc service add endpoint payload.
func NewAddPayload(a int, b int) *calc.AddPayload {
	return &calc.AddPayload{
		A: a,
		B: b,
	}
}