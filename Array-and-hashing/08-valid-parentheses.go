func isValid(s string) bool {
	// store last open bracket
	// use stack and pop

	var st []byte // stack

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
            char :=s[i]
			st = append(st, char)
		} else {
			if len(st) == 0 {
				return false
			}
			top := st[len(st)-1]
			if top == '(' && s[i] == ')' || top == '[' && s[i] == ']' || top == '{' && s[i] == '}' {
				st = st[:len(st)-1]
			} else if top == '(' || top == '[' || top == '{' {
				return false
			}
		}
	}
	if len(st) == 0 {
		return true
	}
	return false
}