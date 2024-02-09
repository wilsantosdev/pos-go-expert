/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"stresser/internal/services"

	"github.com/spf13/cobra"
)

var stressCmd = &cobra.Command{
	Use:   "stress",
	Short: "Stress test a URL, concurrent requests",
	Long:  ``,
	RunE:  runStressCmd,
}

var concurrency int
var requests int
var url string

func init() {
	rootCmd.AddCommand(stressCmd)
	stressCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Number of concurrent requests")
	stressCmd.Flags().IntVarP(&requests, "requests", "r", 1, "Number of requests")
	stressCmd.Flags().StringVarP(&url, "url", "u", "", "URL to stress test")
	stressCmd.MarkFlagRequired("concurrency")
	stressCmd.MarkFlagRequired("requests")
	stressCmd.MarkFlagRequired("url")
}

func runStressCmd(cmd *cobra.Command, args []string) error {
	looper := services.NewLooper(
		services.NewRequester(),
	)
	report, err := looper.Loop(concurrency, requests, url)
	if err != nil {
		return err
	}

	services.NewReporter(report).Report()

	return nil
}
