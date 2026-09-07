class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        result = []

        def bt_dfs(cur, index):
            result.append(cur[:])
            for i in range(index, len(nums)):
                cur.append(nums[i])
                bt_dfs(cur, i + 1)
                cur.pop()

        bt_dfs([], 0)
        return result
