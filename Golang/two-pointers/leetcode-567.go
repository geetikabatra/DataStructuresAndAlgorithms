package main


func checkInclusions(s1 string, s2 string)bool{
	
	//agar queen ka scroll khali hai to false
	if len(s1) == 0{
		return false
	}
	// agar corrodor bhi khali hai to false
	if len(s2) ==0{
		return false
	}
	
	//aggar scrolln corridor se bada hai to bhi false
	if len(s1) > len(s2) {
		return false
	}

	//window bins banaennge, count 1 and count 2
	//this is an array and not slice, slice me length fix nhi hoti to count1== count2 nhi work karega
	var count1, count2 [26]int	
	var k = len(s1)
	
	// scrooll me each character and corridor ka each character apne apne bins me dal diye
	for i:= 0; i<len(s1); i++{
		count1[s1[i] -'a'] ++
		count2[s2[i]- 'a'] ++
	}
	
	//check if equal, if so return true
	if count1==count2{
		return true
	}

	// now start checking from the window
	for r:=k; r<len(s2); r++{
		//scroll
		count2[s2[r]-'a']++
		//peeche chute hue characters scroll length ke characters hatate jao
		//I did a mistake here, count2 ko count likha tha aur a se compare bhi nhi kiya
		count2[s2[r-k]-'a']--

		//check karte jao ki scroll aur corridor same hai kya
		if count1==count2{
			return true
		}

	}
	return false

}
//s1 = "ab", s2 = "eidbaooo"

// count1-> count1[0]=1,count1[1]=1, count2[4]=1, count2[8]=1
// count1!=count2
// for loop
// we are comparing from second element of s2 because first 2 elements were already covered.
// r=2 r<8 count1[0]=1 and count1[1]=1, count2[3]=1 count2[2-2=0]=0 count1!=count2
// r=3 r<8 count1[0]=1and count1[3]=1, s2=Count2[1]=1 count1!=count2 s2 = db
// r=4  r<8                             count[1]=1 count[2]=0 s2= ba count1 ==count 2 return true

