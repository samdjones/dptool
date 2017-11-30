package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd is the root spf13/cobra command
var RootCmd = &cobra.Command{
	Use:   "dptool [command] [command arguments] [flags]",
	Short: "dptool is a tool for working with IBM DataPower gateway appliances",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("For help, try: dptool help")
	},
}

// Verbose is the verbose output flag
var Verbose bool

// User is the gateway username flag
var User string

// Pass is the gateway password flag
var Pass string

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output (particularly when errors occur)")
	RootCmd.PersistentFlags().StringVarP(&User, "user", "u", "", "gateway username")
	RootCmd.PersistentFlags().StringVarP(&Pass, "pass", "p", "", "gateway password")
}

func initConfig() {}
