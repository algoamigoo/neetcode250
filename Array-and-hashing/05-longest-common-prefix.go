func longestCommonPrefix(strs []string) string {
    // we will iterate all the string's ith index and compare if they are equal to ith index of our reference string

    if len(strs) == 0 {
        return ""
    }
    ans := ""
    ref_str := strs[0]

    for idx := range ref_str {
        for j := range strs {
            if idx >= len(strs[j]) || ref_str[idx] != strs[j][idx] {
                return ans
            }
        }
        ans += string(ref_str[idx])
    }
    
    return ans
}