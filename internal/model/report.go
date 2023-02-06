package model

import "github.com/oschwald/geoip2-golang"

type LogData struct {
	GeoData  *geoip2.City
	PagePath string
}

type UnitReport struct {
	Name            string
	Views           int64
	MostVisitedPage string
}
