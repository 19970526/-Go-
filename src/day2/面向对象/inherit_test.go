package class_test

import (
	"fmt"
	"testing"
)

type Fluit struct{}

func (f *Fluit) Say() {
	fmt.Println("I am a kind of fluit")
}

type Apple struct {
	f *Fluit
}

func (a *Apple) Say() {
	fmt.Println("I am an apple")
}

func TestInherit(t *testing.T) {
	obj := new(Apple)
	obj.Say()
}
