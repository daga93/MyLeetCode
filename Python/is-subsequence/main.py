class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        t = iter(t)
        
        for c in s:
            t_c = next(t, None)
            while c != t_c:
                t_c = next(t, None)
                if t_c == None: break

            if t_c == None:
                break
            
        else:
            return True

        return False
