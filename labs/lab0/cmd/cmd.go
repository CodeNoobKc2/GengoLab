package main

import (
	"fmt"
	"k8s.io/gengo/args"
	"k8s.io/gengo/namer"
	"lab/constvar"
	"lab/labs/lab0/pkg"
	"os"
)

const (
	labNo = 0
)

func NameSystems() namer.NameSystems {
	return namer.NameSystems{
		"public":  namer.NewPublicNamer(0),
		"private": namer.NewPrivateNamer(0),
		"raw":     namer.NewRawNamer("", nil),
	}
}

func main() {
	arguments := args.Default()
	arguments.InputDirs = []string{"../../../pkg0", "../../../pkg1"}
	arguments.OutputBase = constvar.LabResultPath(labNo)

	if err := arguments.Execute(
		NameSystems(),
		"public",
		pkg.Packages,
	); err != nil {
		fmt.Errorf("lab 0 failed: %s", err.Error())
		os.Exit(1)
	}
}
