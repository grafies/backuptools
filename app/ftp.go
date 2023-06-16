package app

import (
	"backupTools/model"
	"errors"
	"fmt"
	"github.com/jlaffaye/ftp"
	"log"
	"os"
	"time"
)

func Ftp(config model.Config, zipFileName string) error {
	path := fmt.Sprintf("%v:%v", config.Ftp.Ip, config.Ftp.Port)
	log.Print("正在连接ftp")
	c, err := ftp.Dial(path, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		return errors.New(fmt.Sprint("连接失败,", err))
	}

	log.Print("正在登录ftp")
	if config.Ftp.Username == "" {
		config.Ftp.Username = "guest"
	}
	if err = c.Login(config.Ftp.Username, config.Ftp.Password); err != nil {
		return errors.New(fmt.Sprint("登录失败,", err))
	}

	log.Print("正在处理上传文件")
	zipFileNamePath := fmt.Sprintf("%v/%v", config.TargetDir.Path, zipFileName)
	file, err := os.Open(zipFileNamePath) //打开文件
	if err != nil {
		return errors.New(fmt.Sprint("打开文件失败,", err))
	}
	defer file.Close()

	log.Printf("正在上传文件...")
	ftpPath := fmt.Sprintf("%v/%v", config.Ftp.Path, zipFileName)
	if err = c.Stor(ftpPath, file); err != nil {
		return errors.New(fmt.Sprint("上传文件失败,", err))
	}

	return nil
}
