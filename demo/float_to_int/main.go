package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a float64
	a = 246.5753424657534
	fmt.Println(a)
	str := strconv.FormatFloat(a, 'f', 2, 64)
	b := strings.Split(str, ".")
	fmt.Println(b[0])
}
