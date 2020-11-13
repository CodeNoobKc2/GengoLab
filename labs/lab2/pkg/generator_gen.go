package pkg

import (
	"fmt"
	"io"
	"path"
	"strings"

	"lab/constvar"

	. "github.com/dave/jennifer/jen"
	"k8s.io/gengo/generator"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

const (
	lab2EnableTag = "lab2"
)

func extractCommentLines(t *types.Type) []string {
	comments := append(append([]string{}, t.CommentLines...), t.SecondClosestCommentLines...)
	return comments
}

func checkEnableTag(comments []string) bool {
	if comments == nil || len(comments) == 0 {
		return false
	}

	tagsValue := types.ExtractCommentTags("+", comments)
	_, ok := tagsValue[lab2EnableTag]
	return ok
}

type AliasPkgNamer struct {
	namer namer.Namer
}

func (apn *AliasPkgNamer) Name(t *types.Type) string {
	return ""
}

type LabGenerator struct {
	generator.DefaultGen
	importTracker namer.ImportTracker
	outputPkg     string
}

func (lg *LabGenerator) Namers(ctx *generator.Context) namer.NameSystems {
	return namer.NameSystems{
		"raw": namer.NewRawNamer(lg.outputPkg, lg.importTracker),
	}
}

func (lg *LabGenerator) PackageConsts(ctx *generator.Context) []string {
	return []string{
		`LAB_2_CONST="CONST"`,
	}
}

func (lg *LabGenerator) Filter(ctx *generator.Context, t *types.Type) bool {
	if t == nil {
		return false
	}

	comments := extractCommentLines(t)
	return checkEnableTag(comments)
}

func makeRelativeImportPathAbsolute(importLine string, packageBase string) string {
	splits := strings.Split(importLine, " ")
	if len(splits) < 2 {
		return importLine
	}

	importPath := strings.Trim(splits[1], "\"")
	if len(importPath) > 2 && importPath[0:2] == "./" {
		splits[1] = "\"" + path.Join(packageBase, importPath) + "\""
		return strings.Join(splits, " ")
	}

	return importLine
}

func (lg *LabGenerator) Imports(ctx *generator.Context) []string {
	lines := lg.importTracker.ImportLines()
	for idx, line := range lines {
		lines[idx] = makeRelativeImportPathAbsolute(line, constvar.PACKAGE_BASE)
	}
	return lines
}

func (lg *LabGenerator) GenerateType(ctx *generator.Context, t *types.Type, writer io.Writer) error {
	//fmt.Println(ctx.Namers["raw"].Name(t))
	if strings.Contains(t.Name.Name, "TestPkg1Func") {
		lg.importTracker.AddType(t)
		f := Func().Id("Lab2GeneratedFunc").Params(Id("a").String()).String().Block(
			Qual(t.Name.Package, t.Name.Name).Call(Id("a"), Lit("b"), Lit(1)),
			Return(Lit("test")),
		)
		fmt.Fprintf(writer, "%#v", f)
	}

	return nil
}
