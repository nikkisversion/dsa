package sumofencryptedintegers

func sumOfEncryptedInt(nums []int) int {

	sum := 0
	for _, e := range nums {
		encrypted := encrypt(e)
		sum += encrypted
	}

	return sum

}

func encrypt(x int) int {

	length := 0
	largestDigit := 0

	for x >= 1 {
		lastDigit := x % 10
		if lastDigit >= largestDigit {
			largestDigit = lastDigit
		}
		length++
		x = x / 10
	}

	res := largestDigit
	for length > 1 {
		res = res*10 + largestDigit
		length--
	}

	return res

}

/*
Notes

Problem Link:
https://leetcode.com/problems/find-the-sum-of-encrypted-integers

Comments:

1. To encrypt a number, we need to find out its
largest digit and the number of digits.

2. Once found, we then build a new number with
the largest digit and the number of digits.

3. The end result is just to add all encrypted
numbers and return the sum.

*/
