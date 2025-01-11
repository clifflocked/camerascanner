package main

import (
	"fmt"
	"os/exec"
	"strings"	
	//"log"
)

func main() {
	out, err := exec.Command("ls", "/dev/v4l/by-id/").Output()
	if err != nil {}

	devices := strings.Split(string(out), " ")
	//fmt.Println(devices)
	var sNums []string
	serial := ""

	for _, e := range devices {
		start := strings.Index(e, "Camera") + 5
		for j := 2; j < 100; j++ {
			if string(e[start + j]) == "n" {
				sNums = append(sNums, serial)
				serial = ""
				break
			} else {
				serial += string(e[start + j])
			}
		}
	}
	fmt.Print(sNums)

}