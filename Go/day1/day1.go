package main

import (
	"fmt"
	"io"
)

const (
	target = 2020
)

func main() {
	var ele int
	var check map[int]bool = make(map[int]bool)

	for {
		_, err := fmt.Scan(&ele)
		if err == io.EOF {
			break
		}
		check[ele] = true
		if _, isPresent := check[target-ele]; isPresent {
			fmt.Println(ele * (target - ele))
			break
		}
	}
}
