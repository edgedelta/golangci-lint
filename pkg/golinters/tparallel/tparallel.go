package tparallel

import (
	"github.com/moricho/tparallel"

	"github.com/edgedelta/golangci-lint/v2/pkg/goanalysis"
)

func New() *goanalysis.Linter {
	return goanalysis.
		NewLinterFromAnalyzer(tparallel.Analyzer).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
