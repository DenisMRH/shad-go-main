//go:build !solution

package hotelbusiness

import "sort"

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	type Event struct {
		Date  int
		Delta int
	}

	// Сбор всех изменений
	var events []Event
	for _, g := range guests {
		events = append(events, Event{g.CheckInDate, 1})
		events = append(events, Event{g.CheckOutDate, -1})
	}

	// Сортируем по дате. На одной дате выезды (-1) идут первыми.
	sort.Slice(events, func(i, j int) bool {
		if events[i].Date == events[j].Date {
			return events[i].Delta < events[j].Delta
		}
		return events[i].Date < events[j].Date
	})

	var result []Load
	currGuests := 0
	prevGuests := -1

	for i := 0; i < len(events); {
		date := events[i].Date
		for i < len(events) && events[i].Date == date {
			currGuests += events[i].Delta
			i++
		}
		// Только если меняется кол-во гостей
		if currGuests != prevGuests {
			result = append(result, Load{StartDate: date, GuestCount: currGuests})
			prevGuests = currGuests
		}
	}

	return result
}
