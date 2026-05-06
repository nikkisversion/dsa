package distancebetweenbusstops

func distanceBetweenBusStops(distance []int, start int, destination int) int {

	clockwise := 0
	s := min(start, destination)
	d := max(start, destination)

	wholeDist := 0
	for i, e := range distance {
		wholeDist = wholeDist + e

		if i >= s && i < d {
			clockwise = clockwise + e
		}
	}

	return min(clockwise, wholeDist-clockwise)
}

/*
Notes

Problem Link:
https://leetcode.com/problems/distance-between-bus-stops

Comments:

1. For alternate paths, we have to calculate the
clockwise and anticlockwise distances.

2. Clockwise distance would be straightforward:
from start to destination.

3. For anticlockwise distance, we can subtract the
clockwise one from the total as the path is a circle.

*/
