package utils

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// IsFileExist 检查path是否存在
func IsFileExist(path string) (string, bool) {
	// 判断是否为绝对路径
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return "", false
		}
		path = absPath
	}
	// 获取当前文件地址
	_, err := os.Stat(path)
	if err != nil {
		return "", false
	}
	return path, true
}

// GetOutboundIp 获取本机Ip
func GetOutboundIp() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return ip, fmt.Errorf("failed to get ip.")
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return ip, nil
}
