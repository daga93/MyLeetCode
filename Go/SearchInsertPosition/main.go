func searchInsert(nums []int, target int) int {
    length := len(nums)
    base := 0
    nums2 := nums
    for length > 1 {
        center := int(length / 2)
        if target > nums2[center-1] {
            // target in second half

            base = base + center
            nums2 = nums2[center:]
        } else {
            // target in first half
            nums2 = nums2[:center]
        }
        length = len(nums2)
    }
    
    if base > 0 && target == nums[base-1] {return base-1}
    if target > nums[base] { 
        base++ 
    }
    if base < 0 {return 0}
    return base
}
