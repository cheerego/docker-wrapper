package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "docker-wrapper",
	Short: "a docker image pull cli tool for gcr.io quay.io",
	Long:  `a docker image pull cli tool for gcr.io quay.io`,
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println("目前版本为1.0.0")
		} else {
			cmd.HelpFunc()(cmd, args)
		}
	},
}

var versionFlag bool

func Execute() {
	rootFlags := rootCmd.PersistentFlags()
	rootFlags.BoolVarP(&versionFlag, "version", "v", false, "show version")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
