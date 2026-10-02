// ==========================================
// SORTING IN GO:
// ==========================================

// 1. MODERN SORTING (The `slices` Package - Go 1.21+)
// - Go now has a generic `slices` package which is faster and cleaner than the old `sort` package.
import "cmp"   // Required for comparing types if needed
import "slices"

nums := []int{4, 1, 3, 2}
slices.Sort(nums)                 // Sorts in-place -> [1, 2, 3, 4]


// 2. CLASSIC SORTING (The `sort` Package)
// - Traditional built-in helpers for basic types (mutates the original slice):
import "sort"

intList := []int{5, 2, 6, 3}
sort.Ints(intList)                // [2, 3, 5, 6] (Ascending)

stringList := []string{"banana", "apple", "cherry"}
sort.Strings(stringList)          // ["apple", "banana", "cherry"]

// Reverse sorting (Classic way):
sort.Slice(intList, func(i, j int) bool {
    return intList[i] > intList[j]  // Descending
})


// 3. SORTING CUSTOM STRUCTS
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Bob", 30},
    {"Alice", 25},
    {"Charlie", 35},
}

// Using sort.Slice (Classic & widely used in LeetCode)
sort.Slice(people, func(i, j int) bool {
    return people[i].Age < people[j].Age  // Sort by Age ascending
})

// Using slices.SortFunc (Modern Go 1.21+)
slices.SortFunc(people, func(a, b Person) int {
    return cmp.Compare(a.Age, b.Age)      // Returns -1, 0, or 1
})