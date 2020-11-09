package main

import (
	"flag"
	"fmt"
	"lab/labs/lab0"
	"os"
)

var (
	labNoToGoal = map[int]string{
		0: "pass include dir path to args",
	}

	labNoFunc = map[int]func(){
		0: lab0.Start,
	}

	labNo = flag.Int("lab", -1, "specific which lab to run")
)

func labStartedPrint(no int) {
	fmt.Println(fmt.Sprintf("Lab: %v. Lab Goal: %s", no, labNoToGoal[no]))
}

func labEndPrint(no int) {
	fmt.Println(fmt.Sprintf("Lab: %v complete", no))
}

func main() {
	flag.Parse()

	_, ok := labNoToGoal[*labNo]
	if !ok {
		fmt.Println("no lab was specified")
		os.Exit(0)
	}

	labStartedPrint(*labNo)
	labNoFunc[*labNo]()
	labEndPrint(*labNo)
}
