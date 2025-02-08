package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/appends"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/buildtag"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/ctrlflow"
	"golang.org/x/tools/go/analysis/passes/defers"
	"golang.org/x/tools/go/analysis/passes/errorsas"
	"golang.org/x/tools/go/analysis/passes/findcall"
	"golang.org/x/tools/go/analysis/passes/httpresponse"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/stdmethods"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/tests"
	"golang.org/x/tools/go/analysis/passes/unmarshal"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"golang.org/x/tools/go/analysis/passes/unusedresult"
	"golang.org/x/tools/go/analysis/passes/waitgroup"
	"honnef.co/go/tools/quickfix"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck"

	"github.com/Makovey/shortener/cmd/staticlint/noexit"
)

func main() {
	analyzers := []*analysis.Analyzer{
		appends.Analyzer,      // incorrect use of append
		assign.Analyzer,       // uninitialized variables
		bools.Analyzer,        // redundant logical expressions with bools
		buildtag.Analyzer,     // incorrect build-tags
		copylock.Analyzer,     // copy of mutex
		ctrlflow.Analyzer,     // unreachable code
		defers.Analyzer,       // incorrect use of defer
		errorsas.Analyzer,     // incorrect use of errors.As
		findcall.Analyzer,     // search for a call to a given function
		httpresponse.Analyzer, // unverified HTTP responses
		loopclosure.Analyzer,  // errors in closures inside loops
		lostcancel.Analyzer,   // lost context.CancelFunc
		nilfunc.Analyzer,      // calling nil functions
		shadow.Analyzer,       // shadowing variable
		stdmethods.Analyzer,   // compliance with standard interfaces
		structtag.Analyzer,    // incorrect structure tags
		tests.Analyzer,        // errors in tests
		unmarshal.Analyzer,    // incorrect work with encoding/json.
		unreachable.Analyzer,  // unreachable code
		unusedresult.Analyzer, // unused result of functions
		waitgroup.Analyzer,    // Incorrect use of sync.WaitGroup
	}

	for _, val := range staticcheck.Analyzers {
		analyzers = append(analyzers, val.Analyzer)
	}

	// ST1012 - poorly chosen name for error variable
	// ST1023 - redundant type in variable declaration
	for _, val := range stylecheck.Analyzers {
		if val.Analyzer.Name == "ST1012" || val.Analyzer.Name == "ST1023" {
			analyzers = append(analyzers, val.Analyzer)
		}
	}

	// QF1009 - Use time.Time.Equal instead of == operator
	// QF1012 - Use fmt.Fprintf(x, ...) instead of x.Write(fmt.Sprintf(...))
	for _, val := range quickfix.Analyzers {
		if val.Analyzer.Name == "QF1009" || val.Analyzer.Name == "QF1012" {
			analyzers = append(analyzers, val.Analyzer)
		}
	}

	analyzers = append(analyzers, noexit.NoExitAnalyzer)

	multichecker.Main(analyzers...)
}
