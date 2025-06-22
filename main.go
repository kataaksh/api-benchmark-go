package main

import (
	"api-benchmark/benchmark"
	"api-benchmark/portfolio"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	var url string
	var requests int
	var concurrency int

	// Root command: benchmark tool (default)
	var rootCmd = &cobra.Command{
		Use:   "apibench",
		Short: "API Benchmarking Tool",
		Run: func(cmd *cobra.Command, args []string) {
			err := benchmark.Run(url, requests, concurrency)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.Flags().StringVarP(&url, "url", "u", "", "Target URL")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 100, "Total number of requests")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Number of concurrent workers")
	rootCmd.MarkFlagRequired("url")

	// Portfolio subcommand
	var portfolioCmd = &cobra.Command{
		Use:   "portfolio",
		Short: "Show interactive portfolio CLI",
		Run: func(cmd *cobra.Command, args []string) {
			p := portfolio.NewPortfolio()
			p.Show()
		},
	}

	// Add portfolio subcommand to root
	rootCmd.AddCommand(portfolioCmd)

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
