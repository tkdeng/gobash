package bash

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"
)

type ProgressBar struct {
	running bool
	msg     string
	level   float64
	size    uint

	w int
	h int
}

// NewProgressBar creates a new terminal progress bar
func NewProgressBar(msg string) *ProgressBar {
	pb := &ProgressBar{
		running: true,
		msg:     msg,
		level:   0,
		size:    100,
	}

	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		w = 80
		h = 20
	}

	pb.w = w
	pb.h = h

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt // Wait here until Ctrl+C is pressed
		pb.running = false

		// Perform cleanup and exit gracefully
		fmt.Printf("\033[r\033[%d;1f\033[K\n", h)
		os.Exit(0)
	}()

	fmt.Printf("\033[2J")
	fmt.Printf("\033[1;%dr", h-1)
	fmt.Printf("\033[H\n")

	go func() {
		for pb.running {
			time.Sleep(100 * time.Millisecond)

			if ww, hh, e := term.GetSize(int(os.Stdout.Fd())); e == nil {
				if hh != h || ww != w {
					fmt.Printf("\033[r\033[%d;1f\033[K", h)

					fmt.Printf("\033[2J")
					fmt.Printf("\033[1;%dr", hh-1)
					fmt.Printf("\033[H")
				}

				w = ww
				h = hh

				pb.w = w
				pb.h = h
			}

			fmt.Printf("\r\033[s\r\033[K")
			fmt.Printf("\033[%d;1f", h)

			if pb.level > 100 {
				pb.level = 100
			} else if pb.level < 0 {
				pb.level = 0
			}

			space := w - 12 - len(pb.msg)
			if space < 10 {
				// fmt.Printf("\033[K\n %s (%d%%) \r\033[1A", pb.msg, uint8(pb.level))
				fmt.Printf(" %s (%d%%) ", pb.msg, uint8(pb.level))
			} else {
				prog := int(math.Min(float64(pb.level)*float64(space)/100, float64(space)))
				// fmt.Printf("\033[K\n %s [%s%s] (%d%%) \r\033[1A", pb.msg, strings.Repeat("=", prog), strings.Repeat("-", space-prog), uint8(pb.level))
				fmt.Printf(" %s [%s%s] (%d%%) ", pb.msg, strings.Repeat("=", prog), strings.Repeat("-", space-prog), uint8(pb.level))
			}

			fmt.Printf("\033[u")
		}

		fmt.Printf("\033[r\033[%d;1f\033[K\n", h)
	}()

	return pb
}

// Msg changes the message for the progress bar
func (pb *ProgressBar) Msg(msg string) {
	pb.msg = msg
}

// SetLevel sets the current percentage level of the progress bar
func (pb *ProgressBar) SetLevel(level float64) {
	pb.level = level
}

// SetSize sets the total size of the progress bar for the Step method
func (pb *ProgressBar) SetSize(size uint) {
	pb.size = size
}

// AddSize adds to the total size of the progress bar for the Step method
func (pb *ProgressBar) AddSize(size uint) {
	pb.size += size
}

// Step increments the progress bar by size (default 1)
func (pb *ProgressBar) Step(size ...int) {
	if len(size) == 0 {
		size = []int{1}
	}

	pb.level += 100 / float64(pb.size) * float64(size[0])
}

// Stop stops the progress bar and cleans up the terminal
func (pb *ProgressBar) Stop() {
	pb.running = false
	time.Sleep(250 * time.Millisecond)
	fmt.Printf("\033[r\033[%d;1f\033[K\n", pb.h)
}
