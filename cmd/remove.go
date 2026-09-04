package cmd

import(
	"fmt"

	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use: "remove",
	Short: "Reomves something",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("removed what you inteded sir")
	},
}

func init(){
	rootCmd.AddCommand(RemoveCmd)
}