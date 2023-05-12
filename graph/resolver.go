package graph

import "context"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// NewResolver returns a new Resolver struct.
func NewResolver() *Resolver {
	return &Resolver{}
}

// Hello resolver function for the "hello" field.
func (r *Resolver) Hello(ctx context.Context) (string, error) {
	return "world", nil
}
