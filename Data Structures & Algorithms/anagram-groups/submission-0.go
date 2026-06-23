func groupAnagrams(strs []string) [][]string {
    groupedStrings := make(map[[26]int][]string)
    for i:=0;i<len(strs);i++ {
        str := strs[i]
        key := [26]int{}
        for _,ch := range str {
            key[ch-'a']++
        }
        groupedStrings[key] = append(groupedStrings[key],str)
    }

    result := [][]string{}
    for _,value := range groupedStrings {
        result = append(result,value)
    }
    return result
}
