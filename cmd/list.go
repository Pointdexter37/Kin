package cmd

import (
	"fmt"
	

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use: "list",
	Short: "List saved snippets",
	Run: func(cmd *cobra.Command, arg []string){
		fmt.Println("List snippets")
	},
}

func init(){
	rootCmd.AddCommand(listCmd)
}