package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	max := 0
	current := 0
	for scanner.Scan() {
		if scanner.Text() == "1" {
			current++
			if current > max {
				max = current
			}
		} else {
			current = 0
		}
	}
	fmt.Println(max)
}
