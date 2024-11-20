package utils

import (
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// 检测文件夹或文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 检测目录是否存在，不存在就创建
func PathExistForCreate(path string) {
	exists, _ := PathExists(path)
	if !exists {
		os.MkdirAll(path, os.ModePerm)
	}
}

// 从文件读取imgage
func ReadImg(imgFile string) (image.Image, error) {
	f, err := os.Open(imgFile)
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func DeleteFile(path string) {

	// 删除文件
	err := os.Remove(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func DeleteFileS() {
	// 定义要检测的目录
	dir := "./"

	// 编译正则表达式，用于匹配文件名模式
	re := regexp.MustCompile(`^code\d+\.png$`)

	// 控制循环的变量
	keepRunning := true

	for keepRunning {
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			continue
		}

		// 记录本次扫描中是否找到了匹配的文件
		foundMatchingFile := false

		// 遍历目录中的所有文件
		for _, file := range files {
			// 检查文件是否符合模式
			if re.MatchString(file.Name()) {
				foundMatchingFile = true
				// 构建完整的文件路径
				filePath := filepath.Join(dir, file.Name())
				// 删除文件
				err := os.Remove(filePath)
				if err != nil {
					fmt.Printf("Failed to delete %s: %v\n", filePath, err)
				} else {
					fmt.Printf("\nDeleted %s\n", filePath)
				}
			}
		}

		// 如果本次扫描没有找到匹配的文件，结束循环
		if !foundMatchingFile {
			keepRunning = false
			fmt.Println("No more matching files found. Exiting...")
		} else {
			// 等待一段时间后再次检查，防止CPU占用过高
			time.Sleep(2 * time.Second)
		}
	}
}
