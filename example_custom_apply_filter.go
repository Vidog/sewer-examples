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
	f1 := func(p Person) Person {
		p.Name = "[" + p.Name + "]"
		return p
	}

	f2 := func(p Person) bool {
		return p.Age >= 50
	}

	data := sw.MakeTupleChan(
		Person{"Dima", 40},
		Person{"Lena", 60},
		Person{"Ivan", 20},
		Person{"Gena", 50},
		Person{"Oleg", 30},
		Person{"Petr", 10},
		Person{"Olga", 70},
	).Apply(f1).Filter(f2)

	for r := range data {
		p := r.Value.(Person)
		fmt.Printf("-- Read %s with age = %d\n", p.Name, p.Age)
	}
}