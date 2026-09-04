package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new snippet",
    Args:  cobra.ExactArgs(1),

    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Add snippet:", args[0])
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}