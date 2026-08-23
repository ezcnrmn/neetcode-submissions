class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        mem = set()
        for i in range(len(nums)):
            if nums[i] in mem:
                return True
            mem.add(nums[i])
            if len(mem) > k:
                mem.remove(nums[i-k])
        return False
