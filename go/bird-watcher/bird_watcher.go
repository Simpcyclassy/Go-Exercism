package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    birdCount := 0
    birdsPerDayLength := len(birdsPerDay)
	for i := 0; i < birdsPerDayLength; i++ {
       birdCount +=  birdsPerDay[i]
    }

	return birdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    weekStart := (week-1) * 7
    weekEnd := week * 7
	birdCount := 0
	for i := weekStart; i < weekEnd; i++ {
       birdCount +=  birdsPerDay[i]
    }

	return birdCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	updatedBirdsPerDay := birdsPerDay

	birdsPerDayLength := len(birdsPerDay)
	for i := 0; i < birdsPerDayLength; i += 2 {
		updatedBirdsPerDay[i] = updatedBirdsPerDay[i] + 1
	}

	return updatedBirdsPerDay
}
