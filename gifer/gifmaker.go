package gifer

import (
	"log"
	"os/exec"
	"strconv"
)

var gifMakerPath string

func init() {
	path, err := exec.LookPath("gifsicle")
	if err != nil {
		log.Fatal(err)
	}
	gifMakerPath = path
}

func gifmaker(frames []string, delay int, loopCount int, saveFileName string) {
	args := []string{"-d", strconv.Itoa(delay), "-l", strconv.Itoa(loopCount), "-o", saveFileName}
	args = append(args, frames...)
	cmd := exec.Command(gifMakerPath, args...)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
