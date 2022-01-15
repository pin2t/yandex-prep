package main

import (
	"fmt"
	"sort"
)

type state struct{ city, len int }

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	var n int
	fmt.Scanln(&n)
	cities := make([]struct{ x, y int }, 0, 1000)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scanln(&x, &y)
		cities = append(cities, struct{ x, y int }{x, y})
	}
	var k, from, to int
	fmt.Scanln(&k)
	fmt.Scanln(&from, &to)
	from--
	to--
	seen := map[int]bool{}
	seen[from] = true
	closer := make([]int, n)
	for i := 0; i < n; i++ {
		closer[i] = i
	}
	distance := func(i, j int) int { return abs(cities[i].x-cities[j].x) + abs(cities[i].y-cities[j].y) }
	sort.Slice(closer, func(i, j int) bool {
		return distance(i, to) < distance(j, to)
	})
	paths := make([]state, 0, 1000000)
	paths = append(paths, state{from, 0})
	for i := 0; i < len(paths); i++ {
		s := paths[i]
		seen[s.city] = true
		for _, city := range closer {
			if distance(city, s.city) <= k && !seen[city] {
				if city == to {
					fmt.Println(s.len + 1)
					return
				}
				paths = append(paths, state{city, s.len + 1})
			}
		}
	}
	fmt.Println(-1)
}
