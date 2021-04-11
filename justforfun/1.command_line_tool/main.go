package main

import (
	"flag"
	"fmt"
)

func main() {
	name_pointer := flag.String("name", "yingshaoxo", "your name")
	age_pointer := flag.Int("age", 0, "your age. It's a number")
	king_pointer := flag.Bool("king", false, "true or false")
	flag.Parse()

	fmt.Printf("your name: %s, your age: %v, is king: %v\n", *name_pointer, *age_pointer, *king_pointer)
}
