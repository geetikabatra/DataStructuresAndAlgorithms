package two-pointers

//
//[2,6,4,8,10,9,15]

func findUnsortedSubarray(nums []int) int {
	//arjun left se check karta jaega with intMin for ascending order, jaha bhi break hoga woh left ptr hoga
	//bhim right se decreasing order check karta jjaega wihth inMax, jaha bhi break aaya woh right ptry
	// return right-left+1
	numLength = len(nums)-1
	currMin:= math.intMin
	currMax:= math.maxInt

	leftPtr := 0
	rightPtr := numsLength
	for i=0; i<=numLength; i++{
		if nums[i]<currMin{
			leftPtr = i-1
			break
		}
		currMin = nums[i]
	}

	for i=numLength; i>=0; i--{
		if nums[i]>currMax{
			rightPtr = i+1
			break
		}
		currMax = nums[i]
	}
	if rightPtr==numLength && leftPtr == 0{
		return 0
	}
	return leftPtr-rightPtr+1
}

func main(){

}


//[2,6,4,8,10,9,15]
//numLength = 6
currMin 	currMax			i		nums[i]<currMin				updatedcurrMin			
intMin		intMax			0			2<IntMin= false					2
2							1			6<2								6
6							2			4<6 true						currMin=i-1 = 1

currMax			i		nums[i]>currMax				updatedcurrMax
intMax			6			15>currMax false				15		
15				5			9>15	false					9
10				4			10>9 true						i+1 = 5

return 5-1 +1 =5

//  [1,2,3,4]
currMin 	currMax			i		nums[i]<currMin				updatedcurrMin			leftPtr = 
intMin		intMax			0			1<IntMin false				1
1							1			2<1		 false				2
2							2			3<2		 false				3
3							3			3<4		 false				3
