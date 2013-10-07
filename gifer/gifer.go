package gifer

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"log"
	"os"
	"runtime"
	"sync"
)

type Gifer struct {
	sourceFilename string
	numberOfFrame  int

	// For ffmpeg
	mp4Info        *MP4Info
	numberOfThread int

	// For gifmaker
	delay     int
	loopCount int
}

func NewGifer(filename string) *Gifer {
	self := &Gifer{
		sourceFilename: filename,
		numberOfFrame:  100,
		mp4Info:        GetMP4Info(filename),
		numberOfThread: runtime.GOMAXPROCS(-1),
		delay:          50,
		loopCount:      0,
	}
	return self
}

func (self *Gifer) Run(output string) {
	if self.sourceFilename == "" {
		log.Fatal("Missing input filename")
	}

	if self.NumberOfFrame() < 0 {
		log.Fatal("Gifer.numberOfFrame should not be negative")
	}

	log.Println("Start make frames")
	sources := self.makeSourceFrames()

	log.Println("Start make animation GIF")
	gifmaker(sources, self.Delay(), self.LoopCount(), output)

	log.Println("Complete!!")
}

func (self *Gifer) SourceFileName() string {
	return self.sourceFilename
}

func (self *Gifer) SetNumberOfFrame(n int) {
	self.numberOfFrame = n
}

func (self *Gifer) SetLoopCount(count int) {
	self.loopCount = count
}

func (self *Gifer) SetDelay(delay int) {
	self.delay = delay
}

func (self *Gifer) NumberOfFrame() int {
	return self.numberOfFrame
}

func (self *Gifer) LoopCount() int {
	return self.loopCount
}

func (self *Gifer) Delay() int {
	return self.delay
}

func (self *Gifer) extractFrameNumbers() []int {
	max := self.numberOfFrame
	diff := self.mp4Info.FrameCount() / self.numberOfFrame
	current := 1
	results := make([]int, max)

	for i := 0; i < max; i++ {
		results[i] = current
		current += diff
	}

	return results
}

func (self *Gifer) sourceTempFile(fileNumber int) (file *os.File, err error) {
	template := "%sgifer_%d.gif"
	filename := fmt.Sprintf(template, os.TempDir(), fileNumber)
	return os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
}

func (self *Gifer) makeSourceFrame(number int) (string, error) {
	file, err := self.sourceTempFile(number)

	if err != nil {
		return "", err
	}

	defer file.Close()

	self.mp4Info.SaveFrame(number, file.Name())
	return file.Name(), nil
}

func (self *Gifer) makeSourceFrames() []string {
	indexes := self.extractFrameNumbers()
	filenames := make([]string, len(indexes))
	queue := make(chan bool, self.numberOfThread)

	progressBar := pb.StartNew(len(indexes))
	var wg sync.WaitGroup
	for i, n := range indexes {
		wg.Add(1)
		go func(index, frameNo int) {
			defer wg.Done()
			defer progressBar.Increment()

			queue <- true
			filename, err := self.makeSourceFrame(frameNo)
			if err == nil {
				filenames[index] = filename
			}
			<-queue
		}(i, n)
	}
	wg.Wait()
	progressBar.Finish()

	return filenames
}
