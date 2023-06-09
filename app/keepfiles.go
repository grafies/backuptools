package app

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Times 获取当前时间
func Times() string {
	times := fmt.Sprintf("%02d%02d%02d%02d%02d%02d",
		time.Now().Year(), int(time.Now().Month()), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())
	return times
}

var delFileName []string //待删除的文件名称

// KeepFiles 保留历史备份文件
func KeepFiles(distDirRoute string, backName string, compressedFormat string, KeepFileVal int) error {

	f, err := os.Open(distDirRoute)
	if err != nil {
		log.Println(err)
		return err
	}

	/**
	如果n>0，Readdir函数会返回一个最多n个成员的切片。这时，
	如果Readdir返回一个空切片，它会返回一个非nil的错误说明
	原因。如果到达了目录f的结尾，返回值err会是io.EOF。

	如果n<=0，Readdir函数返回目录中剩余所有文件对象的FileInfo
	构成的切片。此时，如果Readdir调用成功（读取所有内容直到结尾）
	，它会返回该切片和nil的错误值。如果在到达结尾前遇到错误，会
	返回之前成功读取的FileInfo构成的切片和该错误。**/
	Files, err := f.Readdir(-1)
	_ = f.Close() //因为已经赋值，所以直接关闭
	if err != nil {
		log.Println(err)
		return err
	}

	//遍历所有名称，并过滤长度不对的
	for _, dir := range Files {
		// i1 := 0                                   //用来判断是否为系统生成的文件
		str := strings.Contains(dir.Name(), backName) //查找文件名称
		if str {
			if (len(dir.Name()) - len(backName)) == (len(Times()) + len(compressedFormat)) { //过滤总长度不对的
				condition := notNumeric(dir.Name(), backName, compressedFormat)
				if condition {
					delFileName = append(delFileName, dir.Name())
				}

			}
		}
	}

	//获取删除文件个数
	if KeepFileVal > 0 { //判断KeepFileVal是否为空，不为空则执行
		delFile := len(delFileName) - KeepFileVal
		for i, v := range delFileName {
			if i < delFile {
				log.Printf("正在清理%v/%v", distDirRoute, v)
				err = os.Remove(fmt.Sprintf("%v/%v", distDirRoute, v))
				if err != nil {
					log.Println(err)
					return err
				}
			}
		}
	}

	return nil
}

// 传入文件名称，备份名称，文件后缀,处理所有正确的时间格式
func notNumeric(filename string, backName string, compressedFormat string) bool {
	// len(filename) 文件总长度 = len(backName) 备份文件名长度 + len(tools.Times()) 时间长度 + len(compressedFormat) 后缀长度
	var a bool
	for i, v := range filename {
		// fmt.Printf("下标=%v,值=%v\n", i, v)
		if i > len(backName) && i < len(backName)+len(Times()) {
			if v >= 48 && v <= 57 {
				a = true
			} else {
				return false
			}
		}
	}
	return a
}
