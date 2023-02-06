package main

import (
	"github.com/spf13/cobra"
	"github.com/zlobste/ala/internal/report"
	"log"

	"github.com/zlobste/ala/internal/analyzer"
)

var logFilePath, geoDBPath string

var rootCmd = &cobra.Command{
	Use:   "ala",
	Short: "A tool for reporting",
	Long:  `A tool for reporting using a log file and a geo database.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		countryReport := report.NewCountryReport("Top views by countries")
		stateReport := report.NewStateReport("Top views by states of US", "US")

		analyzer := analyzer.NewAnalyzer(logFilePath, geoDBPath, []report.Report{
			countryReport,
			stateReport,
		})

		return analyzer.Run()
	},
}

func init() {
	rootCmd.Flags().StringVarP(&logFilePath, "log", "l", "", "Path to the log file")
	rootCmd.MarkFlagRequired("log")

	rootCmd.Flags().StringVarP(&geoDBPath, "db", "d", "", "Path to the geo database")
	rootCmd.MarkFlagRequired("db")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to run reporter: %v", err)
	}
}
