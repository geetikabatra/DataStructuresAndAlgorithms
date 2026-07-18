

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

// 🏗️ The Crafty Crane Yard — Memory Palace
// Station 1 — Yard Gate (guard clause)
// Tum ek construction yard ke gate par khade ho. Gate pe bada board tanga hai: "Beam hi nahi hai to kaam hi nahi — 0 bol ke ghar jao."
// → if len(intervals) == 0 { return 0 }
// Station 2 — Sorting Conveyor Belt
// Crafty Crane saare steel beams ko ek conveyor belt pe daalta hai, aur belt unhe START point ke hisaab se chhote-se-bade line mein laga deti hai.
// → sort.Slice(..., intervals[i][0] < intervals[j][0])
// Station 3 — The Red Flagpole
// Pehla beam yard ke beech mein gad jaata hai, aur uske END point par ek laal flagpole gadh diya jaata hai. Yeh flag hamesha batata hai — "abhi tak yard mein sabse door tak kya cover ho chuka hai."
// → endSoFar := intervals[0][1]
// Station 4 — Control Panel Counter
// Crane ke control panel par ek digital counter hai, abhi 0 dikha raha hai. Yeh ginega ki kitne beams scrap yard mein bhejne pade.
// → flag := 0
// Station 5 — Crane Arm Swing
// Crafty apna arm ghumaata hai agli beam ki taraf, line mein doosri beam ki taraf.
// → for i := 1; i < len(intervals); i++
// Station 6 — Start Tag Scanner
// Arm us beam ke START tag ko scan karta hai.
// → startSoFar := intervals[i][0]
// Station 7 — The Alarm Gate
// Crane ek gate se guzarta hai jahan do sensors hain — ek flagpole se wired, doosra naye start-tag se. Agar flagpole, start-tag se AAGE nikla hua hai (endSoFar > startSoFar) → ALARM BAJTA HAI: OVERLAP!
// Station 8a — Overlap: Taraazu (Scale) aur Scrap Yard
// Alarm baajte hi Crafty dono beams ko taraazu par tolta hai — unke END points compare karta hai. Jo beam zyada lamba reach karta hai (bada end), usse uthaake scrap yard mein phek deta hai (discard). Flagpole hamesha chhote reach wale end tak hi rehta hai — kyunki chhota end rakhoge to aage kam beams overlap karenge.
// → endSoFar = min(endSoFar, intervals[i][1]); flag++
// Station 8b — No Overlap: Flagpole Shift
// Alarm nahi baja — matlab beams sirf touch kar rahe hain ya gap mein hain, dono rakhe jaate hain. Crafty flagpole ko ukhaad kar naye beam ke end tak gaad deta hai (chahe woh end chhota ho ya bada — kyunki koi discard nahi hua, sequence continue ho rahi hai).
// → endSoFar = intervals[i][1]
// Station 9 — Exit Gate
// Saare beams process, counter ki final value hi answer hai.
// → return flag

// 🪤 Edge-case side rooms

// Touching beams [1,2][2,3]: alarm sensor pe "2>2 false" likha — touching = allowed, not overlap.
// The tricky one [1,2][2,4][3,4]: yeh room isliye alag rakha hai kyunki yahin bug pakadte hain — agar Station 8b mein flagpole shift na kiya jaaye (matlab endSoFar hamesha stale rahe), to [2,4] vs [3,4] ka overlap kabhi detect nahi hoga.

func eraseOverlapIntervals(intervals [][]int) int {
	// ek crane hai jiska hand move hota hai, sare beams ko ye sort karte hai start point and end point ke according,
	// ab gate pe frst ye check karta hai ki beams hai ki nhi, nhi to 0 return karta hai.
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//start beam ka flag endSofFar me note down kar lo
	endSofar := intervals[0][1]
	//crane ke pas ek counter hai
	flag := 0
	for i := 1; i < len(intervals); i++ {
		//new beams aate ja rhe hai.
		//yaha hum less than end point se hi compare kar rhe hai, kyunki ques says ki end point aur start pont equal pe bhi non overlapping mane jate hai
		if intervals[i][0] < endSofar {
			flag++
			//overlap hone pe hum us beam ko hata denge jiska end bhot dur hai
			// kyunki chance hai ki naye beam sse chote hi honge, to overlap zada badne ke chance hai and
			//chote wale beams hatane padenge to count badh jaega, but question says ki min beams hatane hai
			//to agar hum bada beam hata denge, to hume last chote beam ka end rakhna hai.
			// to do current beams me se chote ka end rakhne ke liye min() use karenge
			endSofar = min(intervals[i][1], endSofar)
		} else {
			//agar overlapping nhi mila to aage badenge as new beam fit ho sakti hai ab properly
			endSofar = intervals[i][1]
		}

	}
	return flag
	// edge cases ki stories
	// 1) beam na aai ho ie len(interval)==0 return 0
	//2)bhot bada interval ho and sabko consume kar rha ho jaise [1,10][2,4][5,6]
	// 3)sab ek dusre se overlapping ho
	// [1,3][2,4][3,5] -> is case me to remove karne padenge
	//4) duplicate intervals ho, jaise [1,2][1,2][1,2]. Is case me sn-1 remove ho jaenge

	//intriguing test case
	//Insight: jab beams nested (Russian-doll style) hoti hain —
	//  har agli beam pichli ke andar hi simat jaati hai — to almost saari hi hatani padti hain,
	// kyunki flagpole har baar aur chhota hota jaata hai, aur nayi beam ka start us chhote flagpole se pehle hi aa jaata hai.
	// sorted: [1,10][2,9][3,8][4,7][5,6]
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