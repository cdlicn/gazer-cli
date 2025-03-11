package cmd

import (
	"fmt"
	"gazer/common"
	"gazer/etcd"
	"gazer/utils"
	"github.com/spf13/cobra"
)

type Config struct {
	EtcdConfig `ini:"etcd"`
}

type EtcdConfig struct {
	Address    string `ini:"address"`
	CollectKey string `ini:"collect_key"`
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new log file.",

	Run: func(cmd *cobra.Command, args []string) {
		// 检查参数
		if len(args) == 0 {
			fmt.Println("Use `gazer add [topic] [path]` to add log files to be collected.")
			return
		}
		if len(args) != 2 {
			fmt.Println("wrong number of parameters, number should be 2.")
			return
		}

		// 取出参数
		topic := args[0]
		path := args[1]

		// 检查文件是否存在，获取绝对路径
		absPath, b := utils.IsFileExist(path)
		if !b {
			fmt.Println("failed to load file.")
			return
		}
		path = absPath

		// 加入etcd
		err := etcd.Add(common.Ip, topic, path)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// 修改add etcd
		err = etcd.Put("add", common.Ip+"\\"+topic)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("add success")
	},
}

func init() {
	addCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Println("Use `gazer add [topic] [path]` to add log files to be collected.")
		return
	})
	rootCmd.AddCommand(addCmd)
}
