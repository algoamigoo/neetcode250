
 ------------------------------------------------------------------------
 // ==========================================
// GO CONSTRUCTORS: KEY LEARNINGS
// ==========================================

// 1. NO NATIVE CONSTRUCTORS
// - Go does not have class-based constructors (like Java, Python, or C++).
// - Instead, Go uses factory functions—idiomatically named New<TypeName>
//   (or simply Constructor() for platforms like LeetCode).

// 2. RETURN TYPE CONSISTENCY (VALUE VS. POINTER)
// - The function's return signature must match what you actually return:
//   - Returning a value: func Constructor() MyStruct { return MyStruct{...} }
//   - Returning a pointer: func Constructor() *MyStruct { return &MyStruct{...} }
// - Tip: For LeetCode, stick to whatever return type the boilerplate provides.

// 3. INITIALIZING FIELDS IN STRUCT LITERALS
// - When creating the struct inside the constructor, field types must match precisely.
// - Example: If your struct has a slice of a custom type ([]kv), you must
//   initialize it with that exact type ([]kv{}), not an unrelated type ([]int{}).
//
//   Correct:
//   return MyHashMap{
//       myMap: []kv{},
//   }

// 4. WHY USE A CONSTRUCTOR FUNCTION?
// - Encapsulates initialization logic (e.g., pre-allocating slices, setting
//   default values, or configuring internal fields).
// - Ensures your struct is ready to use immediately without runtime panics.

// 5. STRUCT LITERAL INITIALIZATION
// - You can instantiate custom structs inline (e.g., inside an append() or return statement) 
//   using the literal syntax: `TypeName{field: value}`.
// - Parameter vs. Field Mapping: You can map differently named parameters or variables 
//   explicitly to your struct fields (e.g., passing function parameter `value` to struct field `val`):
//   
//   this.myMap = append(this.myMap, kv{key: key, val: value})