package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgPath string

// RootCmd base information of goproxy
var RootCmd = cobra.Command{
	Use:   "goproxy",
	Short: "goproxy is a proxy build with golang",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().String("config", ".goproxy", "your config file path, default is $HOME/.goproxy.json)")
}
