package foobar

import "fmt"

type FooBar struct {
}

func NewFooBarService() *FooBar {
	return &FooBar{}
}

func (s FooBar) DoSomeDummyAction() error {

	fmt.Println("asd")
	return nil
}
