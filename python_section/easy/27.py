class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
       for ind, _ in enumerate(nums):
            while(ind < len(nums) and nums[ind] == val):
                nums.pop(ind)
       return len(nums)
