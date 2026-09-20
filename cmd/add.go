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
		tags, err := cmd.Flags().GetStringSlice("tag")
		if err != nil {
			return err
		}
		newSnippet, err := snippet.Add(args[0], tags...)
		if err != nil {
			return err
		}

		fmt.Printf("Added %d: %s\n", newSnippet.ID, newSnippet.Command)
		return nil
	},
}

func init() {
	addCmd.Flags().StringSliceP("tag", "t", nil, "tags for the snippet")
	rootCmd.AddCommand(addCmd)
}
