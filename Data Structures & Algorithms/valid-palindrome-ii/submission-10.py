class Solution:
    def validPalindrome(self, s: str) -> bool:
        start, end = 0, len(s)-1
        while start < end:
            if s[start] == s[end]:
                start += 1
                end -= 1
                continue
            
            return self.isPalindrome(s[start:end]) or self.isPalindrome(s[start+1:end+1])
            
        return True
    
    def isPalindrome(self, s: str) -> bool:
        start, end = 0, len(s)-1
        while start < end:
            if s[start] != s[end]:
                return False

            start += 1
            end -= 1
            
        return True