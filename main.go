package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

var Version = ""

func main() {

	Version = strconv.FormatInt(time.Now().Unix(), 10)

	log.Println(Version)

	cmd := exec.Command("go", "build", "-ldflags", fmt.Sprintf("-X 'main.Version=%s'", Version))

	outPut, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	log.Println(string(outPut))
}
