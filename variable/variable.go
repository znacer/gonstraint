package gonstraint

import (
	"errors"
)

type Variable struct {
	Name   string
	Domain []int
	Value  *int
}

func NewVariable(name string, domain []int) (*Variable, error) {
	if len(domain) == 0 {
		return nil, errors.New("ErrEmptyDomain")
	}
	return &Variable{Name: name, Domain: domain}, nil
}
