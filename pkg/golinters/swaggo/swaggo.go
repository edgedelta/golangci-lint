package swaggo

import (
	"github.com/edgedelta/golangci-lint/v2/pkg/goanalysis"
	"github.com/edgedelta/golangci-lint/v2/pkg/goformatters"
	"github.com/edgedelta/golangci-lint/v2/pkg/goformatters/swaggo"
	"github.com/edgedelta/golangci-lint/v2/pkg/golinters/internal"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(
			goformatters.NewAnalyzer(
				internal.LinterLogger.Child(swaggo.Name),
				"Check if swaggo comments are formatted",
				swaggo.New(),
			),
		).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
