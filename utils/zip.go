package utils

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ZipDir(SrcDir string, ZipFileName string) error {
	fz, err := os.Create(ZipFileName) //创建zip文件
	if err != nil {
		log.Println("创建zip文件失败,", err)
		return err
	}
	defer fz.Close() //延时关闭

	w := zip.NewWriter(fz) //把文件fz开启一个写入方法，准备向fz写入数据
	defer w.Close()        //延时关闭

	/*
		1.filepath包实现了兼容各操作系统的文件路径的实用操作函数。
		2.Walk函数会遍历SrcDir指定的目录下的文件树
		3.SrcDir为你指定的路径
		4.引用WalkFunc函数func(path string, info os.FileInfo, err error) error
			path：遍历文件的路径赋值给path
			info：FileInfo用来描述一个文件对象，为文件时返回true，不为文件时返回flash
	*/

	//遍历SrcDir下所有文件
	filepath.Walk(SrcDir, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() { //判断是否为文件夹，当不为文件夹时，创建文件
			fDest, err := w.Create(path[len(SrcDir)+1:]) //根据路径创建文件
			if err != nil {
				log.Println("创建失败，", err)
				return nil
			}

			fSrc, err := os.Open(path) //打开path文件
			if err != nil {
				log.Println("文件打开失败，", err)
				return nil
			}
			defer fSrc.Close() //延时关闭

			_, err = io.Copy(fDest, fSrc) //把源文件fSrc拷贝写入到压缩包fDest
			if err != nil {
				log.Println("文件复制失败，", err)
				return nil
			}

		}
		return nil
	})
	return nil
}
