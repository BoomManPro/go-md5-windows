package main

import (
	"crypto/md5"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"os"
)

func main() {
	//编译命令
	//go build -ldflags="-H windowsgui"
	filePath := os.Args[1]
	bytes, _ := ioutil.ReadFile(filePath)
	clipboard.WriteAll(fmt.Sprintf("%x", md5.Sum(bytes)))
}
