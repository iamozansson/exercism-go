package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	numberOfBirds := 0
	for i := 0; i < len(birdsPerDay); i++ {
		numberOfBirds = numberOfBirds + birdsPerDay[i]
	}
	return numberOfBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	numberOfBirds := 0
	if week == 1 {
		for i := 0; i < len(birdsPerDay[:7]); i++ {
			numberOfBirds = numberOfBirds + birdsPerDay[i]
		}
	}
	if week == 2 {
		for i := 7; i < len(birdsPerDay[:14]); i++ {
			numberOfBirds = numberOfBirds + birdsPerDay[i]
		}
	}
	if week == 3 {
		for i := 14; i < len(birdsPerDay[:21]); i++ {
			numberOfBirds = numberOfBirds + birdsPerDay[i]
		}
	}
	return numberOfBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}
