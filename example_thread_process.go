package main

import (
	sw "github.com/Vidog/sewer"

	"fmt"
)

func main() {
	data1 := sw.MakeTupleFloat64Chan(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	data2 := sw.MakeTupleChanFromSlice( []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} )

	fx := func(s sw.StreamI) {
		for x := range s.Items() {
			fmt.Printf("---- Read %#v from stream %#v\n", x.Value, s.Id())
		}
	}

	s1 := sw.NewStreamProcess("Stream 1", data1, func(i float64) float64 {
		return i * 2
	})

	s2 := sw.NewStreamProcess("Stream 2", data2, func(i float64) float64 {
		return i * 3
	})

	go fx(s1)
	go fx(s2)

	for s := range sw.RunStreams(s1, s2) {
		fmt.Printf("-- Stream %#v done\n", s.Id())
	}
}