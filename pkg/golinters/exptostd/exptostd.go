package exptostd

import (
	"github.com/ldez/exptostd"

	"github.com/edgedelta/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(exptostd.NewAnalyzer()).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
