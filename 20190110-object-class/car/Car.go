package car

import "fmt"

type Car struct {
	name string
}

func (c *Car) Set(model string)  {
	c.name = model
}

func (c *Car) Echo()  {
	fmt.Println("this is a car " + c.name)
}

func (c Car) Test(model string)  {
	c.name = "test " + model
}