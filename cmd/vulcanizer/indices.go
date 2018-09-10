package main

import (
	"fmt"

	v "github.com/github/vulcanizer"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdIndices)
}

var cmdIndices = &cobra.Command{
	Use:   "indices",
	Short: "Display the indices of the cluster.",
	Long:  `Show what indices are created on the give cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, port := getConfiguration()
		rows, header := v.GetIndices(host, port)
		table := renderTable(rows, header)
		fmt.Println(table)
	},
}