package main

import "fmt"

type Name []string

type Man struct {
	name Name
	age  int
}

func main() {
	name := Name{"董洋", "a", "b"}
	man := Man{name: name, age: 12}
	fmt.Println(man.name)
	name[8] = "hhh"
	fmt.Println(man.name)

}
