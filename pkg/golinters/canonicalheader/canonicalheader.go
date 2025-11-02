package canonicalheader

import (
	"github.com/lasiar/canonicalheader"

	"github.com/edgedelta/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(canonicalheader.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
