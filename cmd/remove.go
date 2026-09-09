package cmd

import (
	"fmt"
	"strconv"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <id>",
	Short: "Remove a saved snippet by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("snippet ID must be a number: %w", err)
		}
		if err := snippet.Remove(id); err != nil {
			return err
		}
		fmt.Println("Removed snippet", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
