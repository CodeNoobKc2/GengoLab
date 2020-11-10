package pkg

import (
	"fmt"
	"k8s.io/gengo/args"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/types"
)

func OutputDivider(name string) string {
	return fmt.Sprintf("=========%s=========", name)
}

func IsPassedInPkg(pkg *types.Package) bool {
	return len(pkg.Path) > 0 && pkg.Path[0] == '.'
}

// this method depend on the way user organize the code and the way user pass in the include dir
func GetSourcePkgName(ctx *generator.Context, arguments *args.GeneratorArgs) []string {
	rtn := make([]string, 0)
	for _, p := range ctx.Universe {
		// cause i pass in a relative path
		if IsPassedInPkg(p) {
			rtn = append(rtn, p.Name)
		}
	}
	return rtn
}

func GetSourcePkgComment(ctx *generator.Context, _ *args.GeneratorArgs) map[string][]string {
	rtn := make(map[string][]string)
	for _, p := range ctx.Universe {
		if IsPassedInPkg(p) {
			rtn[p.Name] = p.Comments
		}
	}
	return rtn
}

func GetSourcePkgTypes(ctx *generator.Context, _ *args.GeneratorArgs) map[string][]string {
	rtn := make(map[string][]string)
	for _, p := range ctx.Universe {
		if IsPassedInPkg(p) {
			types := make([]string, 0)
			for tName, _ := range p.Types {
				types = append(types, tName)
			}

			rtn[p.Name] = types
		}
	}
	return rtn
}

func Packages(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
	fmt.Println(OutputDivider("Package Name"))
	for _, name := range GetSourcePkgName(context, arguments) {
		fmt.Println(fmt.Sprintf("Package name: %s", name))
	}

	fmt.Println(OutputDivider("Package comments"))
	for name, comments := range GetSourcePkgComment(context, arguments) {
		fmt.Println(fmt.Sprintf("Package name: %s", name))
		for _, comment := range comments {
			fmt.Println(comment)
		}
	}

	fmt.Println(OutputDivider("Package Types"))
	for name, types := range GetSourcePkgTypes(context, arguments) {
		fmt.Println(fmt.Sprintf("Package name: %s, type: %v", name, types))
	}

	return generator.Packages{}
}
