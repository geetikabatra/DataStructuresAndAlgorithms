// package DP

// //[2,3,1,1,4]

// import math

// func jumpRecursive(nums []int, stepIndex, stepCount int) int {
// 	if stepIndex == len(nums)-1 {
// 		return stepCount
// 	}
// 	recurseResult := math.maxInt
// 	for i := 1; i <= nums[stepIndex]; i++ {
// 		recurseResult = min(recurseResult, jumpRecursive(nums, stepIndex+i, stepCount+1))
// 	}
// 	return recurseResult + stepCount
// }

// func jump(nums []int) int {
// 	return jumpRecursive(nums, 0, 0)
// }

// func main() {

// }
// //[2,3,1,1,4]
// // dry run
// call 	stepIndex	stepCOunt	len(nums)-1 stepIndex==len(nums)-1(4)	recurseResult	i 	stepIndex+i StepCount+1 newRecurseResult	returnResult
// 0		0			0				4			no						Inf				1		1			1			pending				pending
// 1 		1			1							no						Inf				1		2			2			pending				pending
// 2		2			2				4			no						Inf				1		3			3			pending				pending
// 3		3			3							no						Inf				1		4			4
// 4		4			4							yes																								4
// 5

func jump2(nums []int) int {
	numLength = len(nums)
	if numLength == 0 {
		return 0
	}

	farthest := 0
	currEnd := 0
	jumps := 0

	for i := 0; i < numLength; i++ {
		farthest = max(farthest, i+nums[i])
		if currEnd == i {
			jumps++
			currEnd = farthest
		}

		if currEnd == numLength-1 {
			break
		}
	}
	return jumps
}

// [2,3,0,1,4]
numLength = 5

i		farthest	currEnd==i 	jumps	 CurrEnd
0			2			yes		 1			2
1			4			no       1			2
2			4			yes		 2			4

// [2,3,0,1,4]

// index	farthest max(farthest, i+nums[i])	currEnd=Index	jumps.           currEnd(isko update nhi karna, best filhal 2nd tak phonch sakte hai, jab tak update karne k need nhi ho)
// 0		2										no				1				2
// 1		4										no				1               2
// 2		4										yes				2				4

// if currIndex == lastIndex

// [1, 3, 0, 0, 1]
// index	farthest currEnd=Index	jumps	currEnd
// 0			1			yes	     1			1
// 1			4			yes		  2			4

// [4, 1, 1, 3, 1, 1,1]
// index	farthest currEnd=Index	jumps	currEnd. currEnd>=umLength-1
// 0			4		yes			1			4
// 1			4		no			1			4
// 2			4		no			1			4
// 3			6		no			1			4
// 4			6		yes			2			6			yes


//[4,1,1,7,1,1,1]
index	farthest	currEnd=Index	jumps	currEnd		currEnd=numLength-1
