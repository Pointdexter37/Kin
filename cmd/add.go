package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
    "github.com/Pointdexter37/kin/internal/snippet"
)

var addCmd = &cobra.Command{
    Use:   "add",
    Short: "Add a new snippet",
    Args:  cobra.ExactArgs(1),

    Run: func(cmd *cobra.Command, args []string) {
    	newSnippet := snippet.Add(args[0])

    
        fmt.Println("Added:", newSnippet.Command)
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}