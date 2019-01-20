package main

import "fmt"

func out1()  {
	fmt.Println("out1")
}

func out2(s string)  {
	fmt.Println(s)
}

func out3(s ...string)  {
	var rs string
	for _, value := range s {
		rs = rs + value + " "
	}
	fmt.Println(rs)
}

func get1(s ...string) string {
	var rs string
	for _, value := range s {
		rs = rs + value + " "
	}
	return rs
}

func test1(f func())  {
	f()
}

func test2(f func(string), p string)  {
	f(p)
}

func test3(f func(string), p ...string)  {
	for _, value := range p {
		f(value)
	}
}

func test4(f func(...string), p ...string) {
	f(p...)
}

func test5(f func(...string)string, p ...string) string {
	return f(p...)
}

func main() {
	test1(out1)
	test2(out2, "out2")
	test3(out2, "out3.1", "out3.2", "out3.3")
	test4(out3, "out4.1", "out4.2", "out4.3")
	rs := test5(get1, "out5.1", "out5.2")
	fmt.Println(rs)
}
