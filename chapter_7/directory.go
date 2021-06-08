package main

import (
	"fmt"
	"os"
)

/*
func Mkdir(name string, perm FileMode) error

创建名称为name的目录，权限设置是perm，例如0777

func MkdirAll(path string, perm FileMode) error

根据path创建多级子目录，例如astaxie/test1/test2。

func Remove(name string) error

删除名称为name的目录，当目录下有文件或者其他目录时会出错

func RemoveAll(path string) error

根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除。
*/

func main() {
	if err := os.Mkdir("mangolau", 0777); err != nil {
		fmt.Println(err)
		return
	}
	if err := os.MkdirAll("mangolau/test1/test2", 0777); err != nil {
		fmt.Println(err)
		return
	}
	if err := os.RemoveAll("mangolau"); err != nil {
		fmt.Println(err)
		return
	}
}
