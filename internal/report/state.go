package report

import (
	"github.com/zlobste/ala/internal/model"
)

// stateReport defines a structure that implements Report interface and generates report for the subdivisions.
type stateReport struct {
	countryReport
	countryCode string
}

// NewStateReport is a constructor for stateReport.
func NewStateReport(name, countryCode string) Report {
	return &stateReport{
		countryCode: countryCode,
		countryReport: countryReport{
			name: name,
			stat: make(map[string]map[string]int64, 0),
		},
	}
}

// AnalyzeRow implements Report interface.
func (r stateReport) AnalyzeRow(logData model.LogData) {
	if logData.GeoData == nil || logData.GeoData.Country.IsoCode != r.countryCode {
		return
	}

	var state string
	if len(logData.GeoData.Subdivisions) != 0 {
		state = logData.GeoData.Subdivisions[0].Names[enLangCode]
	} else {
		state = unknownUnitName
	}

	pages, ok := r.stat[state]
	if !ok {
		r.stat[state] = make(map[string]int64, 0)
	}

	r.stat[state][logData.PagePath] = pages[logData.PagePath] + 1
}
