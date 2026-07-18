package two-pointers

import "sort"

func numRescueBoats(people []int, limit int) int {
	//sort kar do first
	sort.Ints(people)

	// ek side se ajun left pointer and ek side se bhim right pointer
	// agar arjun + bhim dono limit se kam hai to add boat and arjun ek aage and bhim ek pche
	//otherwise bhim ek piche and boat ++

	leftptr:= 0
	rightptr:=len(people)-1
	boats:=0
	for leftptr<=rightptr{
		if people[leftptr] + people[rightptr]<=limit{
			
			leftptr++
		}
		boats++
		rightptr--
	}

	return boats
}

func main() {
	people:= []int{3,2,2,1}
	limit:= 3

	fmt.Println(numRescueBoats(people, limit))
}
//Dry Run
//[3,2,2,1]
//limit = 3
// sort.Ints = [1,2,3,3]
leftptr	rightPtr	leftptr<=rightptr	people[leftptr] + people[rightptr].  people[leftptr] + people[rightptr]<=limit.  boats	updatedleftPtr  updatedrightptr
0			3			yes						1+3=4								4<=3	yes								1		1				2
1			2			yes						2+3									5<=3.   no								2		1				1
1			1			yes						3+3 = 6								6<=3	no								3		1				0	
//