package utils

import (
	"fmt"
	"log"
	"os"
)

func Log() {
	filepath := "./log/backupTools.log"
	logPath := "./log"
	_, err := os.Stat(logPath) //查找log目录是否存在
	if err != nil {
		err = os.MkdirAll(logPath, os.ModePerm) //不存在则创建目录
		if err != nil {
			fmt.Println(err)
		}

		_, err = os.Stat(filepath)
		if err != nil {
			_, err = os.Create(filepath)
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	logFile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[backupTools]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}
