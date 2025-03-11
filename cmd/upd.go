package cmd

import (
	"fmt"
	"gazer/common"
	"gazer/etcd"
	"gazer/utils"

	"github.com/spf13/cobra"
)

// updCmd represents the upd command
var updCmd = &cobra.Command{
	Use:   "upd",
	Short: "update a log file.",

	Run: func(cmd *cobra.Command, args []string) {
		// 检查参数数量
		if len(args) == 0 {
			fmt.Println("Use `gazer upd [topic] [path]` to update log files to be collected.")
			return
		}
		if len(args) != 2 {
			fmt.Println("wrong number of parameters, number should be 2.")
			return
		}
		// 取出参数
		topic := args[0]
		path := args[1]
		absPath, b := utils.IsFileExist(path)
		if !b {
			fmt.Println("failed to load file.")
			return
		}
		path = absPath

		// 加入etcd
		err := etcd.Update(common.Ip, topic, path)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// 修改upd etcd
		err = etcd.Put("upd", common.Ip+"\\"+topic)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("update success")
	},
}

func init() {
	updCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `gazer upd [topic] [path]` to update log files to be collected.")
		return
	})
	rootCmd.AddCommand(updCmd)
}
