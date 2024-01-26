/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stresser",
	Short: "A brief description of your application",
	Long: `Aplicações de teste de stress que permitem simular o comportamento de um sistema em condições de carga.
	Como resultado, é possível avaliar o comportamento do sistema e identificar possíveis problemas de performance.
	
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("root called")
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "url to stress test")
	rootCmd.Flags().IntP("concurrency", "c", 1, "concurrency")
	rootCmd.Flags().IntP("requests", "r", 1, "requests")
	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("concurrency")
	rootCmd.MarkFlagRequired("requests")

}
