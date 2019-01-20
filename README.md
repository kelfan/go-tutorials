# go-tutorials
tutorial codes for go 

# new slice array of int  ` a := []int{3,9,4,6,2,1,6,4} `
# new channel int  ` sortResult := []<-chan int {} `
# new string array 
var a = [...]string{"a","b"}

# 测试文件命名 _test.go
# 规则: 大写变量才能被外部引用


# 从多个channel中select 
```go
select {
    case v := <-c1:
        fmt.println("channel 1 sends", v)
    case v := <-c2:
        fmt.println("channel 2 sends", v)
    default:
        fmt.println("neither channel was ready)
}
```

# interface 多肽
```go
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

```

# extends 继承
```go
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

```

# object class 
```go 
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
```

# go concurrent channel 
```go
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 500; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		if out, ok := <-ch; ok {
			fmt.Println(out)
		} else {
			break
		}
	}

	// use function
	ch2 := make(chan string)
	go printLine(1000, ch2)

	for {
		if msg, ok2 := <-ch2; ok2 {
			fmt.Println(msg)
		} else {
			break
		}
	}
}

func printLine(num int, ch chan string) <-chan string {
	for j := 0; j < num; j++ {
		ch <- fmt.Sprintln(j, ": -------------")
	}
	close(ch)
	return ch
}
```

# string to double float 
```go
i, err := strconv.ParseFloat(elem, 64)
```

# sort.Ints(a)
```go
	// Create a slice of int
	a := []int{3,9,4,6,2,1,6,4}
	sort.Ints(a)

	for i, v := range a {
		fmt.Println(i, v)
    }
```

# concurrent print
```go
func main() {
	for i := 0; i < 5; i++ {
		go printword3(i)
	}
	time.Sleep(time.Millisecond)
}
func printword3(num int) {
	fmt.Printf("test: %d\n", num)
}
```

# concurrent channel 
```go
func main() {
	ch := make(chan string)
	for i := 0; i < 500; i++ {
		go printword5(i, ch)
	}

	// dead loop and need stop manually
	// output result
	for{
		msg := <- ch
		fmt.Println(msg)
	}

	//time.Sleep(time.Millisecond)
}
func printword5(num int, ch chan string) {
	// dead loop
	for{
		ch <- fmt.Sprintf("test: %d\n", num)
	}
}
```

# function unused result or variable 
```go
func main() {
	// Create a slice of int
	a := []int{3,9,4,6,2,1,6,4}
	sort.Ints(a)

	// print key and value
	for i, v := range a {
		fmt.Println(i, v)
	}

	// print only key
	for _, v := range a{
		fmt.Println(v)
	}
}
```

# := 定义并且赋值变量 等于 `var 变量=`
# <- 取出一个元素
# defer 等于finally 然后先进后出 
```go
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() 
```

# 保存到Excel
作者：aloris
链接：https://www.jianshu.com/p/99456156e1ce
來源：简书
简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。
```go
func SaveSchoolRank(schoolArray [] SchoolStruct.SchoolObj,excelName string,sheetName string)  {

    var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
    var err error

    file,err = xlsx.OpenFile(excelName + ".xlsx")

    if err != nil {
        file = xlsx.NewFile()
        sheet,err = file.AddSheet(sheetName)
    } else {
       sheet = file.Sheet[sheetName]
    }

    if err == nil {

        for i := 0; i < len(schoolArray); i++ {
            obj := schoolArray[i]

            row = sheet.AddRow()
            cell = row.AddCell()
            cell.Value = obj.RankIndex

            cell = row.AddCell()
            cell.Value = obj.SchoolName

            cell = row.AddCell()
            cell.Value = obj.StarLevel

            cell = row.AddCell()
            cell.Value = obj.LocationName

            cell = row.AddCell()
            cell.Value = obj.EnrollOrder

            cell = row.AddCell()
            cell.Value = obj.SchoolType

            cell = row.AddCell()
            cell.Value = obj.UrlAddress


            var tagStr string
            for _,value := range obj.SchoolTags {
                tagStr += "+" + value
            }
            cell = row.AddCell()
            cell.Value = tagStr


            if err != nil {
                fmt.Printf(err.Error())
            }
        }

    }

    err = file.Save(excelName + ".xlsx")
    if err != nil {
        fmt.Printf(err.Error())
    }
}
```

# gotests: 自动生成测试模板
https://github.com/cweill/gotests
例子: https://github.com/polaris1119/The-Golang-Standard-Library-by-Example/blob/master/chapter09/09.1.md

# print vs fmt vs log
- Printing via package log is safe from concurrent goroutines (which plain fmt isn't)
- Log can add timing information automatically.

- The log functions print to stderr by default and can directed to an arbitrary writer. The fmt.Printf function prints to stdout.
- The log functions can print timestamp, source code location and other info.
- The log functions and fmt.Printf are both thread safe, but concurrent writes by fmt.Printf above an OS dependent size can be interleaved.


# pass functions as argument or parameters
```go
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
```



# variable arguments and variable types 
```go
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
```


# object pipe 
```go
type pipeString struct {
	preChan chan string
}

func PipeStringInit(f1 func(chan string)) *pipeString {
	p := pipeString{}
	p.preChan = make(chan string)
	go f1(p.preChan)
	return &p
}

func (p *pipeString) Next(f1 func(chan string, chan string)) *pipeString {
	nextChan := make(chan string)
	go f1(p.preChan, nextChan)
	p.preChan = nextChan
	return p
}

func (p *pipeString) Out() {
	for {
		if out, ok := <-p.preChan; ok {
			fmt.Println(out)
		} else {
			break
		}
	}
}
```
