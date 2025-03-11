package test

import (
	"log"
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {
	path := "D:\\project\\gazer\\gazer\\test\\mysql.log"
	// 判断是否为绝对路径
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			log.Fatalf("Error converting path to absolute: %v", err)
		}
		log.Printf("Absolute path is: %s", absPath)
	} else {
		log.Println("Given path is already absolute.")
	}
}
