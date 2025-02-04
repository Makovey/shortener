package noexit

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// NoExitAnalyzer анализатор, проверяющий наличие os.Exit в main
var NoExitAnalyzer = &analysis.Analyzer{
	Name: "noexit",
	Doc:  "сhecks usage os.Exit in the main function",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	if pass.Pkg.Name() != "main" {
		return nil, nil
	}

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if fn, ok := n.(*ast.FuncDecl); ok && fn.Name.Name == "main" {
				checkForOsExit(fn.Body, pass)
			}
			return true
		})
	}
	return nil, nil
}

func checkForOsExit(body *ast.BlockStmt, pass *analysis.Pass) {
	ast.Inspect(body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		pkgIdent, ok := sel.X.(*ast.Ident)
		if ok && pkgIdent.Name == "os" && sel.Sel.Name == "Exit" {
			pass.Reportf(call.Pos(), "it is forbidden to use os.Exit in main package")
		}

		return true
	})
}
