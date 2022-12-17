package models

import "errors"

type Namespace struct {
	Name   string
	Status interface{}
}

func NewNamespace(name string) (*Namespace, error) {
	namespace := Namespace{
		Name: name,
	}

	err := namespace.validate()
	if err != nil {
		return nil, err
	}

	return &namespace, nil
}

func (n Namespace) validate() error {
	if len(n.Name) > 63 {
		return errors.New("namespace name cannot exceed 63 characters")
	}

	return nil
}
