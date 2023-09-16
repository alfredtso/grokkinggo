package main

import (
	"fmt"
    "math/bits"
)

func sortByBits(arr []int) []int {
    m := make(map[int][]int)

    for i := 0; i < len(arr); i++ {
        c := bits.OnesCount(uint(arr[i]))
        if _, ok := m[c]; ok {
            m[c] = append(m[c], arr[i])
        } else {
            m[c] = []int{arr[i]}
        }
    }

    res := make([]int)
    return []int;
}

func main() {
    c := sortByBits()
	fmt.Println("Hello, World!", c)
}
