class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        n = len(candidates)
        result = set()

        def bt_dfs(path, index):
            cur_sum = sum(path)
            if cur_sum > target:
                return
            elif cur_sum == target:
                result.add(tuple(path))
                return
            
            for i in range(index, n):
                path.append(candidates[i])
                bt_dfs(path, i+1)
                path.pop()
        
        bt_dfs([], 0)

        return [list(t) for t in result]
