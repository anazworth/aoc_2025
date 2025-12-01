package day

import "sort"

var days = make(map[int]Day)

func Register(num int, d Day) {
	days[num] = d
}

func AllDays() []Day {
	nums := make([]int, 0, len(days))
	for num := range days {
		nums = append(nums, num)
	}
	sort.Ints(nums)

	result := make([]Day, len(nums))
	for i, num := range nums {
		result[i] = days[num]
	}

	return result
}
