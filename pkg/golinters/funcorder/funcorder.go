package funcorder

import (
	"github.com/manuelarte/funcorder/analyzer"

	"github.com/edgedelta/golangci-lint/v2/pkg/config"
	"github.com/edgedelta/golangci-lint/v2/pkg/goanalysis"
)

func New(settings *config.FuncOrderSettings) *goanalysis.Linter {
	var cfg map[string]any

	if settings != nil {
		cfg = map[string]any{
			analyzer.ConstructorCheckName:  settings.Constructor,
			analyzer.StructMethodCheckName: settings.StructMethod,
			analyzer.AlphabeticalCheckName: settings.Alphabetical,
		}
	}

	return goanalysis.
		NewLinterFromAnalyzer(analyzer.NewAnalyzer()).
		WithConfig(cfg).
		WithLoadMode(goanalysis.LoadModeSyntax)
}
