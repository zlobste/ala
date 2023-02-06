package report

import (
	"fmt"
	"github.com/zlobste/ala/internal/model"
	"sort"
)

const (
	unknownUnitName = "Unknown"
	enLangCode      = "en"
)

// countryReport defines a structure that implements Report interface and generates report for the countries.
type countryReport struct {
	name string
	stat map[string]map[string]int64
}

// NewCountryReport is a constructor for countryReport.
func NewCountryReport(name string) Report {
	return &countryReport{
		name: name,
		stat: make(map[string]map[string]int64, 0),
	}
}

// AnalyzeRow implements Report interface.
func (r countryReport) AnalyzeRow(logData model.LogData) {
	var country string
	if logData.GeoData != nil {
		country = logData.GeoData.Country.Names[enLangCode]
	} else {
		country = unknownUnitName
	}

	pages, ok := r.stat[country]
	if !ok {
		r.stat[country] = make(map[string]int64, 0)
	}

	r.stat[country][logData.PagePath] = pages[logData.PagePath] + 1
}

// PrintReport implements Report interface.
func (r countryReport) PrintReport() {
	fmt.Println(r.name)

	reportData := r.createReport(r.stat)
	topUnits := r.getTopUnits(reportData, 10)

	for i, val := range topUnits {
		fmt.Println(fmt.Sprintf("%d. %s: %d views, page: %s", i+1, val.Name, val.Views, val.MostVisitedPage))
	}
}

// createReport aggregates logs data.
func (r countryReport) createReport(data map[string]map[string]int64) []model.UnitReport {
	reportData := make([]model.UnitReport, 0, len(data))

	for unit, pages := range data {
		var totalUnitViews int64
		var mostVisitedPage string
		var mostVisitedPageViews int64

		for page, views := range pages {
			totalUnitViews += views

			if views > mostVisitedPageViews && page != "/" {
				mostVisitedPage = page
				mostVisitedPageViews = views
			}
		}

		reportData = append(reportData, model.UnitReport{
			Name:            unit,
			Views:           totalUnitViews,
			MostVisitedPage: mostVisitedPage,
		})
	}

	return reportData
}

// getTopUnits returns top n units.
func (r countryReport) getTopUnits(data []model.UnitReport, count int) []model.UnitReport {
	sort.Slice(data, func(i, j int) bool {
		return data[i].Views > data[j].Views
	})

	topViews := data
	if len(topViews) > count {
		topViews = topViews[:count]
	}

	return topViews
}
