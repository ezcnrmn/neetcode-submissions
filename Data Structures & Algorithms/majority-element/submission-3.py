class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        el, amount = None, 0
        for n in nums:
            if n == el:
                amount += 1
                continue

            if amount == 0:
                el = n
                amount = 1
            else:
                amount -= 1

        return el
