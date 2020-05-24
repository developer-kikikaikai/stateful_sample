package statemachine

import "github.com/bykof/stateful"

var (
	BEGIN    = stateful.DefaultState("BEGIN")
	InCart   = stateful.DefaultState("InCart")
	Ordered  = stateful.DefaultState("Ordered")
	Canceled = BEGIN
	Shipped  = BEGIN
)

type Product struct {
	Fee int
}
