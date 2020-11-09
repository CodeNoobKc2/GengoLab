package pkg1

import (
"fmt"
"sync"
)

var (
	lock =&sync.Mutex{}
	Lock =&sync.Mutex{}
)

type Pkg1 struct{
	field int
	Field0 int
	Field1 string
}

func NewPkg1(field int)Pkg1{
	return Pkg1{
		field:  field,
		Field0: 1,
		Field1: "test",
	}
}

func (p *Pkg1) func0(){
	fmt.Print(p.field)
}

func (p *Pkg1) Func(output string)string{
	return output
}

// TestPkg0Func
func TestPkg1Func(arg0,arg1 string,arg2 int)(string,error) {
	return fmt.Sprintf("%s%s%v",arg0,arg1,arg2),nil
}

// testPkg0Func
func testPkg1Func()(string,error){
	return "test",nil
}
