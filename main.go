package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

var Version = ""

func main() {
	outFile := ""
	flag.StringVar(&outFile, "o", "", "输入文件，默认为当前包名")
	flag.Parse()
	org := make([]string, 0)
	Version = strconv.FormatInt(time.Now().Unix(), 10)
	org = append(org, "build", "-ldflags", fmt.Sprintf("-X 'main.Version=%s'", Version))
	if outFile != "" {
		org = append(org, "-o", outFile)
	}

	log.Println(Version)

	cmd := exec.Command("go", org...)

	outPut, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	log.Println(string(outPut))
}
