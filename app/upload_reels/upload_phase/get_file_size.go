package uploadphase

import (
	"fmt"
	"os"
)

func GetFileSize(file_name string) int {
	openFile, _ := os.Stat(file_name)
	fmt.Println(file_name)
	fileSize := openFile.Size()
	return int(fileSize)
}
