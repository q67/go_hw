package main

import (
	"fmt"
	"sort"
)

type number struct {
	id int
}

func includesInt(what int, nums []int) bool {
	for _, num := range nums {
		if num == what {
			return true
		}
	}
	return false
}

func uniqAndSort(data []number) []number {
	var uniqInts = []int{}
	for _, value := range data {
		if !includesInt(value.id, uniqInts) {
			uniqInts = append(uniqInts, value.id)
		}
	}

	sort.Ints(uniqInts)
	uniqAndSortedNumbers := []number{}
	for _, value := range uniqInts {
		var num number
		num.id = value
		uniqAndSortedNumbers = append(uniqAndSortedNumbers, num)
	}

	return uniqAndSortedNumbers
}

func main() {
	values := []number{{3}, {2}, {1}, {2}}
	uniqsAndSorted := uniqAndSort(values)
	fmt.Println(uniqsAndSorted)
}
