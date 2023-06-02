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
	zipFileName := fmt.Sprintf("%v/%v%v%v",
		config.TargetDir.Path, config.BackName, Times(), ".zip")

	log.Print("正在压缩文件到", zipFileName)
	if err := utils.ZipDir(config.SourceDir.Path, zipFileName); err != nil {
		log.Print("压缩文件失败", err)
		return
	}

	log.Print("正在清理历史文件")
	if err := KeepFiles(config.TargetDir.Path, config.BackName, ".zip", config.KeepFiles); err != nil {
		log.Print("清理文件失败", err)
		return
	}

}
