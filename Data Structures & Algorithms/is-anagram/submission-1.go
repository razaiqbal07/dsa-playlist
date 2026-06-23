func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var countS [26]int
    for _,ch := range s {
        ch = unicode.ToLower(ch)
        countS[ch-'a']++
    }

    for _,ch := range t {
        countS[ch-'a']--
    }

    for i:=0;i<26;i++ {
        if(countS[i] != 0) {
            return false
        }
    }
    return true
}
