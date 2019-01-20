package main

import "fmt"

func testParams3(args ...interface{}) {
	fmt.Println(args[0])
	fmt.Println(args)
	fmt.Println(args...)
}


func main() {
	s := []string{"4", "5", "6"}
	var d []interface{} = []interface{}{s[0], s[1], s[2]}
	testParams3(d...)

}
