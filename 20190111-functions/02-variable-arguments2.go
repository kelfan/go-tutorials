package main

import "fmt"

func testParams2(args ...interface{}) {
	for i, v := range args {
		if s, ok := v.(string); ok {
			fmt.Println("----", s)
		}
		if s, ok := v.([]string); ok {
			for i, v := range s {
				fmt.Println(i, "[]----", v)
			}
		}
		fmt.Println(i, v)
	}
}


func main() {
	s := []string{"4", "5", "6"}
	var d []interface{}
	d = append(d, s)
	testParams2(d...)
}
