// this is pkg0 test
package pkg0

import (
	"fmt"
	"sync"
)

var (
	lock = &sync.Mutex{}
	Lock = &sync.Mutex{}
)

// +lab2
type Pkg0 struct {
	field  int
	Field0 int
	Field1 string
}

func NewPkg0(field int) Pkg0 {
	return Pkg0{
		field:  field,
		Field0: 1,
		Field1: "test",
	}
}

func (p *Pkg0) func0() {
	fmt.Print(p.field)
}

func (p *Pkg0) Func(output string) string {
	return output
}

// TestPkg0Func
func TestPkg0Func(arg0, arg1 string, arg2 int) (string, error) {
	return fmt.Sprintf("%s%s%v", arg0, arg1, arg2), nil
}

func testPkg0Func() (string, error) {
	return "test", nil
}
