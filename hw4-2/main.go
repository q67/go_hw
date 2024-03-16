package main

import (
	"fmt"
	"sort"
)

type mystruct struct {
	id     int
	field1 int
	field2 string
	field3 bool
}

func getUniqIds(ids []int) []int {
	uniqs := []int{}

	for _, id := range ids {
		isUniqId := true
		for _, uniqId := range uniqs {
			if id == uniqId {
				isUniqId = false
				break
			}
		}
		if isUniqId {
			uniqs = append(uniqs, id)
		}
	}

	return uniqs
}

func getIndexById(id int, elements []mystruct) int {
	for index, elem := range elements {
		if id == elem.id {
			return index
		}
	}

	return -1
}

func sortAndUniq(sliceIn []mystruct) []mystruct {
	var ids = make([]int, len(sliceIn))
	for i, value := range sliceIn {
		ids[i] = value.id
	}

	sort.Ints(ids)
	ids = getUniqIds(ids)

	sliceOut := []mystruct{}
	for _, id := range ids {
		index := getIndexById(id, sliceIn)
		if index >= 0 {
			sliceOut = append(sliceOut, sliceIn[index])
		}
	}

	return sliceOut
}

func main() {
	values := []mystruct{
		{5, 0, "ttt", false},
		{4, 9, "sus", true},
		{1, 4, "str", false},
		{4, 2, "ws", true},
	}
	uniqsAndSorted := sortAndUniq(values)
	fmt.Println(values, "->", uniqsAndSorted)
}
