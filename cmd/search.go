package cmd

import (
	"fmt"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search <text>",
	Short: "Find snippets containing text",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		snippets, err := snippet.Search(args[0])
		if err != nil {
			return err
		}
		for _, item := range snippets {
			fmt.Println(item.ID, item.Command)
		}
		if len(snippets) == 0 {
			fmt.Println("No snippets found.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
