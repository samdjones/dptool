package cmd

import (
	"fmt"
	"os"

	dptool "github.com/samdjones/dptool/lib"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync [local src dir] [gateway dest dir]",
	Short: "Continuously syncs files from local dir to gateway dir (non-recursive, ignoring .* files)",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dpt := dptool.NewDPTool(Verbose, User, Pass)
		if err := dpt.SyncDir(args[0], args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
