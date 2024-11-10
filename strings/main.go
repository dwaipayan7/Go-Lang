package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple. orange. banana.dwaipayan"
	parts := strings.Split(data, ".")
	fmt.Println(parts)

	str := "one two two three four five"
	count := strings.Count(str,"two");
	fmt.Println(count)

	str = "        Hello, Go!   "
	trimmed := strings.TrimSpace(str);
	fmt.Println(trimmed)

	str1 := "Dwaipayan"
	str2 := "Biswas"
	result := strings.Join([]string{str1,str2}, " ")
	fmt.Println("result:",result)

}
