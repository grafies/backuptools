package utils

import (
	"testing"
	"time"
)

func TestUtils(t *testing.T) {
	//t.Run("测试查询用户", testGetUser)
	t.Run("测试zip", testZip)

}

func testZip(t *testing.T) {
	progress := &Progress{}
	go ProgressBar(progress)
	for {
		for i := 0; i < 100; i++ {
			progress.value++
			time.Sleep(100 * time.Millisecond)
		}
		if progress.value == 100 {
			break
		}
	}
}
