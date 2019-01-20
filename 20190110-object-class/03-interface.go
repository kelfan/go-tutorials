package main

import "fmt"

type Plant interface {
	grow()
}

type Grass struct {}
type Tree struct {}

func (g Grass) grow()  {fmt.Println("grass")}
func (t Tree)grow()  {fmt.Println("tree")}

func main() {
	var p Plant
	p = Grass{}
	p.grow()
	p = Tree{}
	p.grow()
}
