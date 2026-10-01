func groupAnagrams(strs []string) [][]string {

    // Map to group strings by their character frequency array key.
    mp := make(map[[26]int][]string)

    for i := 0; i < len(strs); i++ {
        curr := strs[i]
        var freq [26]int
        
        for j := 0; j < len(curr); j++ {
            freq[curr[j]-'a']++
        }
        
        mp[freq] = append(mp[freq], curr)
        
        // freq is an array of size 26, storing count of each alphabet at index 0 to 25
        // mp stores a slice of strings having same freq arrays 
    }
    
    var ans [][]string
    for _, val := range mp {
        ans = append(ans, val)
    }
    
    return ans
}

/*
Map as a value (map[string]map[byte]int) is fully valid, but you must initialize the inner map with make() before writing to it.

Map as a key (map[map[byte]int]...) is invalid because Go maps are not comparable.

The workaround: Use a fixed-size array (like [26]int) as your map key instead, since arrays are comparable in Go.

*/