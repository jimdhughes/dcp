package main

import "log"

type Schedule struct {
	Start int16
	End   int16
}

func calcRoomsRequired(classes []Schedule) int {
	i := 1
	j := 0
	rooms := 1
	maxRooms := 1
	for i < len(classes) && j < len(classes) {
		if classes[i].Start <= classes[j].End {
			rooms = rooms + 1
			i = i + 1
			if rooms > maxRooms {
				maxRooms = rooms
			}
		} else {
			rooms = rooms - 1
			j = j + 1
		}
	}
	return maxRooms
}

func main() {
	classes := []Schedule{
		{Start: 0, End: 50},
		{Start: 30, End: 75},
		{Start: 60, End: 150},
	}
	log.Printf("We need %d rooms\n", calcRoomsRequired(classes))
}
