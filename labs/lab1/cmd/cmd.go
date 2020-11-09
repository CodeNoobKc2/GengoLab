package main

import (
	"fmt"
	"os"

	"lab/constvar"
	"lab/generator"
	"lab/labs/lab1/pkg"

	"k8s.io/gengo/args"
)

const (
	labNo = 1
)

func main() {
	arguments := args.Default()
	arguments.InputDirs = []string{"../../../pkg0", "../../../pkg1"}
	arguments.OutputBase = constvar.LabResultPath(labNo)

	if err := arguments.Execute(
		generator.NameSystems(),
		generator.DefaultNameSystem(),
		pkg.Packages,
	); err != nil {
		fmt.Errorf("lab 0 failed: %s", err.Error())
		os.Exit(1)
	}
}
