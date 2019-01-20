package main

import "fmt"

type Animal struct {
	name string
}

type Cat struct {
	Animal
}

func (a *Animal)echo()  {
	fmt.Println(a.name)
}


func main() {
	c := Cat{Animal{name:"cat"}}
	c.echo()
}
