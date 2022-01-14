package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	prev := ""
	for scanner.Scan() {
		s := scanner.Text()
		if s != prev {
			fmt.Println(s)
		}
		prev = s
	}
}
