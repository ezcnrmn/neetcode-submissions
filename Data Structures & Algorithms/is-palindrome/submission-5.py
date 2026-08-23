class Solution:
    def isPalindrome(self, s: str) -> bool:
        start, end = 0, len(s)-1
        while start < end:
            while not s[start].isalnum() and start < len(s):
                start+=1
            while not s[end].isalnum() and end >= 0:
                end -= 1
            if s[start].lower() != s[end].lower():
                return False
            start, end = start+1, end-1
        return True