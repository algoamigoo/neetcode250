// ==========================================
// GO SLICES:
// ==========================================

// 1. BASICS & DECLARATION
// - A slice is a dynamic view of an underlying array (pointer + length + capacity).
s1 := []int{1, 2, 3, 4}          // Literal syntax
s2 := make([]int, length, cap)    // Using make()
// len = current number of elements | cap = max size before memory reallocation


// 2. APPENDING & THE SPREAD OPERATOR (...)
s := []int{1, 2}
s = append(s, 3, 4)               // Adds individual elements -> [1 2 3 4]
// *Always reassign: s = append(...)

extra := []int{5, 6}
s = append(s, extra...)           // Unpacks 'extra' slice elements and appends them


// 3. SUB-SLICING (Accessing parts)
base := []int{10, 20, 30, 40, 50}
part := base[1:4]                 // [20, 30, 40] (start inclusive, end exclusive)
// *Shares the same underlying array as 'base'!


// 4. SHARED MEMORY & INDEPENDENT COPIES
part[0] = 99                      // Modifies 'base' as well!
// To make an independent copy (no shared memory):
dst := make([]int, len(base))
copy(dst, base)


// 5. NIL VS EMPTY SLICES
var nilSlice []int                // len = 0, cap = 0, is nil
emptySlice := []int{}             // len = 0, cap = 0, not nil
// *Both are safe to use with append().

//append vs Strings: append() is only for slices. For strings, you use string concatenation (+=).


// 6. REMOVING AN ELEMENT FROM A SLICE (Idiomatic Pattern)
// - To remove an element at index 'i' efficiently:
s = append(s[:i], s[i+1:]...)

// *Why this one-liner is superior to multi-step approaches:
//   - Prevents Bounds Panic: Truncating early (s[:i]) instantly shrinks length, 
//     making s[i+1:] trigger a "slice bounds out of range" panic. This evaluates 
//     both parts using the original length atomically.
//   - Correct Type Handling: The '...' spread operator unpacks the right-side 
//     sub-slice so append() accepts it item-by-item instead of throwing a type error.
//   - Clean & Idiomatic: Avoids extra temporary variables (like 'remaining') and 
//     keeps code concise.