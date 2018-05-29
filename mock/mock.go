package mock

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gokit/mockkit/static"
	"github.com/influx6/moz/ast"
	"github.com/influx6/moz/gen"
)

// ImplOnlyGen generates a implementation source file for a giving interface type.
func ImplOnlyGen(toPkg string, an ast.AnnotationDeclaration, itr ast.InterfaceDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
	dirPath := pkgDeclr.Dir
	var templateName string
	if strings.HasSuffix(dirPath, toPkg) {
		templateName = "impl-only.tml"
	} else {
		templateName = "impl.tml"
	}

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

	for path, name := range imports {
		implImports = append(implImports, gen.Import(path, name))
	}

	implGen := gen.Block(
		gen.Package(
			gen.Name(itr.Package),
			gen.Imports(implImports...),
			gen.Block(
				gen.SourceText(
					"mockonly",
					string(static.MustReadFile(templateName, true)),
					struct {
						InterfaceName string
						Package       ast.Package
						Methods       []ast.FunctionDefinition
						Itr           ast.InterfaceDeclaration
						Pkg           ast.PackageDeclaration
					}{
						Itr:           itr,
						Package:       pkg,
						Pkg:           pkgDeclr,
						Methods:       methods,
						InterfaceName: itr.Object.Name.Name,
					},
				),
			),
		),
	)

	return []gen.WriteDirective{
		{
			Writer:   implGen,
			FileName: fmt.Sprintf("%s.mockkit.go", strings.ToLower(itr.Object.Name.Name)),
		},
	}, nil
}

// ImplGen generates a mock package with implementation and mock types for giving interface type.
func ImplGen(toPkg string, an ast.AnnotationDeclaration, itr ast.InterfaceDeclaration, pkgDeclr ast.PackageDeclaration, pkg ast.Package) ([]gen.WriteDirective, error) {
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
					"mock",
					string(static.MustReadFile("mock.tml", true)),
					struct {
						InterfaceName string
						Package       ast.Package
						Methods       []ast.FunctionDefinition
						Itr           ast.InterfaceDeclaration
						Pkg           ast.PackageDeclaration
					}{
						Itr:           itr,
						Pkg:           pkgDeclr,
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
					"mock-impl",
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
		Writer:   mockGen,
		FileName: fmt.Sprintf("%s.go", packageMockName),
		Dir:      filepath.Join(packageName, packageMockName),
	})

	directives = append(directives, gen.WriteDirective{
		Writer:   implGen,
		FileName: fmt.Sprintf("%s.go", packageName),
		Dir:      packageName,
	})

	return directives, nil
}
