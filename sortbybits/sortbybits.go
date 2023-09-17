package sortbybits

import (
	"fmt"
	"math/bits"
	"slices"
	"sort"
)

func Hello(name string) string {
	message := fmt.Sprintf("Hello %s", name)
	return message
}

func sortByBits(arr []int) []int {
	// make a map having key as number of 1s in binary representation of number
	m := make(map[int][]int)

	for i := 0; i < len(arr); i++ {
		count := bits.OnesCount(uint(arr[i]))
		// Common idiom to check if key exists in map
		if _, ok := m[count]; ok {
			m[count] = append(m[count], arr[i])
			slices.Sort(m[count])
		} else {
			m[count] = []int{arr[i]}
		}
	}

	var keys []int

	for k, _ := range m {
		keys = append(keys, k)
		sort.Ints(m[k])
	}

	var res []int

	for i := 0; i < len(keys); i++ {
		res = append(res, m[keys[i]]...)
	}

	return res

}
