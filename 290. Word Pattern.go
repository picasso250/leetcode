package leetcode

func wordPattern(pattern string, str string) bool {
    words := strings.Fields(str)
    tr:=make(map[rune]string)
    if len(words) != len(pattern) {
        return false
    }
    for i,c := range pattern {
        w:=words[i]
        tr[c]=w
    }
    for i,c := range pattern {
        w:=words[i]
        if tr[c]!=w {
            return false
        }
    }
    return true
}

func main(){
	pattern := "abba"
	
	str := "dog dog dog dog"
	
}