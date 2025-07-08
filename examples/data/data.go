package data

import (
	"os"
	"path/filepath"
)

// ReadFromFile 读取文件内容
func ReadFromFile(filename string) ([]byte, error) {
	if filename == "" {
		return nil, os.ErrInvalid
	}
	//	判断文件是否存在
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []byte

	_, err = file.Read(data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// WriteToFile 将数据写入文件（自动添加换行）
// filename: 目标文件路径
// data: 要写入的内容
func WriteToFile(filename string, data []string) error {

	// 文件所在目录不存在时，自动创建)
	targetDir := filepath.Dir(filename)
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return err
	}

	// 以追加模式打开文件（不存在时自动创建）
	file, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, // 组合模式
		0644,                                // Unix 权限位
	)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range data {
		content := []byte(line)

		// 自动添加换行符（不同系统保持兼容）
		content = append(content, '\n')

		// 执行写入操作
		if _, err := file.Write(content); err != nil {
			return err
		}
	}

	// 确保数据落盘
	return file.Sync()
}
