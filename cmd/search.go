package cmd

import (
	"fmt"
	"strconv"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var interactiveSearch bool

var searchCmd = &cobra.Command{
	Use:   "search <text>",
	Short: "Find snippets containing text",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		snippets, err := snippet.Search(args[0])
		if err != nil {
			return err
		}
		if interactiveSearch {
			return chooseSnippet(cmd, snippets)
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
	searchCmd.Flags().BoolVarP(&interactiveSearch, "interactive", "i", false,
		"choose one result interactively")
	rootCmd.AddCommand(searchCmd)
}

func chooseSnippet(cmd *cobra.Command, snippets []snippet.Snippet) error {
	if len(snippets) == 0 {
		fmt.Fprintln(cmd.OutOrStdout(), "No snippets found.")
		return nil
	}
	for index, item := range snippets {
		fmt.Fprintf(cmd.OutOrStdout(), "%d) %s\n", index+1, item.Command)
	}
	fmt.Fprint(cmd.OutOrStdout(), "Choose a snippet number: ")

	var input string
	if _, err := fmt.Fscanln(cmd.InOrStdin(), &input); err != nil {
		return fmt.Errorf("reading selection: %w", err)
	}
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(snippets) {
		return fmt.Errorf("selection must be a number from 1 to %d", len(snippets))
	}
	fmt.Fprintln(cmd.OutOrStdout(), snippets[choice-1].Command)
	return nil
}
