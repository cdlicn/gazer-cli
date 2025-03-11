package cmd

import (
	"fmt"
	"gazer/common"
	"gazer/etcd"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "get the list of log files.",

	Run: func(cmd *cobra.Command, args []string) {
		// 加入etcd
		list, err := etcd.List(common.Ip)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for k, v := range list {
			fmt.Printf("%s\t%s\n", k, v)
		}

		fmt.Println("get list success")
	},
}

func init() {
	listCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `gazer list` to get the list of log files.")
		return
	})
	rootCmd.AddCommand(listCmd)
}
