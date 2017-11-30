package cmd

import (
	"fmt"
	"os"

	dptool "github.com/samdjones/dptool/lib"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "delete [gateway file]",
	Short: "Delete file on the gateway",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dpt := dptool.NewDPTool(Verbose, User, Pass)
		if err := dpt.DeleteFile(args[0]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
