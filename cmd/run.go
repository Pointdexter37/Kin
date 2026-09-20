package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <id>",
	Short: "Preview a saved snippet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			return fmt.Errorf("snippet ID must be a positive number")
		}

		item, err := snippet.Get(id)
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("snippet %d was not found", id)
		}
		if err != nil {
			return err
		}

		// Preview-only is intentional: no shell or operating-system process runs.
		fmt.Fprintf(cmd.OutOrStdout(), "Preview (not executed): %s\n", item.Command)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
