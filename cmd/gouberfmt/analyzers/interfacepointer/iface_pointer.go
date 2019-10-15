package interfacepointer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "interfacepointer",
	Doc:  "reports uses of pointer of interface type",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			decl, ok := node.(*ast.GenDecl)
			if !ok {
				return true
			}

			if decl.Tok == token.VAR {
				for _, spec := range decl.Specs {
					varSpec, ok := spec.(*ast.ValueSpec)
					if !ok {
						return true
					}

					pointer, ok := varSpec.Type.(*ast.StarExpr)
					if !ok {
						return true
					}

					_, ok = pointer.X.(*ast.InterfaceType)
					if !ok {
						return true
					}

					for _, varName := range varSpec.Names {
						pass.Reportf(varName.Pos(), "interface-pointer: var %s uses pointer to interface", varName.Name)
					}

				}
			}
			return true
		})
	}
	return nil, nil
}
