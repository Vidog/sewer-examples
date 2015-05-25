package main

import (
	sw "github.com/Vidog/sewer"

	"fmt"
	"strings"
	"time"
)

func someAsyncWork() <- chan string {
	out := make(chan string)

	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(500 * time.Millisecond)
			out <- fmt.Sprintf("    Hello %d!    ", i)
		}

		close(out)
	}()

	return out
}

func main() {
	data := sw.MakeTupleChanFromChan( someAsyncWork() ).ApplyAll(strings.ToLower, strings.TrimSpace)

	for r := range data {
		s := r.Value.(string)
		fmt.Printf("-- Read %s\n", s)
	}
}