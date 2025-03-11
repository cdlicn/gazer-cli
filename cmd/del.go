package cmd

import (
	"fmt"
	"gazer/common"
	"gazer/etcd"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete a log file.",

	Run: func(cmd *cobra.Command, args []string) {
		// 检查参数数量
		if len(args) == 0 {
			fmt.Println("Use `gazer upd [topic]` to delete a log files.")
			return
		}
		if len(args) != 1 {
			fmt.Println("wrong number of parameters, number should be 1.")
			return
		}

		// 取出参数
		topic := args[0]

		// 删除etcd
		err := etcd.Delete(common.Ip, topic)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// del etcd
		err = etcd.Put("del", common.Ip+"\\"+topic)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("delete success")
	},
}

func init() {
	delCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `gazer upd [topic]` to delete a log files.")
		return
	})
	rootCmd.AddCommand(delCmd)
}
