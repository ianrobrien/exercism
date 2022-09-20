package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	result := 0
	for i := 0; i < len(birdsPerDay); i++ {
		result += birdsPerDay[i]
	}
	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	startWeekIndex := (week - 1) * 7
	result := 0
	for i := startWeekIndex; i < startWeekIndex+7; i++ {
		result += birdsPerDay[i]
	}
	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i++ {
		if i == 0 || i%2 == 0 {
			birdsPerDay[i] += 1
		}
	}
	return birdsPerDay
}
