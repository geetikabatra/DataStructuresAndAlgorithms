// class Solution {
// public:
//     void sortColors(vector<int>& nums) {
        
//     }
// };



// left at 0the index
// right at last index
// while left<=right
// if left[i]>right[i]//[2,0][2,1][1,0]
//     swap(left, right)
//     left++
// else: //[0,1][0,2][1,2][0,0][1,1][2,2]
//     right--

// [2,0,2,1,1,0]
// leftptr        rightPtr      


// Let us take 3 pointers
// zpointer:=0
// twoPointer:=len(nums)-1
// onePointer:=0

// while onePointer<2pointer:

//     //check for onePointer 
//     if onePointe[i] = 1
//     onepointer++
//     if onePointer[i]=2//swap with two pointer\
//         rightpointer--
//     if onepointer[i]=0 //swap wih zpointer
//         zpointer++
//         onepointer++


// [2,0,2,1,1,0]

// zpointer    1pointer    2pointer    1pointer[i]==0  1pointer[i]=1   1pointer[i]=2  array
// 0               0           5                                           true        [0,0,2,1,1,2]
// 0               0           4               true                                    [0,0,2,1,1,2]
// 1               1           4               true                                    [0,0,2,1,1,2]
// 2               2           4                                           true        [0,0,1,1,2,2]
// 2               3            3.                          true                       [0,0,1,1,2,2]
//                 4                                         true 
//                                                                          true       
                            