package main

import (
	"fmt"
	"io"
	"sort"
)

const (
	target = 2020
)

// assuming arr is sorted
func part1(arr []int) int {
	var i, j int
	n := len(arr)

	i = 0
	j = n - 1

	for i < n && j >= 0 {
		if arr[i]+arr[j] == target {
			break
		}
		if arr[i]+arr[j] < target {
			i++
		} else {
			j--
		}
	}
	return arr[i] * arr[j]
}

// assuming arr is sorted
func part2(arr []int) int {
	var i, j, k int
	n := len(arr)

	var found bool = false
	for i = 0; i < n; i++ {
		j = i + 1
		k = n - 1
		var tempTarget int = target - arr[i]
		for j < n && k >= 0 {
			if arr[j]+arr[k] == tempTarget {
				found = true
				break
			}
			if arr[j]+arr[k] < tempTarget {
				j++
			} else {
				k--
			}
		}
		if found {
			break
		}
	}
	return arr[i] * arr[j] * arr[k]
}

func main() {
	var ele int

	var check = make([]int, 200, 500)
	var i int = 0
	for {
		_, err := fmt.Scan(&ele)
		if err == io.EOF {
			break
		}
		check[i] = ele
		i++
	}

	// sort the input
	sort.Ints(check)

	part1Ans := part1(check)
	part2Ans := part2(check)

	fmt.Println("Part 1:", part1Ans)
	fmt.Println("Part 2:", part2Ans)
}
