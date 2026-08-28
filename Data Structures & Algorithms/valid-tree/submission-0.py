class Solution:
	def validTree(self, n: int, edges: List[List[int]]) -> bool:
		if n-1 != len(edges):
			return False
		
		return True