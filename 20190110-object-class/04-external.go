package main

import "./car"

func main() {
	c1 := car.Car{}
	c1.Set("A")
	c1.Echo()

	c1.Test("B")
	c1.Echo()
}
