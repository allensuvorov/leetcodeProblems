class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        n = len(nums)
        for i in range(n):
            nums[i] = nums[i] * nums[i]
        
        l, r = 0, n - 1
        res = [0] * n

        for i in range (n - 1, -1, -1):
            if nums[l] < nums[r]:
                res[i] = nums[r]
                r -= 1
            else:
                res[i] = nums[l]
                l += 1
        return res
