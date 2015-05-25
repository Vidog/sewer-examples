package main

import (
	sw "github.com/Vidog/sewer"

	"fmt"
	"time"
)

func main() {
	s1 := sw.NewStream("Stream #1", func () {
		fmt.Printf("----- Stream #1 working...\n")
		time.Sleep(1000 * time.Millisecond)
	})

	s2 := sw.NewStream("Stream #2", func () {
		fmt.Printf("----- Stream #2 working...\n")
		time.Sleep(2000 * time.Millisecond)
	})

	for s := range sw.RunStreams(s1, s2) {
		fmt.Printf("----- Stream %#v done\n", s.Id())
	}
}