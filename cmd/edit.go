package cmd

import (
	"fmt"
	"strconv"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit <id> <new command>",
	Short: "Edit a saved snippet",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil || id < 1 {
			return fmt.Errorf("snippet ID must be a positive number")
		}

		tags, err := cmd.Flags().GetStringSlice("tag")
		if err != nil {
			return err
		}

		if err := snippet.Update(id, args[1], tags...); err != nil {
			return err
		}

		fmt.Printf("Updated snippet %d\n", id)
		return nil
	},
}

func init() {
	editCmd.Flags().StringSliceP("tag", "t", nil, "new tags for the snippet")
	rootCmd.AddCommand(editCmd)
}
