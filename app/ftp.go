package app

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"os"
	"time"
)

func Ftp() {
	//连接ftp
	c, err := ftp.Dial("192.168.10.13:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Panic("连接失败,", err)
	}

	//登录ftp
	if err = c.Login("guest", ""); err != nil {
		log.Panic("登录失败,", err)
	}

	//打开上传的文件
	file, err := os.Open("E:/backuptools/笔记20230602134852.zip")
	if err != nil {
		log.Panic("打开文件失败,", err)
	}
	defer file.Close()

	output, _ := simplifiedchinese.GB18030.NewDecoder().String("笔记20230602134852.zip")
	fmt.Println(output)
	if err = c.Stor(output, file); err != nil {
		log.Panic("上传文件失败,", err)
	}

}
