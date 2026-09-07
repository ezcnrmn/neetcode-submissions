class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:
        combinations = set()

        def cur_to_list(cur):
            l = []
            for n in nums:
                for _ in range(cur[n]):
                    l.append(n)
            return l

        def bt_dfs(cur):
            s = cur["sum"]
            if s > target:
                return
            if s == target:
                combinations.add(tuple(cur_to_list(cur)))
                return

            for n in nums:
                cur[n] += 1
                cur["sum"] += n
                bt_dfs(cur)
                cur[n] -= 1
                cur["sum"] -= n

        cur = defaultdict(int)
        bt_dfs(cur)

        return [list(c) for c in list(combinations)]
