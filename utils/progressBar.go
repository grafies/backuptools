package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type Bar struct {
	percent int64  //百分比
	cur     int64  //当前进度位置
	total   int64  //总进度条
	rate    string //进度条
	graph   string //显示符号
}
type Progress struct {
	value int64 //进度值
	to    int
}

func ProgressBar(progress *Progress) {
	var bar Bar
	bar.NewOption(0, 100)
	for {
		if progress.value > 100 { //达到100后退出
			break
		}
		time.Sleep(40 * time.Millisecond)
		bar.Play(progress.value)
	}
	bar.Finish() //打印一个换行符号
}

func (bar *Bar) NewOption(start int64, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" { //显示符号为空是添加显示符号
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph //初始化进度条位置
	}
}

func (bar *Bar) getPercent() int64 { //当前进度条除以总进度条乘100
	return int64(float32(bar.cur) / float32(bar.total) * 100)
}

func (bar *Bar) Play(cur int64) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%%  %8d%%/%d%%", bar.rate, bar.percent, bar.cur, bar.total)
}

func (bar *Bar) Finish() {
	fmt.Println()
}

func DirSize(path string) (int64, error) {
	var size int64
	if !PathExists(path) {
		return size, errors.New(fmt.Sprintf("The target file does not exist"))
	}

	err := filepath.Walk(path, func(_ string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// PathExists 判断路径是否正确
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil { //文件或目录存在
		return true
	}

	//返回一个布尔值说明该错误是否表示一个文件或目录不存在。
	//ErrNotExist和一些系统调用错误会使它返回真。
	if os.IsNotExist(err) {
		return false
	}
	return false
}
