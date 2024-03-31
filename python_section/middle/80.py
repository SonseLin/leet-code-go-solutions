class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        num = 1;
        for i in range(1, len(nums)):
            if num == 1 or nums[i] != nums[num - 2]:
                nums[num] = nums[i]
                num += 1
        return len(nums)
