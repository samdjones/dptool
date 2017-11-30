package cmd

import (
	"fmt"
	"os"

	dptool "github.com/samdjones/dptool/lib"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(writeCmd)
}

var writeCmd = &cobra.Command{
	Use:   "write [local src file] [gateway dest file]",
	Short: "Write file to gateway",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dpt := dptool.NewDPTool(Verbose, User, Pass)
		if err := dpt.PutFile(args[0], args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
