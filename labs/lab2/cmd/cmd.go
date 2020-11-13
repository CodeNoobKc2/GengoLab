package main

import (
	"fmt"
	"k8s.io/gengo/args"
	"lab/generator"
	"lab/labs/lab2/pkg"
	"os"
)

const (
	labNo = 2
)

func main() {
	arguments := args.Default()
	arguments.InputDirs = []string{"./pkg0", "./pkg1", "./pkg2", "./pkg0/pkg00"}
	arguments.OutputBase = "./"
	arguments.OutputPackagePath = "labs/lab2/result"

	if err := arguments.Execute(
		generator.NameSystems(),
		generator.DefaultNameSystem(),
		pkg.Packages,
	); err != nil {
		fmt.Errorf("failed: %s", err.Error())
		os.Exit(1)
	}
}
