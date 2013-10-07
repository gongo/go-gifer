Gifer
======

Gifer is maker for animation GIF.

## Required

The following programs must be in the `$PATH`:

- [ffmpeg](http://www.ffmpeg.org/)
- [gifsicle](http://www.lcdf.org/gifsicle/)

## Installation

    $ go get github.com/gongo/go-gifer

## Usage

    $ go-gifer -h
    Usage of go-gifer:
      -delay=50: Set frame delay to TIME (1/100 sec)
      -h=false: Show this message
      -i="": Input movie filename (required)
      -loopcount=0: Set loop extension to N (0 is forever
      -n=100: Number of frames to extract
      -o="gifer.gif": Output gif filename

![demo gif](https://github.com/gongo/go-gifer/raw/master/demo.gif)

## License

MIT License.
