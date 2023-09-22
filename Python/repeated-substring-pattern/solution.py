class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:
        substr_length = 1
        while substr_length <= (len(s)/2):
            cur_str = s[0:substr_length]
            if self.check_two_strings(s, cur_str):
                return True
            substr_length += 1


    def check_two_strings(self, main_string, substring):
        nb_chars = len(substring)
        if len(main_string) % nb_chars != 0: return False
        for i in range(0, len(main_string), nb_chars):
            if main_string[i:i+nb_chars] != substring:
                return False
        return True
