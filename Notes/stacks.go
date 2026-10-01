// ==========================================
// GO STACKS
// ==========================================

// 1. OVERVIEW & DECLARATION
// - LIFO (Last-In, First-Out) data structure where the last item added is the first removed.
// - Go doesn't have a built-in stack type; we use a slice ([]T) to implement it natively.
st1 := []int{10, 20, 30}          // Stack literal
var st2 []int                    // Nil stack (safe to append to, expands dynamically)
st3 := make([]int, 0)            // Empty slice using make()


// 2. BASIC OPERATIONS (CRUD)
st := make([]int, 0)

// Push (Add to top)
st = append(st, 100)             // Appends element to the end of the slice (top of stack)

// Peek / Top (View top element without removing)
// *Note: Always check `len(st) > 0` first to prevent panic.
top := st[len(st)-1]             

// Pop (Remove and return top element)
top := st[len(st)-1]             // 1. Get the top element
st = st[:len(st)-1]              // 2. Slice off the last element

// IsEmpty Check
isEmpty := len(st) == 0          // True if stack has no elements

// Length / Size
count := len(st)                 // Number of elements in the stack


// 3. LOOPING & ITERATION (DRAINING A STACK)
stLoop := []int{1, 2, 3}

// Process and empty the stack from top to bottom
for len(stLoop) > 0 {
    top := stLoop[len(stLoop)-1] // Access top
    stLoop = stLoop[:len(stLoop)-1] // Pop
    // fmt.Println(top)
}


// - Out of Bounds Panic: Trying to pop or peek from an empty stack (`len(st) == 0`) causes a runtime panic. Always check length first.
// - Slice Capacity: Slices expand underlying arrays automatically when appending, but old memory might linger unless cleaned up or garbage collected.
// - Indexing Mistakes: Never use loop indices (like `i`) to track or slice stack operations. Always use the dynamic `len(st)`.
// - Use Cases: Perfect for matching brackets, parsing expressions, Depth-First Search (DFS), and undo/redo workflows.