/*
Copyright © 2023 Rory McCune rorym@mccune.org.uk

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "kubelet_dumper",
	Short:   "A Utility to remotely Dump Kubelet configs",
	Long:    `The goal of this utility is to dump the kubelet config from remote node(s)`,
	Version: "0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringP("nodename", "n", "", "The name of the node to dump the kubelet config from")
}
