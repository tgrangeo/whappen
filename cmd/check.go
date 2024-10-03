/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/tgrangeo/whappen/rss"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "check for new articles",
	Long: `check for new articles`,
	Run: func(cmd *cobra.Command, args []string) {
		rss.FetchRSS()
		fmt.Println("check called")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
