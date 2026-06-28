package DP

// func canJumpRecursive(step int, nums []int) bool{
// 	n := len(nums)

// 	//Base case: reached or passed the last Index
// 	if step >= n-1{
// 		return true
// 	}

// 	for i:= 1; i<=nums[step]; ++{
// 		if canJumpRecursive(i+step, nums){
// 			return true
// 		}
// 	}

// 	return false
// }

func canJump(nums []int) bool {

	//binoculars
	farthest := 0
	numLength := len(nums)
	// har stone pe jake check karo
	for i := 0; i < numLength; i++ {
		// agar binoculars abhi wale point se pehle dekh sakta tha to false
		if i > farthest {
			return false
		}

		//naya bincoculars update karo, jo aur dur dekh sakta hai
		farthest = max(farthest, i+nums[i])
		//agar farthest flag post ya uske aage bhi dekh p[a rha hai to update, yaha numLength -1 hai kyunki farthest bhi index 0 se start hua hai
		//TODO: I make a mistake here
		if farthest >= numLength-1 {
			return true
		}
	}
	return false
}

//dry run
//[2,3,1,1,4]

farthest 	numLength 	i 	i>farthest	i+num[i]	updatedfarthest	farthest>=numsLength-1
0			5.           0.  0>0 no.     0+2 =2			2					2>=4 no
2						 1	 1>2 no		  1+3=4			4					4>=4 yes return true  
