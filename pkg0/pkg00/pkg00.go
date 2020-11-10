// a pkg with in the pkg0
package pkg00

import "os"

type Pkg00 struct {
	field  int
	Field0 int
	Field1 string
}

func TestCaseTransitiveIncomingImports() {
	// import os package to test transitive incoming imports
	os.Exit(0)
}
