package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var seq []rune
var output *bufio.Writer

func generate(n int, opened int, i int) {
	if i == 2*n {
		for _, c := range seq {
			fmt.Fprint(output, string(c))
		}
		fmt.Fprintln(output)
	} else {
		if opened == 2*n-i {
			seq[i] = ')'
			generate(n, opened-1, i+1)
		} else {
			seq[i] = '('
			generate(n, opened+1, i+1)
			if opened > 0 {
				seq[i] = ')'
				generate(n, opened-1, i+1)
			}
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	output = bufio.NewWriter(os.Stdout)
	n, _ := strconv.ParseInt(scanner.Text(), 0, 0)
	seq = make([]rune, 2*n)
	generate(int(n), 0, 0)
	output.Flush()
}
