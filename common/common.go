package common

import (
	"fmt"
	"gazer/etcd"
	"gazer/utils"
	"gopkg.in/ini.v1"
)

type CollectEntry struct {
	Topic string `json:"topic"`
	Path  string `json:"path"`
}

var (
	ConfigObj Config
	Ip        string
)

type Config struct {
	EtcdConfig `ini:"etcd"`
}

type EtcdConfig struct {
	Address string `ini:"address"`
}

func Init() error {
	// Load config
	err := ini.MapTo(&ConfigObj, "conf/config.ini")
	if err != nil {
		return fmt.Errorf("failed to load config: %v", err)
	}

	// Get ip
	Ip, err = utils.GetOutboundIp()
	if err != nil {
		return fmt.Errorf("failed to get ip: %v", err)
	}

	// Init etcd
	err = etcd.Init(ConfigObj.Address)
	if err != nil {
		return fmt.Errorf("failed to configure etcd client: %v", err)
	}

	return nil
}
