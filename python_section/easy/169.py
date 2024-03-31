class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        dn = dict()
        for i in nums:
            if i in dn:
                dn[i] += 1
            else:
                dn[i] = 1
            if dn[i] > len(nums)//2:
                return i
        return -1
