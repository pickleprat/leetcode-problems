from typing import List
class Solution: 
    def encode(self, strs: List[str]) -> str:
        if len(strs) == 0: 
            return ""
        
        encoded_str: str = ""
        
        for word in strs:
            encoded_str += str(len(word)) + "|" + word
        
        return encoded_str
                    
    def decode(self, s: str) -> List[str]:
        stream_start: bool = False
        i = 0 
        output = []
        current_charcount : int = 0

        while i < len(s):
            if not stream_start and s[i] == "0":
                output.append("")
                i+= 2 

            elif not stream_start and s[i] in "123456789":
                stream_start = True
                amount = ""
                while(s[i] != "|"):
                    amount += s[i]
                    i+=1
                current_charcount = int(amount)
                i+=1 

            elif stream_start and current_charcount != 0 :
                current_word : str = ""
                while(current_charcount > 0):
                    current_word += s[i]
                    current_charcount -= 1     
                    i+= 1           
                stream_start = False  
                output.append(current_word)

        return output 

sln = Solution() 
tc = ["neet", "code", "love", "you"] 

print(sln.encode(tc)) 
print(sln.decode(sln.encode(tc)))

