/*
Copyright © 2023 Rory McCune rorym@mccune.org.uk

*/
package cmd

import (
	kubelet_dumper "github.com/raesene/kubelet_dumper/pkg/kublet_dumper"
	"github.com/spf13/cobra"
)

// dumpKubeletCmd represents the dumpKubelet command
var dumpKubeletCmd = &cobra.Command{
	Use:   "dumpKubelet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("dumpKubelet called")
		//Dump the config from a specific nodename
		kubelet_dumper.Dump(cmd.Flag("nodename").Value.String())
	},
}

func init() {
	rootCmd.AddCommand(dumpKubeletCmd)
	dumpKubeletCmd.Flags().StringP("nodename", "n", "", "The name of the node to dump the kubelet config from")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpKubeletCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpKubeletCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
