package app

import (
	"backupTools/model"
	"backupTools/utils"
	"flag"
	"fmt"
	"log"
)

func Terminal() {
	var conf string
	flag.StringVar(&conf, "f", "./config/config.yaml", "配置文件路径")
	flag.Parse() //转换，调用该方法

	config := model.ServerConfig(conf)

	//生成目标路径
	zipFileName := fmt.Sprintf("%v%v%v", config.BackName, Times(), ".zip")
	zipFileNamePath := fmt.Sprintf("%v/%v",
		config.TargetDir.Path, zipFileName)

	log.Printf("正在压缩%v文件到%v", config.SourceDir, zipFileNamePath)
	if err := utils.ZipDir(config.SourceDir.Path, zipFileNamePath); err != nil {
		log.Print("压缩文件失败", err)
		return
	}
	log.Print("压缩成功！！！")

	if config.Ftp.FtpSwitch {
		log.Printf("正在上传%v文件到ftp", zipFileNamePath)
		if err := Ftp(*config, zipFileName); err != nil {
			log.Print(err)
		}
		log.Print("ftp上传成功！！！")
	}

	log.Print("正在清理历史文件")
	if err := KeepFiles(config.TargetDir.Path, config.BackName, ".zip", config.KeepFiles); err != nil {
		log.Print("清理文件失败", err)
		return
	}
	log.Print("清理完成，备份结束！！！")

}
