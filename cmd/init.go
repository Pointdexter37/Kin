package cmd 

import (
	"fmt"

	"github.com/spf13/cobra"

	
)

var initCmd = &cobra.Command{
	Use : "init",
	Short: "Initialize kin",
	Run : func(cmd *cobra.Command, arge []string){
		fmt.Println("kin initilized")
	},
}

func init(){
	rootCmd.AddCommand(initCmd)
} // Go has a special function called: init(), its run automatically befor main()