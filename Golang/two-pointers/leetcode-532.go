func findPairs(nums []int, k int) int {
    
    l:=0
    r:=1
    n:= len(nums)
    count:=0

    for r<n{
        if abs(nums[l]-nums[r])<k{
            r++
        }else if abs(nums[l]-nums[r])==k){
            r++
			if l==r{  
                continue
            }
            count++
            l++
			for nums[l-1]=nums[l]{
                l++
            }
        }else if  abs(nums[l]-nums[r])>k{
            l++
        }
    }

}

// [1,1,3,4]
// l       r       |nums[l]-nums[r]|    =K       <K      >K      l==r      nums[l-1]=nums[l] count
// 0       0           1-1=0            true                      true         no                         
// 0       1               2-1=1                true               true        yes             
// 0       



// [3,1,4,1,5]
// [1, 5, 8,10]

// l=0 r=len(n-1)

// l = 0 r =0
// while r<n
// |nums[l]-nums[r]|<k r++.  
// nums[l]-nums[r]=k count++, l++ while(I forget this and put an if) nums[l-1]=nums[l] l++ r++, if l==r r++
// nums[l]-nums[r]>k l++ 




// [1, 5, 8,10]
// [1,2,3,4]



import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func findPairs(nums []int, k int) int {
    
    sort.Ints(nums)
    l:=0
    r:=1
    n:= len(nums)
    count:=0

    for r<n && l<n{
        kDiff := nums[r]-nums[l]
        if kDiff<k || l==r{
            r++
        }else if kDiff==k{
            r++
            count++
            l++
			for (l<n && nums[l-1]==nums[l]){
                l++
            }
            // if l>r{
            //     r=l+1
            // }
        }else if  kDiff>k{
            l++
        }
    }
    return count
}
//pair me ek chiz ko fix rakho and dusre ko move karte rehna, 3sum and 4 sum kar ke dekhenge
//hume dhyan rakhna hai ki dusre ko move karne se kahi pura logic n kharab ho jae, so move only 1 pointer
// [1,1,1,2,2]
// l       r       |nums[l]-nums[r]|    =K       <K      >K      l==r      nums[l-1]=nums[l] count
    // 0       1           0               true                                                    1
    // 3       2           1                                 true  
    // 4       2           1                                  true        
    // 5        2                         

//[1,1,3,4,5]
// l       r       |nums[l]-nums[r]|    =K       <K      >K      l==r      nums[l-1]=nums[l] count
    // 0       1                                                                                  1
    // 2       2               0           true                                                   2                                  
    //         3
// [1,1,3,4]
// l       r       |nums[l]-nums[r]|    =K       <K      >K      l==r      nums[l-1]=nums[l] count
// 0       0           1-1=0            true                      true         no                         
// 0       1               2-1=1                true               true        yes             
// 0       



// [3,1,4,1,5]
// [1, 5, 8,10]

// l=0 r=len(n-1)

// l = 0 r =0
// while r<n
// |nums[l]-nums[r]|<k r++.  
// nums[l]-nums[r]=k count++, l++ while(I forget this and put an if) nums[l-1]=nums[l] l++ r++, if l==r r++
// nums[l]-nums[r]>k l++ 



// [1, 5, 8,10]
// [1,2,3,4]



