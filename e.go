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
	var chars1, chars2 [26]int
	for _, c := range s1 {
		chars1[byte(c)-'a']++
	}
	for _, c := range s2 {
		chars2[byte(c)-'a']++
	}
	for i := 0; i < 26; i++ {
		if chars1[i] != chars2[i] {
			fmt.Println("0")
			return
		}
	}
	fmt.Println("1")
}
