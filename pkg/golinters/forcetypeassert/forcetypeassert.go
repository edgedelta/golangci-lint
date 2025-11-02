package forcetypeassert

import (
	"github.com/gostaticanalysis/forcetypeassert"

	"github.com/edgedelta/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(forcetypeassert.Analyzer).
		WithDesc("Find forced type assertions").
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
