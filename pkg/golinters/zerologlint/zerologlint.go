package zerologlint

import (
	"github.com/ykadowak/zerologlint"

	"github.com/edgedelta/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(zerologlint.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
