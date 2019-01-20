package main

import "fmt"

type Food struct {
	name string
}

func (f *Food) echo()  {
	fmt.Println(f.name)
}

func main() {
	f := Food{
		name: "apple",
	}
	f.echo()
}
