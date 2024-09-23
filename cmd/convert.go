/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/hectoraldairah/curcuma/pkg/converter"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	value float64
	from  string
	to    string
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		rates, err := converter.FechRates(from, to)

		if err != nil {
			fmt.Printf("Error in convertert:  %v", err)
		}

		convertedVal := converter.ConvertValue(value, rates[to])

		fmt.Printf("%.2f " +to + "\n", convertedVal)

	},
}

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error to load .env")
	}

	convertCmd.Flags().Float64VarP(&value, "value", "v", 1.0, "Value to convert")
	convertCmd.Flags().StringVarP(&from, "from", "f", "USD", "Source currency")
	convertCmd.Flags().StringVarP(&to, "to", "t", "MXN", "Target currency")

	rootCmd.AddCommand(convertCmd)
}
