package testableexamples

import (
	"github.com/maratori/testableexamples/pkg/testableexamples"

	"github.com/edgedelta/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(testableexamples.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
