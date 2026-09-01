package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Add a new snippet",

	Run: func(cmd *cobra.Command, agr []string){
		fmt.Println("Add snippet")
	},
}

func init(){
	rootCmd.AddCommand(addCmd)
}