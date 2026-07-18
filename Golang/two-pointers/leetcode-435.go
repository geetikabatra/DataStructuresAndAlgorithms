

// ] return 0
// [1,2][3,4][5,6] returns 0
// [1,2][2,3][3,4] return 0
// [1,10][2,6][7.9] return 1
// [1,4][2,6][7,9] return 1
// [1,4][2,6][4,7] return 1

// [[1,2],[2,3],[3,4],[1,3]]
// [1,2][1,3][2,3][3,4]

//

// overlapping
// first ontervals end point is greater than second intervals start point

// non o
// [1,2] [2,3]
// 2>2 false non pverlapping
// [1,10] [2,6]
// 10>2 false non overlappinbg
// [1,4][2,6]
// 4>2 true overlapping

// rakhna kaunsa hai and hatana kaunsa hai

// har bar second nhi hataenge,
// jiska end point bada hai, usi ko hatana sahi rahga taki aage comparisons har bar bade end point se honge to zada interval hatane padeneg

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	endSoFar := intervals[0][1] // humne end so far sepearet varable me isliye rakha hai taki jo nterval hatana hai uska end point track kar sake
	flag := 0
	for i := 1; i < len(intervals); i++ {
		startSoFar := intervals[i][0]
        
		if endSoFar > startSoFar {
			endSoFar = min(endSoFar, intervals[i][1])
			flag++
		}else{
			endSoFar = intervals[i][1]
		}
		
	}
	return flag
}
[1,10][2,6][7,9]

i		endSoFar	startSoFar	endSoFar>startSoFar		flag	interavals[i][1]
1		  10			2			10>2. true			1			6				
2.   		6.          7.           6>7. false.                     9

[[1,2],,[2,3],[3,4],[1,3]]

i		endSoFar	startSoFar	endSoFar>startSoFar		flag	interavals[i][1]
1		  2				1			2>1					1		  3
2		  2				2			2>2 false					  3
3		  2				3			2>3 false					  4

[1,2][1,2][1,2]
i		endSoFar	startSoFar	endSoFar>startSoFar		flag	interavals[i][1]
1			2			1			2>1					1				2	
2			2			1								2				2

test case, sort ke baad bhi consexutive wale overlap a kar rhe ho

[1,2][2,3][3,4]
i		endSoFar	startSoFar	endSoFar>startSoFar		flag	interavals[i][1]
1			2			2			2>2								3
2			2			3			2>3								4

[1,2][2,4][3,4]
i		endSoFar	startSoFar	endSoFar>startSoFar		flag	interavals[i][1]
1		  2				2			2>2								4
2		  2				
is case me 2,4 and 3,4 overlapping hai , but comparison me endso far abhi bhi 0th interval ka hi chal rha hai ie 2, flag will never increase

[1,2][2,4][3,4]
i		endSoFar	startSoFar	endSoFar>startSoFar		flag	interavals[i][1]
1		  2				2			2>2					0			4
2		  4				3			4>3					1			4

[[1,2],[2,3],[3,4],[1,3]]