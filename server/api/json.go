package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"server/model"
)

// 读json文件
func Read(fileName string) []model.Word {
	filePtr, err := os.Open("./json/" + fileName + ".json")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
	}
	defer filePtr.Close()
	var data []model.Word
	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&data)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	} else {
		fmt.Println("解码成功")

	}
	return data
}

// 写json文件
func Write(data []model.Word, name string) {
	// 创建文件
	filePtr, err := os.Create("./json/" + name + ".json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return
	}
	defer filePtr.Close()
	// 创建json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(&data)
	if err != nil {
		fmt.Println("编码失败", err.Error())
	} else {
		fmt.Println("编码成功")
	}
}
func GetList() []string {
	files, err := os.ReadDir("./json")
	if err != nil {
		fmt.Println(err)
	}
	var arr []string
	for i := 0; i < len(files); i++ {
		arr = append(arr, files[i].Name()[:len(files[i].Name())-5])
	}
	return arr
}

// 获取每张表的长度
func GetLen(fileName string) int {
	list := Read(fileName)
	return len(list)
}

// 获取所有单词表的详细信息
func GetDetail() []model.Detail {
	dir := "./json" // 设置目录路径
	var details []model.Detail
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// 排除目录本身
		if path == dir {
			return nil
		}
		// 获取文件名
		fileName := filepath.Base(path)[:len(filepath.Base(path))-5]
		// 获取文件的创建时间
		createdTime := info.ModTime()
		// 文件名、创建时间和json数据长度
		details = append(details, model.Detail{Name: fileName, Init: createdTime.String()[:19], Len: GetLen(fileName)})
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
	return details
}
