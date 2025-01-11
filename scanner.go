package main

import (
	"fmt"
	"os/exec"
	"strings"	
	//"log"
)

func main() {
	out, err := exec.Command("ls", "/dev/v4l/by-id/", "|", "grep", "video0").Output()
	if err != nil {}

	devices := strings.Split(string(out), "\n ")
	var sNums []string
	serial := ""

	for _, e := range devices {
		start := strings.Index(e, "Camera_")
		for j := 1; j < 10; j++ {
			if string(e[start + j]) == "-" {
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