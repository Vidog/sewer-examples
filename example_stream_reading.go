package main

import (
	sw "github.com/Vidog/sewer"

	"fmt"
)

func main() {
	data1 := sw.MakeTupleFloat64Chan(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	data2 := sw.MakeTupleChanFromSlice( []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} )

	s1 := sw.NewStreamProcess("Stream 1", data1, func(i float64) float64 {
		return i * 2
	})

	s2 := sw.NewStreamProcess("Stream 2", data2, func(i float64) float64 {
		return i * 3
	})

	go sw.ReadStream(s1, func(t sw.Tuple) {
		fmt.Printf("------ Read %#v from stream 1\n", t.Value)
	})

	go sw.ReadStream(s2, func(t sw.Tuple) {
		fmt.Printf("------ Read %#v from stream 2\n", t.Value)
	})

	for s := range sw.RunStreams(s1, s2) {
		fmt.Printf("-- Stream %#v done\n", s.Id())
	}
}