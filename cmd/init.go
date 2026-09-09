package cmd

import (
	"fmt"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize kin",
	RunE: func(cmd *cobra.Command, arge []string) error {
		if err := snippet.Init(); err != nil {
			return err
		}
		fmt.Println("kin initialized")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
} // Go has a special function called: init(), its run automatically befor main()
