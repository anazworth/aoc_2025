package day

var days = make(map[int]Day)

func Register(num int, d Day) {
	days[num] = d
}

func AllDays() map[int]Day {
	return days
}
