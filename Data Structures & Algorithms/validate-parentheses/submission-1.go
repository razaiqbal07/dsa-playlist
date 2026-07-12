func isValid(s string) bool {
    check := []string{} 
    
    validMap := map[string]string{
        "(": ")",
        "[": "]",
        "{": "}",
    }
    
    for _, ch := range s {
        strCh := string(ch)
        if strCh == "(" || strCh == "{" || strCh == "[" {
            check = append(check, strCh)
        } else {
			if len(check) == 0 {
                return false
            }
            
            lastIndex := len(check) - 1
            lastItem := check[lastIndex] 
            
            if validMap[lastItem] != strCh {
                return false
            }
            
            check = check[:lastIndex]
        }
    }
    
    return len(check) == 0
}