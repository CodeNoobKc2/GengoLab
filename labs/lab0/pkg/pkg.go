package pkg

import (
	"fmt"
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
)

func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	fmt.Println(fmt.Sprintf("Input dirs: %s", arguments.InputDirs))
	fmt.Println(fmt.Sprintf("Output base: %s", arguments.OutputBase))

	return generator.Packages{}
}
