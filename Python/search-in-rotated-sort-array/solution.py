class Solution:
    def search(self, nums: List[int], target: int) -> int:
        min_bound = 0
        max_bound = len(nums)
        middle = middle_index(min_bound, max_bound)
        while True:
            if len(nums[min_bound:max_bound]) == 1:
                if nums[min_bound] == target:
                    return min_bound
                else:
                    return -1
            a, b = bisect_it(nums[min_bound:max_bound], target, middle)
            min_bound += a
            max_bound += b
            middle = middle_index(min_bound, max_bound)

def is_rotated_part(nums):
    return nums[-1] < nums[0]

def middle_index(min_bound, max_bound):
    return int((max_bound-min_bound)/2)

def bisect_it(nums, target, middle):
    if is_rotated_part(nums[0:middle]):
        if rotated_valid(nums[0:middle], target):
            return 0, -(middle-1)
        else:
            return middle, 0
    else:
        if nums[0:middle][0] <= target <= nums[0:middle][-1]:
            return 0, -1
        return middle, 0


def rotated_valid(nums, target): 
    return target >= nums[0] or target <= nums[-1]
