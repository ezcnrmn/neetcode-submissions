class Solution:
    def validPalindrome(self, s: str) -> bool:
        start, end, deleted = 0, len(s)-1, False
        while start < end:
            if s[start] == s[end]:
                start += 1
                end -= 1
                continue
            
            if deleted:
                return False
            
            if s[end-1] == s[start]:
                start += 1
                end -= 2
                delete = True
            elif s[end] == s[start+1]:
                start += 2
                end -= 1
                delete = True
            else:
                return False
        return True