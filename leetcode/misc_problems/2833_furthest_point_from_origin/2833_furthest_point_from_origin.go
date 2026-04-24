package furthestpointfromorigin

func furthestDistanceFromOrigin(moves string) int {

	nL := 0
	nR := 0
	nD := 0

	for _, e := range moves {
		if string(e) == "L" {
			nL++
		} else if string(e) == "R" {
			nR++
		} else {
			nD++
		}
	}

	if nR == nL {
		return nD
	}

	maxOne := max(nR, nL)
	minOne := min(nR, nL)
	return maxOne - minOne + nD
}

/*
Notes

Problem Link:
https://leetcode.com/problems/furthest-point-from-origin?envType=daily-question&envId=2026-04-24
OR
https://leetcode.com/problems/furthest-point-from-origin

Comments:

1. From the problem statement, it looks like a complex question.
But the solution is to understand that we are moving only on the
x-axis, so to say. And this means that we do not have to move as
per the moves and then find out where we are going.

2. A point to note is that if we have a move towards L, it will
be nullified with a move towards R.

3. Another point to note is that for the dashes, it would make
the most sense to fill them with one direction, so as to move
farther.

4. This makes it a simple matter of calculation of the number
of moves to the right, to the left, and the blanks.

5. Once we know the number of each of L, R, and _, we can
calculate the furthest distance from this. Each pair of L
and R wil cancel themselves out, resulting in a net 0. So
if the number of Ls is equal to that of Rs, then the answer
would be what we fill in the dashes. Which makes the number
of dashes the solution.

6. If they are unequal, then the answer would be
max - min + nd. This is because whatever we are left with
on max - min is the number of single L's or R's that we
will have that aren't cancelled out. Which means we can
then add the number of dashes to find our solution.

*/
