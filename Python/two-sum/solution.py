class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        start_index = 0 
        for i, nb_one in enumerate(nums):
            for j in range(i + 1, len(nums)):
                if nb_one + nums[j] == target:
                    return [i, j]
