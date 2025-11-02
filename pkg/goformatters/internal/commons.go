package internal

import "github.com/edgedelta/golangci-lint/v2/pkg/logutils"

// FormatterLogger must be used only when the context logger is not available.
var FormatterLogger = logutils.NewStderrLog(logutils.DebugKeyFormatter)
