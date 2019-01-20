// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "datid=12345"
	re, error := regexp.Compile("(?:datid=)([0-9]{5})")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(re.FindStringSubmatch(s)[1])
}
