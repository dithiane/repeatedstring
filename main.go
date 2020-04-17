package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a"
	n := 1000000000000
	lengthOfWhole := n / len(str)
	numberOfLeftsymbols := n - lengthOfWhole*len(str)
	aOccurences := strings.Count(str, "a")
	sliceOfLeft := str[0:numberOfLeftsymbols]
	totala := aOccurences*lengthOfWhole + strings.Count(sliceOfLeft, "a")
	fmt.Println(totala)

}
