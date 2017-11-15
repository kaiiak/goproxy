package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd base information of goproxy
var RootCmd = cobra.Command{
	Use:   "goproxy",
	Short: "goproxy is a proxy build with golang",
}

// Execute run goproxy with cobra excute
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
