class Solution:
    def romanToInt(self, s: str) -> int:
        nb = 0
        translation = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
            'IV': 4,
            'IX': 9,
            'XL': 40,
            'XC': 90,
            'CD': 400,
            'CM': 900,
        }
        pos = 0
        while len(s[pos:])>1:
            if s[pos:pos+2] in translation:
                nb += translation[s[pos:pos+2]]
                pos += 2
            else:
                nb += translation[s[pos]]
                pos += 1
        if len(s[pos:]) == 1:
            nb += translation[s[pos]]
            pos += 1
        return nb
