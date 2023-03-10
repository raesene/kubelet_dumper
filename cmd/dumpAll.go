/*
Copyright © 2023 Rory McCune rorym@mccune.org.uk

*/
package cmd

import (
	kubelet_dumper "github.com/raesene/kubelet_dumper/pkg/kublet_dumper"
	"github.com/spf13/cobra"
)

// dumpAllCmd represents the dumpAll command
var dumpAllCmd = &cobra.Command{
	Use:   "dumpAll",
	Short: "Dump the kubelet config from all nodes",
	Long:  `Dumps the kubelet config from all nodes in the cluster`,
	Run: func(cmd *cobra.Command, args []string) {
		kubelet_dumper.DumpAll()
	},
}

func init() {
	rootCmd.AddCommand(dumpAllCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dumpAllCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dumpAllCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
