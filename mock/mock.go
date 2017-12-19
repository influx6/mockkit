package mock

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gokit/mockkit/static"
	"github.com/influx6/faux/fmtwriter"
	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
)

// ImplGen generates a mock package with implementation for giving type.
func ImplGen(toDir string, an ast.AnnotationDeclaration, itr ast.InterfaceDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	interfaceName := itr.Object.Name.Name
	packageName := fmt.Sprintf("%simpl", strings.ToLower(interfaceName))
	packageMockName := fmt.Sprintf("%smock", strings.ToLower(interfaceName))
	methods := itr.Methods(&pkgDeclr)

	imports := make(map[string]string, 0)

	for _, method := range methods {
		// Retrieve all import paths for arguments.
		func(args []ast.ArgType) {
			for _, argument := range args {
				if argument.Import2.Path != "" {
					imports[argument.Import2.Path] = argument.Import2.Name
				}
				if argument.Import.Path != "" {
					imports[argument.Import.Path] = argument.Import.Name
				}
			}
		}(method.Args)

		// Retrieve all import paths for returns.
		func(args []ast.ArgType) {
			for _, argument := range args {
				if argument.Import2.Path != "" {
					imports[argument.Import2.Path] = argument.Import2.Name
				}
				if argument.Import.Path != "" {
					imports[argument.Import.Path] = argument.Import.Name
				}
			}
		}(method.Returns)
	}

	var implImports []gen.ImportItemDeclr
	implImports = append(implImports,
		gen.Import(pkg.Path, ""),
	)

	var mockImports []gen.ImportItemDeclr
	mockImports = append(mockImports,
		gen.Import("time", ""),
		gen.Import("runtime", ""),
		gen.Import("github.com/influx6/faux/reflection", ""),
		gen.Import(pkg.Path, ""),
	)

	for path, name := range imports {
		mockImports = append(mockImports, gen.Import(path, name))
		implImports = append(implImports, gen.Import(path, name))
	}

	var directives []gen.WriteDirective

	mockGen := gen.Block(
		gen.Package(
			gen.Name(packageMockName),
			gen.Imports(mockImports...),
			gen.Block(
				gen.SourceText(
					string(static.MustReadFile("mock.tml", true)),
					struct {
						InterfaceName string
						Package       ast.Package
						Methods       []ast.FunctionDefinition
					}{
						Package:       pkg,
						Methods:       methods,
						InterfaceName: interfaceName,
					},
				),
			),
		),
	)

	implGen := gen.Block(
		gen.Package(
			gen.Name(packageName),
			gen.Imports(implImports...),
			gen.Block(
				gen.SourceText(
					string(static.MustReadFile("impl.tml", true)),
					struct {
						InterfaceName string
						Package       ast.Package
						Methods       []ast.FunctionDefinition
					}{
						Package:       pkg,
						Methods:       methods,
						InterfaceName: interfaceName,
					},
				),
			),
		),
	)

	directives = append(directives, gen.WriteDirective{
		Writer:   fmtwriter.New(mockGen, true, true),
		FileName: fmt.Sprintf("%s.go", packageMockName),
		Dir:      filepath.Join(packageName, packageMockName),
	})

	directives = append(directives, gen.WriteDirective{
		Writer:   fmtwriter.New(implGen, true, true),
		FileName: fmt.Sprintf("%s.go", packageName),
		Dir:      packageName,
	})

	return directives, nil
}
