package gifer

import (
	"fmt"
	"log"
	"path/filepath"
)

type MP4Info struct {
	fileName string
	fps      float32
	duration float32
}

func (self *MP4Info) FileName() string {
	return self.fileName
}

func (self *MP4Info) FPS() float32 {
	return self.fps
}

func (self *MP4Info) Duration() float32 {
	return self.duration
}

func (self *MP4Info) FrameCount() int {
	return int(self.Duration() * self.FPS())
}

func (self *MP4Info) Position(frameNumber int) float32 {
	return float32(frameNumber) / self.FPS()
}

func (self *MP4Info) SaveFrame(frameNumber int, pathTo string) {
	startPosition := self.Position(frameNumber)
	args := []string{
		"-ss", fmt.Sprintf("%.2f", startPosition),
		"-i", self.FileName(),
		"-vframes", "1",
		"-f", "image2",
		"-an",
		"-y", pathTo,
	}
	ffmpeg(args)
}

func GetMP4Info(path string) *MP4Info {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatal(err)
	}

	args := []string{
		"-i", path,
		"-f", "null",
		"-vn", "-an",
		"-t", "1",
		"/dev/null",
	}
	result := ffmpeg(args)

	return &MP4Info{
		fileName: absPath,
		fps:      result.framesPerSecond(),
		duration: result.duration(),
	}
}
