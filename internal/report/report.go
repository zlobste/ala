package report

import (
	"github.com/zlobste/ala/internal/model"
)

// Report defines the interface for custom reports based on logs data.
type Report interface {
	// AnalyzeRow is used to analyze a row.
	AnalyzeRow(logData model.LogData)
	// PrintReport is used to print the report after full log analysis.
	PrintReport()
}
