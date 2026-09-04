package cmd

import (
	"fmt"

	"github.com/Pointdexter37/kin/internal/snippet"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List saved snippets",
	Run: func(cmd *cobra.Command, arg []string){
		snippets := snippet.List()

		for _, s := range snippets {
			fmt.Println(s.ID, s.Command)
		}
			
		
		fmt.Println("List snippets")
	},
}

func init(){
	rootCmd.AddCommand(listCmd)
}