class Solution(object):
    def merge(self, nums1, m, nums2, n):
        A = m - 1
        B = n - 1
        C = m + n - 1
        
        while B >= 0:
            if A >= 0 and nums1[A] > nums2[B]:
                nums1[C] = nums1[A]
                B -= 1
            else:
                nums1[C] = nums2[B]
                B -= 1
            C -= 1
