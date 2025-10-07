package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const maxCapacity = 512*1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	var s1 = scanner.Text()
	scanner.Scan()
	var s2 = scanner.Text()
    if len(s1) != len(s2) { fmt.Println(0); return }
	var chars1 = make(map[rune]int)
    var chars2 = make(map[rune]int)
	for _, c := range s1 { chars1[c]++ }
	for _, c := range s2 { chars2[c]++ }
    if len(chars1) != len(chars2) { fmt.Println(0); return }
	for c, n := range chars1 {
		if n2, found := chars2[c]; !found || n != n2 {
			fmt.Println(0)
			return
		}
	}
	fmt.Println(1)
}
