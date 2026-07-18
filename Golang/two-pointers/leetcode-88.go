// [1,2,7,0,0,0]
// [2,5,6]

// i = m= 2
// j= n = 2
// k = M+n-1=5

// work from right to leftP

// i    j.    k 	  nums1[i]>nums2[j].   updted array
// 2	 2	   5	  7>6	  				[1,2,7,0,0,7 ]
// 1	 2		4	   2>6 false			[1,2,7,0,6,7]
// 1	 1		3	   2>5 false			[1,2,7,5,6,7]
// 1	 0.     2.     2>2 false			[1,2,2,5,6,7]
// 1.   -1		1

there are was an edge case where 1,2,7 was there, so we started from the e

First thing we thought of was the pointers should be placed in the beginng, 
whatever we wrote in the code where we gpo wrong, these are specific to each person, 
once. you go through the code, check for your own pattern where you could go wrong. Check
and fix them there itself.

Start making notes out of this.


q2: intervals[i] = [starti, endi]
how many intervals do we need to remove ,so that there ar no overlapping intervals.

[1,2] [2,3]

my test case : [1,2][1,4],[3,5]: remobe 1 ie [1,2]
[1,2][2,3][3,4]: remove 0, none
[0,0], 

overlapping scenarios

[1,4], [2,5], [3,7]
[0,3], [2,6].[3,5], [5,7]

we can sort the entire array based on the starting time and then what will you do
so, in case of [0,3], [2,6].[3,5], [5,7]

I think we also, need to make sure that kaunsa wala remove karna hai, 

