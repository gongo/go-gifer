package main

import (
	"flag"
	"github.com/gongo/go-gifer/gifer"
	"log"
	"os"
)

type giferOptions struct {
	input        string
	output       string
	frameCount   int
	showHelpFlag bool
	delay        int
	loopCount    int
	logLevel     int
}

var opts giferOptions

func init() {
	flag.StringVar(&opts.input, "i", "", "Input movie filename (required)")
	flag.StringVar(&opts.output, "o", "gifer.gif", "Output gif filename")
	flag.IntVar(&opts.frameCount, "n", 100, "Number of frames to extract")
	flag.IntVar(&opts.delay, "delay", 50, "Set frame delay to TIME (1/100 sec)")
	flag.IntVar(&opts.loopCount, "loopcount", 0, "Set loop extension to N (0 is forever")

	flag.BoolVar(&opts.showHelpFlag, "h", false, "Show this message")

	flag.Parse()

	if opts.showHelpFlag {
		flag.Usage()
		os.Exit(0)
	}

	if opts.input == "" {
		log.Fatal("options: Missing input filename")
	}

	if opts.frameCount < 0 {
		log.Fatal("options: -n should not be negative")
	}

	func() {
		f, err := os.Open(opts.input)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func main() {
	giferInstance := gifer.NewGifer(opts.input)
	giferInstance.SetNumberOfFrame(opts.frameCount)
	giferInstance.SetDelay(opts.delay)
	giferInstance.SetLoopCount(opts.loopCount)
	giferInstance.Run(opts.output)
}
