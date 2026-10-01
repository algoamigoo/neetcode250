// my soln

func isAnagram(s string, t string) bool {
    // have a map m1 and m2, store count of each alphabet, if its same good else return false

    mp1:=make(map[byte]int)
    mp2:=make(map[byte]int)


    if(len(s)!=len(t)){
    return false
    }

    for idx := range len(s) {
        mp1[s[idx]]++
        mp2[t[idx]]++
    }
    
    for key :=range mp1 {
        if val,ok:=mp2[key]; ok==false || mp1[key]!=val{
            return false
        }
    } 

    for key :=range mp2 {
        if  val,ok:=mp1[key]; ok==false || mp2[key]!=val{
            return false
        }
    }
    return true
}

// best soln

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    counts := make(map[byte]int)

    for i := range s {
        counts[s[i]]++
        counts[t[i]]--
    }

    for _, val := range counts {
        if val != 0 {
            return false
        }
    }

    return true
}