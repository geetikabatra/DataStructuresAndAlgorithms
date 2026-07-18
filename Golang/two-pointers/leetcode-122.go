
func maxProfit(prices []int) int {
	

	if len(prices)==0{
		return 0

	}
	sumSoFar:=0
	for i=1; i<len(prices);i++{
		if prices[i]>prices[i-1]{
			sumSoFar = sumSoFar+prices[i]-prices[i-1]
		}
	}
	return sumSoFar
}

[1,3,5]

1 pe liya and 5 pe sell kar diya proft at 4
1 leke 3 pe sell kar diya - 2 and 3 pe leke 5 pe sell kar diya to total 4 aa jaega profit, meri approach jo hai and is approach se dono se same proft aa rha hai

[7,1,5,3,6,4]

// start so far maintain kareneg from oth index 
//end so far 2nd element hoga
// as we move forward, if startsofar >endsofar, update startso far to endsofar 
endSofar = -1
startsofar = -1

every increasng order s a bucket, 

[1,3,8,9] 8, 2+5+1 = 8

[1,3,7,9] 8, 2+4+2

----------------------------------------------------------------

number line me jaha bhi stop bana de, sum to 1 se 9 ke 

hum virtually number line hi le rhe hai and jo break karne wale points the unko neglect kar rhe hai

min number of transactions me karna hota to ye logic fail karega
