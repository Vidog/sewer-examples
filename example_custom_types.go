package main

import (
	sw "github.com/Vidog/sewer"

	"fmt"
)

type Person struct {
	Name string
	Age int
}

func main() {
	data := sw.MakeTupleChan(
		Person{"Dima", 40},
		Person{"Lena", 60},
		Person{"Ivan", 20},
		Person{"Gena", 50},
		Person{"Oleg", 30},
		Person{"Petr", 10},
		Person{"Olga", 70},
	)

	s1 := sw.NewStreamProcess("Stream 1", data, func(p Person) Person {
		p.Age = p.Age * 2
		return p
	})

	go sw.ReadStream(s1, func(t sw.Tuple) {
		p := t.Value.(Person)
		fmt.Printf("------ Read %s from stream %s with age = %d\n", p.Name, s1.Id(), p.Age)
	})

	for s := range sw.RunStreams(s1) {
		fmt.Printf("-- Stream %#v done\n", s.Id())
	}
}