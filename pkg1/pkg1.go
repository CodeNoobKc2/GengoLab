// pkg 1 refs a sub package pkg00
package pkg1

import (
	"fmt"
	"lab/pkg0/pkg00"
	aliasPkg00 "lab/pkg0/pkg00"
	"sync"
)

var (
	lock = &sync.Mutex{}
	Lock = &sync.Mutex{}
)

type Pkg1 struct {
	field  int
	Field0 int
	Field1 string
}

func NewPkg1(field int) Pkg1 {
	return Pkg1{
		field:  field,
		Field0: 1,
		Field1: "test",
	}
}

func (p *Pkg1) func0() {
	fmt.Print(p.field)
}

func (p *Pkg1) Func(output string) string {
	return output
}

// +lab2
// TestPkg0Func
func TestPkg1Func(arg0, arg1 string, arg2 int) (string, error) {
	return fmt.Sprintf("%s%s%v", arg0, arg1, arg2), nil
}

func RefPkg00Struct() *pkg00.Pkg00 {
	pkg00.TestCaseTransitiveIncomingImports()
	return &pkg00.Pkg00{}
}

func RefAliasPkg00Struct() *aliasPkg00.Pkg00 {
	return &aliasPkg00.Pkg00{}
}

// testPkg0Func
func testPkg1Func() (string, error) {
	return "test", nil
}
