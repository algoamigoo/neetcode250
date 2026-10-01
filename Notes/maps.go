// ==========================================
// GO MAPS
// ==========================================

// 1. OVERVIEW & DECLARATION
// - Unordered collection of key-value pairs (hash table). Reference type.
m1 := map[string]int{"apple": 1, "banana": 2} // Map literal
m2 := make(map[string]int)                   // Using make()
// *Note: An uninitialized map is `nil` and cannot be written to; always use `make()` or literal.


// 2. BASIC OPERATIONS (CRUD)
m := make(map[string]int)

// Create / Update
m["score"] = 95             // Adds or updates a key-value pair

// Read
val := m["score"]           // Retrieves value (returns zero-value if key doesn't exist)

// Check if key exists (Comma-ok idiom)
val, exists := m["score"]   // exists is true if key is present, false otherwise

// Delete
delete(m, "score")          // Removes key-value pair (safe even if key doesn't exist)

// Length
count := len(m)             // Number of key-value pairs


// 3. LOOPING & RANGE ITERATION
mLoop := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

// Key and Value
for key, value := range mLoop {
    // fmt.Println(key, value)
}


// 4. IMPORTANT GOTCHAS & RULES
// - Unordered Iteration: Map iteration order in Go is completely random and randomized intentionally on every run. Never rely on a fixed order.
// - Reference Type: Passing a map to a function modifies the original map.
// - Comparable Keys: Keys must be comparable (strings, ints, structs). Slices and maps CANNOT be keys.
// - Concurrency: Maps are NOT thread-safe. Concurrent writes cause a fatal panic, use mutex locks on them

//Map as a value (map[string]map[byte]int) is fully valid, but you must initialize the inner map with make() before writing to it.
//Map as a key (map[map[byte]int]...) is invalid because Go maps are not comparable.