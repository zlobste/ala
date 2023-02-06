package analyzer

import (
	"bufio"
	"github.com/oschwald/geoip2-golang"
	"github.com/pkg/errors"
	"github.com/zlobste/ala/internal/model"
	"github.com/zlobste/ala/internal/report"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

// Defines a regular expression to match the log lines.
var logLineRegex = regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}).*?".*?(\S+)\s+HTTP`)

// Defines a slice of strings that represent excluded patterns.
var excludePatterns = []string{
	"/entry-images/",
	"/images/",
	"/user-images/",
	"/static/",
	"/css/",
	"/js/",
	".txt",
	".ico",
	".rss",
	".atom",
	".png",
}

// Analyzer defines the interface for log analyzer.
type Analyzer interface {
	Run() error
}

// analyzer defines a structure of the Apache Log Analyzer that parses log files and generates custom reports.
type analyzer struct {
	logFilePath string
	geoDBPath   string
	reports     []report.Report
}

// NewAnalyzer creates and returns a new instance of Analyzer.
func NewAnalyzer(logFilePath, geoDBPath string, reports []report.Report) Analyzer {
	return &analyzer{
		logFilePath: logFilePath,
		geoDBPath:   geoDBPath,
		reports:     reports,
	}
}

// Run reads the log file and enriches the data with data from the database and generates custom reports.
func (a analyzer) Run() error {
	// Load the GeoLite2 City database.
	db, err := geoip2.Open(a.geoDBPath)
	if err != nil {
		return errors.Wrap(err, "error opening GeoLite2 City database")
	}
	defer db.Close()

	file, err := os.Open(a.logFilePath)
	if err != nil {
		return errors.Wrap(err, "error opening log file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Read the log file and send the lines to the lines channel.
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return errors.Wrap(err, "error reading log file")
		}

		line := scanner.Text()

		// Check if the line matches the log line format and does not contain any of the excluded patterns.
		if logLineRegex.MatchString(line) && !containsExcludedPattern(line, excludePatterns) {
			// Extract the IP address and page path from the log line.
			matches := logLineRegex.FindStringSubmatch(line)
			ip := matches[1]
			pagePath := matches[2]

			// Lookup the geographical information for the IP address.
			record, err := db.City(net.ParseIP(ip))
			if err != nil {
				log.Println("Error looking up IP address:", err)

				continue
			}

			logData := model.LogData{
				PagePath: pagePath,
				GeoData:  record,
			}

			for _, val := range a.reports {
				val.AnalyzeRow(logData)
			}
		}
	}

	for _, val := range a.reports {
		val.PrintReport()
	}

	return nil
}

// containsExcludedPattern checks if a string contains any of the excluded patterns.
func containsExcludedPattern(s string, exclude []string) bool {
	for _, pattern := range exclude {
		if strings.Contains(s, pattern) {
			return true
		}
	}

	return false
}
