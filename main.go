package main

import (
	"fmt"
	"gazer/cmd"
	"gazer/common"
	"gazer/etcd"
)

func main() {
	err := common.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer etcd.Close()
	//fmt.Printf("%+v", config.ConfigObj)
	cmd.Execute()
}
