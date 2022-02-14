package main

import (
	"log"
)

var Version = ""

func main() {

	if Version == "" {
		log.Println("没有版本号")

	} else {
		log.Println("版本号为：", Version)
	}
}
