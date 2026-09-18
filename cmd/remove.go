package cmd

import (
	"database/sql"
	"errors"
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
			return fmt.Errorf("snippet ID must be a number: %q", args[0])
		}
		if id < 1 {
			return fmt.Errorf("snippet ID must be greater than zero")
		}
		if err := snippet.Remove(id); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return fmt.Errorf("snippet %d was not found", id)
			}
			return err
		}
		fmt.Println("Removed snippet", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
