package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s1 := scanner.Text()
	scanner.Scan()
	s2 := scanner.Text()
	var chars1, chars2 [26]int := make([]int, 26)
	for _, c := range s1 {
		chars1[c-'a']++
	}
	chars2 := make([]int, 26)
	for _, c := range s2 {
		chars2[c-'a']++
	}
	for i, c := range chars1 {
		if c != chars2[i] {
			fmt.Println("0")
			return
		}
	}
	fmt.Println("1")
}
