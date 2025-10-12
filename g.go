package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	type city struct{ x, y int }
	var n int
	fmt.Scanln(&n)
	cities := make([]city, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&cities[i].x, &cities[i].y)
	}
	var k, from, to int
	fmt.Scanln(&k)
	fmt.Scanln(&from, &to)
	from--
	to--
	distance := func(i, j int) int { return abs(cities[i].x-cities[j].x) + abs(cities[i].y-cities[j].y) }
	var c = from
	var len = 0
	var path = make(map[int]bool)
	path[c] = true
	for c != to {
		var m = 3000000000
		var step = c
		for i, _ := range cities {
			if distance(c, i) <= k && distance(i, to) < m && !path[i] {
				step = i
				m = distance(i, to)
			}
		}
		if step == c { break }
		c = step
		len++
		path[step] = true
	}
	if c != to { len = -1 }
	fmt.Println(len)
}
