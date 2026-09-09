package cmd

import (
	"fmt"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new snippet",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		newSnippet, err := snippet.Add(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("Added %d: %s\n", newSnippet.ID, newSnippet.Command)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
