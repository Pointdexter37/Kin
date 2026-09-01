package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use : "kin",
	Short: "A simple command snippet manager",
	//This is the description that appears in: kin --help
}

func Execute() error{
	return  rootCmd.Execute()
}