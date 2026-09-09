package cmd

import (
	"fmt"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved snippets",
	RunE: func(cmd *cobra.Command, arg []string) error {
		snippets, err := snippet.List()
		if err != nil {
			return err
		}

		for _, s := range snippets {
			fmt.Println(s.ID, s.Command)
		}
		if len(snippets) == 0 {
			fmt.Println("No snippets found.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
