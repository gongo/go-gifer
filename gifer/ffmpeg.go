package gifer

import (
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

type ffmpegResult []byte

func (self ffmpegResult) duration() float32 {
	reg, _ := regexp.Compile(`Duration: (\d+):(\d+):(\d+).(\d+), start`)
	result := reg.FindSubmatch(self)

	hour, _ := strconv.Atoi(string(result[1]))
	minute, _ := strconv.Atoi(string(result[2]))
	second, _ := strconv.Atoi(string(result[3]))
	millsecond, _ := strconv.Atoi(string(result[4]))

	return float32((hour*3600)+(minute*60)+second) + (float32(millsecond) / 100.0)
}

func (self ffmpegResult) framesPerSecond() float32 {
	reg, _ := regexp.Compile(`Video:.*,\s*(\d+)(?:\.(\d+))?\s*fps,`)
	result := reg.FindSubmatch(self)

	count, _ := strconv.ParseFloat(string(result[1]), 32)

	if len(result[2]) != 0 {
		decimalPart, _ := strconv.ParseFloat(string(result[2]), 32)
		count += decimalPart / 100.0
	}

	return float32(count)
}

var ffmpegPath string

func init() {
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatal(err)
	}
	ffmpegPath = path
}

func ffmpeg(args []string) ffmpegResult {
	cmd := exec.Command(ffmpegPath, args...)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Start()

	output, err := ioutil.ReadAll(stderr)
	if err != nil {
		log.Fatal(err)
	}

	cmd.Wait()

	return output
}
