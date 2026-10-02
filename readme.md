# NeetCode 250 — Go Solutions

All [NeetCode 250](https://neetcode.io/practice/practice/neetcode250) problems solved in Go.
Each file holds one problem; the file name is `<position>-<leetcode-slug>.go`.

**Progress:** 12 / 250 solved

| Category | Solved | Total |
| --- | --- | --- |
| [Arrays & Hashing](#arrays--hashing) | 11 | 22 |
| [Two Pointers](#two-pointers) | 0 | 13 |
| [Sliding Window](#sliding-window) | 0 | 9 |
| [Stack](#stack) | 1 | 15 |
| [Binary Search](#binary-search) | 0 | 14 |
| [Linked List](#linked-list) | 0 | 14 |
| [Trees](#trees) | 0 | 23 |
| [Heap / Priority Queue](#heap--priority-queue) | 0 | 12 |
| [Backtracking](#backtracking) | 0 | 16 |
| [Tries](#tries) | 0 | 4 |
| [Graphs](#graphs) | 0 | 21 |
| [Advanced Graphs](#advanced-graphs) | 0 | 10 |
| [1-D Dynamic Programming](#1-d-dynamic-programming) | 0 | 17 |
| [2-D Dynamic Programming](#2-d-dynamic-programming) | 0 | 16 |
| [Greedy](#greedy) | 0 | 14 |
| [Intervals](#intervals) | 0 | 7 |
| [Math & Geometry](#math--geometry) | 0 | 13 |
| [Bit Manipulation](#bit-manipulation) | 0 | 10 |

## Contents

- [Arrays & Hashing](#arrays--hashing) (22)
- [Two Pointers](#two-pointers) (13)
- [Sliding Window](#sliding-window) (9)
- [Stack](#stack) (15)
- [Binary Search](#binary-search) (14)
- [Linked List](#linked-list) (14)
- [Trees](#trees) (23)
- [Heap / Priority Queue](#heap--priority-queue) (12)
- [Backtracking](#backtracking) (16)
- [Tries](#tries) (4)
- [Graphs](#graphs) (21)
- [Advanced Graphs](#advanced-graphs) (10)
- [1-D Dynamic Programming](#1-d-dynamic-programming) (17)
- [2-D Dynamic Programming](#2-d-dynamic-programming) (16)
- [Greedy](#greedy) (14)
- [Intervals](#intervals) (7)
- [Math & Geometry](#math--geometry) (13)
- [Bit Manipulation](#bit-manipulation) (10)

## Arrays & Hashing

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 1929 | [Concatenation of Array](https://leetcode.com/problems/concatenation-of-array/) | Easy | ✅ [`01-concatenation-of-array.go`](./Array-and-hashing/01-concatenation-of-array.go) |
| 217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | Easy | ✅ [`02-contains-duplicate.go`](./Array-and-hashing/02-contains-duplicate.go) |
| 242 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | Easy | ✅ [`03-valid-anagram.go`](./Array-and-hashing/03-valid-anagram.go) |
| 1 | [Two Sum](https://leetcode.com/problems/two-sum/) | Easy | ✅ [`04-two-sum.go`](./Array-and-hashing/04-two-sum.go) |
| 14 | [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) | Easy | ✅ [`05-longest-common-prefix.go`](./Array-and-hashing/05-longest-common-prefix.go) |
| 49 | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | Medium | ✅ [`06-group-anagrams.go`](./Array-and-hashing/06-group-anagrams.go) |
| 27 | [Remove Element](https://leetcode.com/problems/remove-element/) | Easy | ✅ [`07-remove-element.go`](./Array-and-hashing/07-remove-element.go) |
| 169 | [Majority Element](https://leetcode.com/problems/majority-element/) | Easy | ✅ [`08-majority-element.go`](./Array-and-hashing/08-majority-element.go) |
| 705 | [Design HashSet](https://leetcode.com/problems/design-hashset/) | Easy | ✅ [`09-design-hashset.go`](./Array-and-hashing/09-design-hashset.go) |
| 706 | [Design HashMap](https://leetcode.com/problems/design-hashmap/) | Easy | ✅ [`10-design-hashmap.go`](./Array-and-hashing/10-design-hashmap.go) |
| 912 | [Sort an Array](https://leetcode.com/problems/sort-an-array/) | Medium | ✅ [`11-sort-an-array.go`](./Array-and-hashing/11-sort-an-array.go) |
| 75 | [Sort Colors](https://leetcode.com/problems/sort-colors/) | Medium | ⬜ [`12-sort-colors.go`](./Array-and-hashing/12-sort-colors.go) |
| 347 | [Top K Frequent Elements](https://leetcode.com/problems/top-k-frequent-elements/) | Medium | ⬜ [`13-top-k-frequent-elements.go`](./Array-and-hashing/13-top-k-frequent-elements.go) |
| — | [Encode and Decode Strings](https://leetcode.com/problems/encode-and-decode-strings/) | Medium | ⬜ [`14-encode-and-decode-strings.go`](./Array-and-hashing/14-encode-and-decode-strings.go) |
| 304 | [Range Sum Query 2D - Immutable](https://leetcode.com/problems/range-sum-query-2d-immutable/) | Medium | ⬜ [`15-range-sum-query-2d-immutable.go`](./Array-and-hashing/15-range-sum-query-2d-immutable.go) |
| 238 | [Product of Array Except Self](https://leetcode.com/problems/product-of-array-except-self/) | Medium | ⬜ [`16-product-of-array-except-self.go`](./Array-and-hashing/16-product-of-array-except-self.go) |
| 36 | [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/) | Medium | ⬜ [`17-valid-sudoku.go`](./Array-and-hashing/17-valid-sudoku.go) |
| 128 | [Longest Consecutive Sequence](https://leetcode.com/problems/longest-consecutive-sequence/) | Medium | ⬜ [`18-longest-consecutive-sequence.go`](./Array-and-hashing/18-longest-consecutive-sequence.go) |
| 122 | [Best Time to Buy and Sell Stock II](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/) | Medium | ⬜ [`19-best-time-to-buy-and-sell-stock-ii.go`](./Array-and-hashing/19-best-time-to-buy-and-sell-stock-ii.go) |
| 229 | [Majority Element II](https://leetcode.com/problems/majority-element-ii/) | Medium | ⬜ [`20-majority-element-ii.go`](./Array-and-hashing/20-majority-element-ii.go) |
| 560 | [Subarray Sum Equals K](https://leetcode.com/problems/subarray-sum-equals-k/) | Medium | ⬜ [`21-subarray-sum-equals-k.go`](./Array-and-hashing/21-subarray-sum-equals-k.go) |
| 41 | [First Missing Positive](https://leetcode.com/problems/first-missing-positive/) | Hard | ⬜ [`22-first-missing-positive.go`](./Array-and-hashing/22-first-missing-positive.go) |

## Two Pointers

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 344 | [Reverse String](https://leetcode.com/problems/reverse-string/) | Easy | ⬜ [`01-reverse-string.go`](./Two-pointers/01-reverse-string.go) |
| 125 | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) | Easy | ⬜ [`02-valid-palindrome.go`](./Two-pointers/02-valid-palindrome.go) |
| 680 | [Valid Palindrome II](https://leetcode.com/problems/valid-palindrome-ii/) | Easy | ⬜ [`03-valid-palindrome-ii.go`](./Two-pointers/03-valid-palindrome-ii.go) |
| 1768 | [Merge Strings Alternately](https://leetcode.com/problems/merge-strings-alternately/) | Easy | ⬜ [`04-merge-strings-alternately.go`](./Two-pointers/04-merge-strings-alternately.go) |
| 88 | [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) | Easy | ⬜ [`05-merge-sorted-array.go`](./Two-pointers/05-merge-sorted-array.go) |
| 26 | [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | Easy | ⬜ [`06-remove-duplicates-from-sorted-array.go`](./Two-pointers/06-remove-duplicates-from-sorted-array.go) |
| 167 | [Two Sum II - Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | Medium | ⬜ [`07-two-sum-ii-input-array-is-sorted.go`](./Two-pointers/07-two-sum-ii-input-array-is-sorted.go) |
| 15 | [3Sum](https://leetcode.com/problems/3sum/) | Medium | ⬜ [`08-3sum.go`](./Two-pointers/08-3sum.go) |
| 18 | [4Sum](https://leetcode.com/problems/4sum/) | Medium | ⬜ [`09-4sum.go`](./Two-pointers/09-4sum.go) |
| 189 | [Rotate Array](https://leetcode.com/problems/rotate-array/) | Medium | ⬜ [`10-rotate-array.go`](./Two-pointers/10-rotate-array.go) |
| 11 | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | Medium | ⬜ [`11-container-with-most-water.go`](./Two-pointers/11-container-with-most-water.go) |
| 881 | [Boats to Save People](https://leetcode.com/problems/boats-to-save-people/) | Medium | ⬜ [`12-boats-to-save-people.go`](./Two-pointers/12-boats-to-save-people.go) |
| 42 | [Trapping Rain Water](https://leetcode.com/problems/trapping-rain-water/) | Hard | ⬜ [`13-trapping-rain-water.go`](./Two-pointers/13-trapping-rain-water.go) |

## Sliding Window

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 219 | [Contains Duplicate II](https://leetcode.com/problems/contains-duplicate-ii/) | Easy | ⬜ [`01-contains-duplicate-ii.go`](./Sliding-window/01-contains-duplicate-ii.go) |
| 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | Easy | ⬜ [`02-best-time-to-buy-and-sell-stock.go`](./Sliding-window/02-best-time-to-buy-and-sell-stock.go) |
| 3 | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | Medium | ⬜ [`03-longest-substring-without-repeating-characters.go`](./Sliding-window/03-longest-substring-without-repeating-characters.go) |
| 424 | [Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/) | Medium | ⬜ [`04-longest-repeating-character-replacement.go`](./Sliding-window/04-longest-repeating-character-replacement.go) |
| 567 | [Permutation in String](https://leetcode.com/problems/permutation-in-string/) | Medium | ⬜ [`05-permutation-in-string.go`](./Sliding-window/05-permutation-in-string.go) |
| 209 | [Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/) | Medium | ⬜ [`06-minimum-size-subarray-sum.go`](./Sliding-window/06-minimum-size-subarray-sum.go) |
| 658 | [Find K Closest Elements](https://leetcode.com/problems/find-k-closest-elements/) | Medium | ⬜ [`07-find-k-closest-elements.go`](./Sliding-window/07-find-k-closest-elements.go) |
| 76 | [Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) | Hard | ⬜ [`08-minimum-window-substring.go`](./Sliding-window/08-minimum-window-substring.go) |
| 239 | [Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/) | Hard | ⬜ [`09-sliding-window-maximum.go`](./Sliding-window/09-sliding-window-maximum.go) |

## Stack

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 682 | [Baseball Game](https://leetcode.com/problems/baseball-game/) | Easy | ⬜ [`01-baseball-game.go`](./Stack/01-baseball-game.go) |
| 20 | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | Easy | ✅ [`08-valid-parentheses.go`](./Array-and-hashing/08-valid-parentheses.go) |
| 225 | [Implement Stack using Queues](https://leetcode.com/problems/implement-stack-using-queues/) | Easy | ⬜ [`03-implement-stack-using-queues.go`](./Stack/03-implement-stack-using-queues.go) |
| 232 | [Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/) | Easy | ⬜ [`04-implement-queue-using-stacks.go`](./Stack/04-implement-queue-using-stacks.go) |
| 155 | [Min Stack](https://leetcode.com/problems/min-stack/) | Medium | ⬜ [`05-min-stack.go`](./Stack/05-min-stack.go) |
| 150 | [Evaluate Reverse Polish Notation](https://leetcode.com/problems/evaluate-reverse-polish-notation/) | Medium | ⬜ [`06-evaluate-reverse-polish-notation.go`](./Stack/06-evaluate-reverse-polish-notation.go) |
| 22 | [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/) | Medium | ⬜ [`07-generate-parentheses.go`](./Stack/07-generate-parentheses.go) |
| 735 | [Asteroid Collision](https://leetcode.com/problems/asteroid-collision/) | Medium | ⬜ [`08-asteroid-collision.go`](./Stack/08-asteroid-collision.go) |
| 739 | [Daily Temperatures](https://leetcode.com/problems/daily-temperatures/) | Medium | ⬜ [`09-daily-temperatures.go`](./Stack/09-daily-temperatures.go) |
| 901 | [Online Stock Span](https://leetcode.com/problems/online-stock-span/) | Medium | ⬜ [`10-online-stock-span.go`](./Stack/10-online-stock-span.go) |
| 853 | [Car Fleet](https://leetcode.com/problems/car-fleet/) | Medium | ⬜ [`11-car-fleet.go`](./Stack/11-car-fleet.go) |
| 71 | [Simplify Path](https://leetcode.com/problems/simplify-path/) | Medium | ⬜ [`12-simplify-path.go`](./Stack/12-simplify-path.go) |
| 394 | [Decode String](https://leetcode.com/problems/decode-string/) | Medium | ⬜ [`13-decode-string.go`](./Stack/13-decode-string.go) |
| 895 | [Maximum Frequency Stack](https://leetcode.com/problems/maximum-frequency-stack/) | Hard | ⬜ [`14-maximum-frequency-stack.go`](./Stack/14-maximum-frequency-stack.go) |
| 84 | [Largest Rectangle in Histogram](https://leetcode.com/problems/largest-rectangle-in-histogram/) | Hard | ⬜ [`15-largest-rectangle-in-histogram.go`](./Stack/15-largest-rectangle-in-histogram.go) |

## Binary Search

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 704 | [Binary Search](https://leetcode.com/problems/binary-search/) | Easy | ⬜ [`01-binary-search.go`](./Binary-search/01-binary-search.go) |
| 35 | [Search Insert Position](https://leetcode.com/problems/search-insert-position/) | Easy | ⬜ [`02-search-insert-position.go`](./Binary-search/02-search-insert-position.go) |
| 374 | [Guess Number Higher or Lower](https://leetcode.com/problems/guess-number-higher-or-lower/) | Easy | ⬜ [`03-guess-number-higher-or-lower.go`](./Binary-search/03-guess-number-higher-or-lower.go) |
| 69 | [Sqrt(x)](https://leetcode.com/problems/sqrtx/) | Easy | ⬜ [`04-sqrtx.go`](./Binary-search/04-sqrtx.go) |
| 74 | [Search a 2D Matrix](https://leetcode.com/problems/search-a-2d-matrix/) | Medium | ⬜ [`05-search-a-2d-matrix.go`](./Binary-search/05-search-a-2d-matrix.go) |
| 875 | [Koko Eating Bananas](https://leetcode.com/problems/koko-eating-bananas/) | Medium | ⬜ [`06-koko-eating-bananas.go`](./Binary-search/06-koko-eating-bananas.go) |
| 1011 | [Capacity To Ship Packages Within D Days](https://leetcode.com/problems/capacity-to-ship-packages-within-d-days/) | Medium | ⬜ [`07-capacity-to-ship-packages-within-d-days.go`](./Binary-search/07-capacity-to-ship-packages-within-d-days.go) |
| 153 | [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/) | Medium | ⬜ [`08-find-minimum-in-rotated-sorted-array.go`](./Binary-search/08-find-minimum-in-rotated-sorted-array.go) |
| 33 | [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) | Medium | ⬜ [`09-search-in-rotated-sorted-array.go`](./Binary-search/09-search-in-rotated-sorted-array.go) |
| 81 | [Search in Rotated Sorted Array II](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/) | Medium | ⬜ [`10-search-in-rotated-sorted-array-ii.go`](./Binary-search/10-search-in-rotated-sorted-array-ii.go) |
| 981 | [Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/) | Medium | ⬜ [`11-time-based-key-value-store.go`](./Binary-search/11-time-based-key-value-store.go) |
| 410 | [Split Array Largest Sum](https://leetcode.com/problems/split-array-largest-sum/) | Hard | ⬜ [`12-split-array-largest-sum.go`](./Binary-search/12-split-array-largest-sum.go) |
| 4 | [Median of Two Sorted Arrays](https://leetcode.com/problems/median-of-two-sorted-arrays/) | Hard | ⬜ [`13-median-of-two-sorted-arrays.go`](./Binary-search/13-median-of-two-sorted-arrays.go) |
| 1095 | [Find in Mountain Array](https://leetcode.com/problems/find-in-mountain-array/) | Hard | ⬜ [`14-find-in-mountain-array.go`](./Binary-search/14-find-in-mountain-array.go) |

## Linked List

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 206 | [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) | Easy | ⬜ [`01-reverse-linked-list.go`](./Linked-list/01-reverse-linked-list.go) |
| 21 | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | Easy | ⬜ [`02-merge-two-sorted-lists.go`](./Linked-list/02-merge-two-sorted-lists.go) |
| 141 | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) | Easy | ⬜ [`03-linked-list-cycle.go`](./Linked-list/03-linked-list-cycle.go) |
| 143 | [Reorder List](https://leetcode.com/problems/reorder-list/) | Medium | ⬜ [`04-reorder-list.go`](./Linked-list/04-reorder-list.go) |
| 19 | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | Medium | ⬜ [`05-remove-nth-node-from-end-of-list.go`](./Linked-list/05-remove-nth-node-from-end-of-list.go) |
| 138 | [Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/) | Medium | ⬜ [`06-copy-list-with-random-pointer.go`](./Linked-list/06-copy-list-with-random-pointer.go) |
| 2 | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | Medium | ⬜ [`07-add-two-numbers.go`](./Linked-list/07-add-two-numbers.go) |
| 287 | [Find the Duplicate Number](https://leetcode.com/problems/find-the-duplicate-number/) | Medium | ⬜ [`08-find-the-duplicate-number.go`](./Linked-list/08-find-the-duplicate-number.go) |
| 92 | [Reverse Linked List II](https://leetcode.com/problems/reverse-linked-list-ii/) | Medium | ⬜ [`09-reverse-linked-list-ii.go`](./Linked-list/09-reverse-linked-list-ii.go) |
| 622 | [Design Circular Queue](https://leetcode.com/problems/design-circular-queue/) | Medium | ⬜ [`10-design-circular-queue.go`](./Linked-list/10-design-circular-queue.go) |
| 146 | [LRU Cache](https://leetcode.com/problems/lru-cache/) | Medium | ⬜ [`11-lru-cache.go`](./Linked-list/11-lru-cache.go) |
| 460 | [LFU Cache](https://leetcode.com/problems/lfu-cache/) | Hard | ⬜ [`12-lfu-cache.go`](./Linked-list/12-lfu-cache.go) |
| 23 | [Merge k Sorted Lists](https://leetcode.com/problems/merge-k-sorted-lists/) | Hard | ⬜ [`13-merge-k-sorted-lists.go`](./Linked-list/13-merge-k-sorted-lists.go) |
| 25 | [Reverse Nodes in k-Group](https://leetcode.com/problems/reverse-nodes-in-k-group/) | Hard | ⬜ [`14-reverse-nodes-in-k-group.go`](./Linked-list/14-reverse-nodes-in-k-group.go) |

## Trees

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 94 | [Binary Tree Inorder Traversal](https://leetcode.com/problems/binary-tree-inorder-traversal/) | Easy | ⬜ [`01-binary-tree-inorder-traversal.go`](./Trees/01-binary-tree-inorder-traversal.go) |
| 144 | [Binary Tree Preorder Traversal](https://leetcode.com/problems/binary-tree-preorder-traversal/) | Easy | ⬜ [`02-binary-tree-preorder-traversal.go`](./Trees/02-binary-tree-preorder-traversal.go) |
| 145 | [Binary Tree Postorder Traversal](https://leetcode.com/problems/binary-tree-postorder-traversal/) | Easy | ⬜ [`03-binary-tree-postorder-traversal.go`](./Trees/03-binary-tree-postorder-traversal.go) |
| 226 | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) | Easy | ⬜ [`04-invert-binary-tree.go`](./Trees/04-invert-binary-tree.go) |
| 104 | [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | Easy | ⬜ [`05-maximum-depth-of-binary-tree.go`](./Trees/05-maximum-depth-of-binary-tree.go) |
| 543 | [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/) | Easy | ⬜ [`06-diameter-of-binary-tree.go`](./Trees/06-diameter-of-binary-tree.go) |
| 110 | [Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/) | Easy | ⬜ [`07-balanced-binary-tree.go`](./Trees/07-balanced-binary-tree.go) |
| 100 | [Same Tree](https://leetcode.com/problems/same-tree/) | Easy | ⬜ [`08-same-tree.go`](./Trees/08-same-tree.go) |
| 572 | [Subtree of Another Tree](https://leetcode.com/problems/subtree-of-another-tree/) | Easy | ⬜ [`09-subtree-of-another-tree.go`](./Trees/09-subtree-of-another-tree.go) |
| 235 | [Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/) | Medium | ⬜ [`10-lowest-common-ancestor-of-a-binary-search-tree.go`](./Trees/10-lowest-common-ancestor-of-a-binary-search-tree.go) |
| 701 | [Insert into a Binary Search Tree](https://leetcode.com/problems/insert-into-a-binary-search-tree/) | Medium | ⬜ [`11-insert-into-a-binary-search-tree.go`](./Trees/11-insert-into-a-binary-search-tree.go) |
| 450 | [Delete Node in a BST](https://leetcode.com/problems/delete-node-in-a-bst/) | Medium | ⬜ [`12-delete-node-in-a-bst.go`](./Trees/12-delete-node-in-a-bst.go) |
| 102 | [Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/) | Medium | ⬜ [`13-binary-tree-level-order-traversal.go`](./Trees/13-binary-tree-level-order-traversal.go) |
| 199 | [Binary Tree Right Side View](https://leetcode.com/problems/binary-tree-right-side-view/) | Medium | ⬜ [`14-binary-tree-right-side-view.go`](./Trees/14-binary-tree-right-side-view.go) |
| 427 | [Construct Quad Tree](https://leetcode.com/problems/construct-quad-tree/) | Medium | ⬜ [`15-construct-quad-tree.go`](./Trees/15-construct-quad-tree.go) |
| 1448 | [Count Good Nodes in Binary Tree](https://leetcode.com/problems/count-good-nodes-in-binary-tree/) | Medium | ⬜ [`16-count-good-nodes-in-binary-tree.go`](./Trees/16-count-good-nodes-in-binary-tree.go) |
| 98 | [Validate Binary Search Tree](https://leetcode.com/problems/validate-binary-search-tree/) | Medium | ⬜ [`17-validate-binary-search-tree.go`](./Trees/17-validate-binary-search-tree.go) |
| 230 | [Kth Smallest Element in a BST](https://leetcode.com/problems/kth-smallest-element-in-a-bst/) | Medium | ⬜ [`18-kth-smallest-element-in-a-bst.go`](./Trees/18-kth-smallest-element-in-a-bst.go) |
| 105 | [Construct Binary Tree from Preorder and Inorder Traversal](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) | Medium | ⬜ [`19-construct-binary-tree-from-preorder-and-inorder-traversal.go`](./Trees/19-construct-binary-tree-from-preorder-and-inorder-traversal.go) |
| 337 | [House Robber III](https://leetcode.com/problems/house-robber-iii/) | Medium | ⬜ [`20-house-robber-iii.go`](./Trees/20-house-robber-iii.go) |
| 1325 | [Delete Leaves With a Given Value](https://leetcode.com/problems/delete-leaves-with-a-given-value/) | Medium | ⬜ [`21-delete-leaves-with-a-given-value.go`](./Trees/21-delete-leaves-with-a-given-value.go) |
| 124 | [Binary Tree Maximum Path Sum](https://leetcode.com/problems/binary-tree-maximum-path-sum/) | Hard | ⬜ [`22-binary-tree-maximum-path-sum.go`](./Trees/22-binary-tree-maximum-path-sum.go) |
| 297 | [Serialize and Deserialize Binary Tree](https://leetcode.com/problems/serialize-and-deserialize-binary-tree/) | Hard | ⬜ [`23-serialize-and-deserialize-binary-tree.go`](./Trees/23-serialize-and-deserialize-binary-tree.go) |

## Heap / Priority Queue

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 703 | [Kth Largest Element in a Stream](https://leetcode.com/problems/kth-largest-element-in-a-stream/) | Easy | ⬜ [`01-kth-largest-element-in-a-stream.go`](./Heap-priority-queue/01-kth-largest-element-in-a-stream.go) |
| 1046 | [Last Stone Weight](https://leetcode.com/problems/last-stone-weight/) | Easy | ⬜ [`02-last-stone-weight.go`](./Heap-priority-queue/02-last-stone-weight.go) |
| 973 | [K Closest Points to Origin](https://leetcode.com/problems/k-closest-points-to-origin/) | Medium | ⬜ [`03-k-closest-points-to-origin.go`](./Heap-priority-queue/03-k-closest-points-to-origin.go) |
| 215 | [Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/) | Medium | ⬜ [`04-kth-largest-element-in-an-array.go`](./Heap-priority-queue/04-kth-largest-element-in-an-array.go) |
| 621 | [Task Scheduler](https://leetcode.com/problems/task-scheduler/) | Medium | ⬜ [`05-task-scheduler.go`](./Heap-priority-queue/05-task-scheduler.go) |
| 355 | [Design Twitter](https://leetcode.com/problems/design-twitter/) | Medium | ⬜ [`06-design-twitter.go`](./Heap-priority-queue/06-design-twitter.go) |
| 1834 | [Single-Threaded CPU](https://leetcode.com/problems/single-threaded-cpu/) | Medium | ⬜ [`07-single-threaded-cpu.go`](./Heap-priority-queue/07-single-threaded-cpu.go) |
| 767 | [Reorganize String](https://leetcode.com/problems/reorganize-string/) | Medium | ⬜ [`08-reorganize-string.go`](./Heap-priority-queue/08-reorganize-string.go) |
| 1405 | [Longest Happy String](https://leetcode.com/problems/longest-happy-string/) | Medium | ⬜ [`09-longest-happy-string.go`](./Heap-priority-queue/09-longest-happy-string.go) |
| 1094 | [Car Pooling](https://leetcode.com/problems/car-pooling/) | Medium | ⬜ [`10-car-pooling.go`](./Heap-priority-queue/10-car-pooling.go) |
| 295 | [Find Median from Data Stream](https://leetcode.com/problems/find-median-from-data-stream/) | Hard | ⬜ [`11-find-median-from-data-stream.go`](./Heap-priority-queue/11-find-median-from-data-stream.go) |
| 502 | [IPO](https://leetcode.com/problems/ipo/) | Hard | ⬜ [`12-ipo.go`](./Heap-priority-queue/12-ipo.go) |

## Backtracking

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 1863 | [Sum of All Subset XOR Totals](https://leetcode.com/problems/sum-of-all-subset-xor-totals/) | Easy | ⬜ [`01-sum-of-all-subset-xor-totals.go`](./Backtracking/01-sum-of-all-subset-xor-totals.go) |
| 78 | [Subsets](https://leetcode.com/problems/subsets/) | Medium | ⬜ [`02-subsets.go`](./Backtracking/02-subsets.go) |
| 39 | [Combination Sum](https://leetcode.com/problems/combination-sum/) | Medium | ⬜ [`03-combination-sum.go`](./Backtracking/03-combination-sum.go) |
| 40 | [Combination Sum II](https://leetcode.com/problems/combination-sum-ii/) | Medium | ⬜ [`04-combination-sum-ii.go`](./Backtracking/04-combination-sum-ii.go) |
| 77 | [Combinations](https://leetcode.com/problems/combinations/) | Medium | ⬜ [`05-combinations.go`](./Backtracking/05-combinations.go) |
| 46 | [Permutations](https://leetcode.com/problems/permutations/) | Medium | ⬜ [`06-permutations.go`](./Backtracking/06-permutations.go) |
| 90 | [Subsets II](https://leetcode.com/problems/subsets-ii/) | Medium | ⬜ [`07-subsets-ii.go`](./Backtracking/07-subsets-ii.go) |
| 47 | [Permutations II](https://leetcode.com/problems/permutations-ii/) | Medium | ⬜ [`08-permutations-ii.go`](./Backtracking/08-permutations-ii.go) |
| 79 | [Word Search](https://leetcode.com/problems/word-search/) | Medium | ⬜ [`09-word-search.go`](./Backtracking/09-word-search.go) |
| 131 | [Palindrome Partitioning](https://leetcode.com/problems/palindrome-partitioning/) | Medium | ⬜ [`10-palindrome-partitioning.go`](./Backtracking/10-palindrome-partitioning.go) |
| 17 | [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | Medium | ⬜ [`11-letter-combinations-of-a-phone-number.go`](./Backtracking/11-letter-combinations-of-a-phone-number.go) |
| 473 | [Matchsticks to Square](https://leetcode.com/problems/matchsticks-to-square/) | Medium | ⬜ [`12-matchsticks-to-square.go`](./Backtracking/12-matchsticks-to-square.go) |
| 698 | [Partition to K Equal Sum Subsets](https://leetcode.com/problems/partition-to-k-equal-sum-subsets/) | Medium | ⬜ [`13-partition-to-k-equal-sum-subsets.go`](./Backtracking/13-partition-to-k-equal-sum-subsets.go) |
| 51 | [N-Queens](https://leetcode.com/problems/n-queens/) | Hard | ⬜ [`14-n-queens.go`](./Backtracking/14-n-queens.go) |
| 52 | [N-Queens II](https://leetcode.com/problems/n-queens-ii/) | Hard | ⬜ [`15-n-queens-ii.go`](./Backtracking/15-n-queens-ii.go) |
| 140 | [Word Break II](https://leetcode.com/problems/word-break-ii/) | Hard | ⬜ [`16-word-break-ii.go`](./Backtracking/16-word-break-ii.go) |

## Tries

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 208 | [Implement Trie (Prefix Tree)](https://leetcode.com/problems/implement-trie-prefix-tree/) | Medium | ⬜ [`01-implement-trie-prefix-tree.go`](./Tries/01-implement-trie-prefix-tree.go) |
| 211 | [Design Add and Search Words Data Structure](https://leetcode.com/problems/design-add-and-search-words-data-structure/) | Medium | ⬜ [`02-design-add-and-search-words-data-structure.go`](./Tries/02-design-add-and-search-words-data-structure.go) |
| 2707 | [Extra Characters in a String](https://leetcode.com/problems/extra-characters-in-a-string/) | Medium | ⬜ [`03-extra-characters-in-a-string.go`](./Tries/03-extra-characters-in-a-string.go) |
| 212 | [Word Search II](https://leetcode.com/problems/word-search-ii/) | Hard | ⬜ [`04-word-search-ii.go`](./Tries/04-word-search-ii.go) |

## Graphs

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 463 | [Island Perimeter](https://leetcode.com/problems/island-perimeter/) | Easy | ⬜ [`01-island-perimeter.go`](./Graphs/01-island-perimeter.go) |
| 953 | [Verifying an Alien Dictionary](https://leetcode.com/problems/verifying-an-alien-dictionary/) | Easy | ⬜ [`02-verifying-an-alien-dictionary.go`](./Graphs/02-verifying-an-alien-dictionary.go) |
| 997 | [Find the Town Judge](https://leetcode.com/problems/find-the-town-judge/) | Easy | ⬜ [`03-find-the-town-judge.go`](./Graphs/03-find-the-town-judge.go) |
| 200 | [Number of Islands](https://leetcode.com/problems/number-of-islands/) | Medium | ⬜ [`04-number-of-islands.go`](./Graphs/04-number-of-islands.go) |
| 695 | [Max Area of Island](https://leetcode.com/problems/max-area-of-island/) | Medium | ⬜ [`05-max-area-of-island.go`](./Graphs/05-max-area-of-island.go) |
| 133 | [Clone Graph](https://leetcode.com/problems/clone-graph/) | Medium | ⬜ [`06-clone-graph.go`](./Graphs/06-clone-graph.go) |
| — | [Walls and Gates](https://leetcode.com/problems/walls-and-gates/) | Medium | ⬜ [`07-walls-and-gates.go`](./Graphs/07-walls-and-gates.go) |
| 994 | [Rotting Oranges](https://leetcode.com/problems/rotting-oranges/) | Medium | ⬜ [`08-rotting-oranges.go`](./Graphs/08-rotting-oranges.go) |
| 417 | [Pacific Atlantic Water Flow](https://leetcode.com/problems/pacific-atlantic-water-flow/) | Medium | ⬜ [`09-pacific-atlantic-water-flow.go`](./Graphs/09-pacific-atlantic-water-flow.go) |
| 130 | [Surrounded Regions](https://leetcode.com/problems/surrounded-regions/) | Medium | ⬜ [`10-surrounded-regions.go`](./Graphs/10-surrounded-regions.go) |
| 752 | [Open the Lock](https://leetcode.com/problems/open-the-lock/) | Medium | ⬜ [`11-open-the-lock.go`](./Graphs/11-open-the-lock.go) |
| 207 | [Course Schedule](https://leetcode.com/problems/course-schedule/) | Medium | ⬜ [`12-course-schedule.go`](./Graphs/12-course-schedule.go) |
| 210 | [Course Schedule II](https://leetcode.com/problems/course-schedule-ii/) | Medium | ⬜ [`13-course-schedule-ii.go`](./Graphs/13-course-schedule-ii.go) |
| — | [Graph Valid Tree](https://leetcode.com/problems/graph-valid-tree/) | Medium | ⬜ [`14-graph-valid-tree.go`](./Graphs/14-graph-valid-tree.go) |
| 1462 | [Course Schedule IV](https://leetcode.com/problems/course-schedule-iv/) | Medium | ⬜ [`15-course-schedule-iv.go`](./Graphs/15-course-schedule-iv.go) |
| — | [Number of Connected Components in an Undirected Graph](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/) | Medium | ⬜ [`16-number-of-connected-components-in-an-undirected-graph.go`](./Graphs/16-number-of-connected-components-in-an-undirected-graph.go) |
| 684 | [Redundant Connection](https://leetcode.com/problems/redundant-connection/) | Medium | ⬜ [`17-redundant-connection.go`](./Graphs/17-redundant-connection.go) |
| 721 | [Accounts Merge](https://leetcode.com/problems/accounts-merge/) | Medium | ⬜ [`18-accounts-merge.go`](./Graphs/18-accounts-merge.go) |
| 399 | [Evaluate Division](https://leetcode.com/problems/evaluate-division/) | Medium | ⬜ [`19-evaluate-division.go`](./Graphs/19-evaluate-division.go) |
| 310 | [Minimum Height Trees](https://leetcode.com/problems/minimum-height-trees/) | Medium | ⬜ [`20-minimum-height-trees.go`](./Graphs/20-minimum-height-trees.go) |
| 127 | [Word Ladder](https://leetcode.com/problems/word-ladder/) | Hard | ⬜ [`21-word-ladder.go`](./Graphs/21-word-ladder.go) |

## Advanced Graphs

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 1631 | [Path With Minimum Effort](https://leetcode.com/problems/path-with-minimum-effort/) | Medium | ⬜ [`01-path-with-minimum-effort.go`](./Advanced-graphs/01-path-with-minimum-effort.go) |
| 743 | [Network Delay Time](https://leetcode.com/problems/network-delay-time/) | Medium | ⬜ [`02-network-delay-time.go`](./Advanced-graphs/02-network-delay-time.go) |
| 332 | [Reconstruct Itinerary](https://leetcode.com/problems/reconstruct-itinerary/) | Hard | ⬜ [`03-reconstruct-itinerary.go`](./Advanced-graphs/03-reconstruct-itinerary.go) |
| 1584 | [Min Cost to Connect All Points](https://leetcode.com/problems/min-cost-to-connect-all-points/) | Medium | ⬜ [`04-min-cost-to-connect-all-points.go`](./Advanced-graphs/04-min-cost-to-connect-all-points.go) |
| 778 | [Swim in Rising Water](https://leetcode.com/problems/swim-in-rising-water/) | Hard | ⬜ [`05-swim-in-rising-water.go`](./Advanced-graphs/05-swim-in-rising-water.go) |
| — | [Alien Dictionary](https://leetcode.com/problems/alien-dictionary/) | Hard | ⬜ [`06-alien-dictionary.go`](./Advanced-graphs/06-alien-dictionary.go) |
| 787 | [Cheapest Flights Within K Stops](https://leetcode.com/problems/cheapest-flights-within-k-stops/) | Medium | ⬜ [`07-cheapest-flights-within-k-stops.go`](./Advanced-graphs/07-cheapest-flights-within-k-stops.go) |
| 1489 | [Find Critical and Pseudo-Critical Edges in Minimum Spanning Tree](https://leetcode.com/problems/find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree/) | Hard | ⬜ [`08-find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree.go`](./Advanced-graphs/08-find-critical-and-pseudo-critical-edges-in-minimum-spanning-tree.go) |
| 2392 | [Build a Matrix With Conditions](https://leetcode.com/problems/build-a-matrix-with-conditions/) | Hard | ⬜ [`09-build-a-matrix-with-conditions.go`](./Advanced-graphs/09-build-a-matrix-with-conditions.go) |
| 2709 | [Greatest Common Divisor Traversal](https://leetcode.com/problems/greatest-common-divisor-traversal/) | Hard | ⬜ [`10-greatest-common-divisor-traversal.go`](./Advanced-graphs/10-greatest-common-divisor-traversal.go) |

## 1-D Dynamic Programming

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 70 | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) | Easy | ⬜ [`01-climbing-stairs.go`](./1-D-dynamic-programming/01-climbing-stairs.go) |
| 746 | [Min Cost Climbing Stairs](https://leetcode.com/problems/min-cost-climbing-stairs/) | Easy | ⬜ [`02-min-cost-climbing-stairs.go`](./1-D-dynamic-programming/02-min-cost-climbing-stairs.go) |
| 1137 | [N-th Tribonacci Number](https://leetcode.com/problems/n-th-tribonacci-number/) | Easy | ⬜ [`03-n-th-tribonacci-number.go`](./1-D-dynamic-programming/03-n-th-tribonacci-number.go) |
| 198 | [House Robber](https://leetcode.com/problems/house-robber/) | Medium | ⬜ [`04-house-robber.go`](./1-D-dynamic-programming/04-house-robber.go) |
| 213 | [House Robber II](https://leetcode.com/problems/house-robber-ii/) | Medium | ⬜ [`05-house-robber-ii.go`](./1-D-dynamic-programming/05-house-robber-ii.go) |
| 5 | [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) | Medium | ⬜ [`06-longest-palindromic-substring.go`](./1-D-dynamic-programming/06-longest-palindromic-substring.go) |
| 647 | [Palindromic Substrings](https://leetcode.com/problems/palindromic-substrings/) | Medium | ⬜ [`07-palindromic-substrings.go`](./1-D-dynamic-programming/07-palindromic-substrings.go) |
| 91 | [Decode Ways](https://leetcode.com/problems/decode-ways/) | Medium | ⬜ [`08-decode-ways.go`](./1-D-dynamic-programming/08-decode-ways.go) |
| 322 | [Coin Change](https://leetcode.com/problems/coin-change/) | Medium | ⬜ [`09-coin-change.go`](./1-D-dynamic-programming/09-coin-change.go) |
| 152 | [Maximum Product Subarray](https://leetcode.com/problems/maximum-product-subarray/) | Medium | ⬜ [`10-maximum-product-subarray.go`](./1-D-dynamic-programming/10-maximum-product-subarray.go) |
| 139 | [Word Break](https://leetcode.com/problems/word-break/) | Medium | ⬜ [`11-word-break.go`](./1-D-dynamic-programming/11-word-break.go) |
| 300 | [Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/) | Medium | ⬜ [`12-longest-increasing-subsequence.go`](./1-D-dynamic-programming/12-longest-increasing-subsequence.go) |
| 416 | [Partition Equal Subset Sum](https://leetcode.com/problems/partition-equal-subset-sum/) | Medium | ⬜ [`13-partition-equal-subset-sum.go`](./1-D-dynamic-programming/13-partition-equal-subset-sum.go) |
| 377 | [Combination Sum IV](https://leetcode.com/problems/combination-sum-iv/) | Medium | ⬜ [`14-combination-sum-iv.go`](./1-D-dynamic-programming/14-combination-sum-iv.go) |
| 279 | [Perfect Squares](https://leetcode.com/problems/perfect-squares/) | Medium | ⬜ [`15-perfect-squares.go`](./1-D-dynamic-programming/15-perfect-squares.go) |
| 343 | [Integer Break](https://leetcode.com/problems/integer-break/) | Medium | ⬜ [`16-integer-break.go`](./1-D-dynamic-programming/16-integer-break.go) |
| 1406 | [Stone Game III](https://leetcode.com/problems/stone-game-iii/) | Hard | ⬜ [`17-stone-game-iii.go`](./1-D-dynamic-programming/17-stone-game-iii.go) |

## 2-D Dynamic Programming

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 62 | [Unique Paths](https://leetcode.com/problems/unique-paths/) | Medium | ⬜ [`01-unique-paths.go`](./2-D-dynamic-programming/01-unique-paths.go) |
| 63 | [Unique Paths II](https://leetcode.com/problems/unique-paths-ii/) | Medium | ⬜ [`02-unique-paths-ii.go`](./2-D-dynamic-programming/02-unique-paths-ii.go) |
| 64 | [Minimum Path Sum](https://leetcode.com/problems/minimum-path-sum/) | Medium | ⬜ [`03-minimum-path-sum.go`](./2-D-dynamic-programming/03-minimum-path-sum.go) |
| 1143 | [Longest Common Subsequence](https://leetcode.com/problems/longest-common-subsequence/) | Medium | ⬜ [`04-longest-common-subsequence.go`](./2-D-dynamic-programming/04-longest-common-subsequence.go) |
| 1049 | [Last Stone Weight II](https://leetcode.com/problems/last-stone-weight-ii/) | Medium | ⬜ [`05-last-stone-weight-ii.go`](./2-D-dynamic-programming/05-last-stone-weight-ii.go) |
| 309 | [Best Time to Buy and Sell Stock with Cooldown](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/) | Medium | ⬜ [`06-best-time-to-buy-and-sell-stock-with-cooldown.go`](./2-D-dynamic-programming/06-best-time-to-buy-and-sell-stock-with-cooldown.go) |
| 518 | [Coin Change II](https://leetcode.com/problems/coin-change-ii/) | Medium | ⬜ [`07-coin-change-ii.go`](./2-D-dynamic-programming/07-coin-change-ii.go) |
| 494 | [Target Sum](https://leetcode.com/problems/target-sum/) | Medium | ⬜ [`08-target-sum.go`](./2-D-dynamic-programming/08-target-sum.go) |
| 97 | [Interleaving String](https://leetcode.com/problems/interleaving-string/) | Medium | ⬜ [`09-interleaving-string.go`](./2-D-dynamic-programming/09-interleaving-string.go) |
| 877 | [Stone Game](https://leetcode.com/problems/stone-game/) | Medium | ⬜ [`10-stone-game.go`](./2-D-dynamic-programming/10-stone-game.go) |
| 1140 | [Stone Game II](https://leetcode.com/problems/stone-game-ii/) | Medium | ⬜ [`11-stone-game-ii.go`](./2-D-dynamic-programming/11-stone-game-ii.go) |
| 329 | [Longest Increasing Path in a Matrix](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) | Hard | ⬜ [`12-longest-increasing-path-in-a-matrix.go`](./2-D-dynamic-programming/12-longest-increasing-path-in-a-matrix.go) |
| 115 | [Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/) | Hard | ⬜ [`13-distinct-subsequences.go`](./2-D-dynamic-programming/13-distinct-subsequences.go) |
| 72 | [Edit Distance](https://leetcode.com/problems/edit-distance/) | Medium | ⬜ [`14-edit-distance.go`](./2-D-dynamic-programming/14-edit-distance.go) |
| 312 | [Burst Balloons](https://leetcode.com/problems/burst-balloons/) | Hard | ⬜ [`15-burst-balloons.go`](./2-D-dynamic-programming/15-burst-balloons.go) |
| 10 | [Regular Expression Matching](https://leetcode.com/problems/regular-expression-matching/) | Hard | ⬜ [`16-regular-expression-matching.go`](./2-D-dynamic-programming/16-regular-expression-matching.go) |

## Greedy

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 860 | [Lemonade Change](https://leetcode.com/problems/lemonade-change/) | Easy | ⬜ [`01-lemonade-change.go`](./Greedy/01-lemonade-change.go) |
| 53 | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) | Medium | ⬜ [`02-maximum-subarray.go`](./Greedy/02-maximum-subarray.go) |
| 918 | [Maximum Sum Circular Subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/) | Medium | ⬜ [`03-maximum-sum-circular-subarray.go`](./Greedy/03-maximum-sum-circular-subarray.go) |
| 978 | [Longest Turbulent Subarray](https://leetcode.com/problems/longest-turbulent-subarray/) | Medium | ⬜ [`04-longest-turbulent-subarray.go`](./Greedy/04-longest-turbulent-subarray.go) |
| 55 | [Jump Game](https://leetcode.com/problems/jump-game/) | Medium | ⬜ [`05-jump-game.go`](./Greedy/05-jump-game.go) |
| 45 | [Jump Game II](https://leetcode.com/problems/jump-game-ii/) | Medium | ⬜ [`06-jump-game-ii.go`](./Greedy/06-jump-game-ii.go) |
| 1871 | [Jump Game VII](https://leetcode.com/problems/jump-game-vii/) | Medium | ⬜ [`07-jump-game-vii.go`](./Greedy/07-jump-game-vii.go) |
| 134 | [Gas Station](https://leetcode.com/problems/gas-station/) | Medium | ⬜ [`08-gas-station.go`](./Greedy/08-gas-station.go) |
| 846 | [Hand of Straights](https://leetcode.com/problems/hand-of-straights/) | Medium | ⬜ [`09-hand-of-straights.go`](./Greedy/09-hand-of-straights.go) |
| 649 | [Dota2 Senate](https://leetcode.com/problems/dota2-senate/) | Medium | ⬜ [`10-dota2-senate.go`](./Greedy/10-dota2-senate.go) |
| 1899 | [Merge Triplets to Form Target Triplet](https://leetcode.com/problems/merge-triplets-to-form-target-triplet/) | Medium | ⬜ [`11-merge-triplets-to-form-target-triplet.go`](./Greedy/11-merge-triplets-to-form-target-triplet.go) |
| 763 | [Partition Labels](https://leetcode.com/problems/partition-labels/) | Medium | ⬜ [`12-partition-labels.go`](./Greedy/12-partition-labels.go) |
| 678 | [Valid Parenthesis String](https://leetcode.com/problems/valid-parenthesis-string/) | Medium | ⬜ [`13-valid-parenthesis-string.go`](./Greedy/13-valid-parenthesis-string.go) |
| 135 | [Candy](https://leetcode.com/problems/candy/) | Hard | ⬜ [`14-candy.go`](./Greedy/14-candy.go) |

## Intervals

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 57 | [Insert Interval](https://leetcode.com/problems/insert-interval/) | Medium | ⬜ [`01-insert-interval.go`](./Intervals/01-insert-interval.go) |
| 56 | [Merge Intervals](https://leetcode.com/problems/merge-intervals/) | Medium | ⬜ [`02-merge-intervals.go`](./Intervals/02-merge-intervals.go) |
| 435 | [Non-overlapping Intervals](https://leetcode.com/problems/non-overlapping-intervals/) | Medium | ⬜ [`03-non-overlapping-intervals.go`](./Intervals/03-non-overlapping-intervals.go) |
| — | [Meeting Rooms](https://leetcode.com/problems/meeting-rooms/) | Easy | ⬜ [`04-meeting-rooms.go`](./Intervals/04-meeting-rooms.go) |
| — | [Meeting Rooms II](https://leetcode.com/problems/meeting-rooms-ii/) | Medium | ⬜ [`05-meeting-rooms-ii.go`](./Intervals/05-meeting-rooms-ii.go) |
| 2402 | [Meeting Rooms III](https://leetcode.com/problems/meeting-rooms-iii/) | Hard | ⬜ [`06-meeting-rooms-iii.go`](./Intervals/06-meeting-rooms-iii.go) |
| 1851 | [Minimum Interval to Include Each Query](https://leetcode.com/problems/minimum-interval-to-include-each-query/) | Hard | ⬜ [`07-minimum-interval-to-include-each-query.go`](./Intervals/07-minimum-interval-to-include-each-query.go) |

## Math & Geometry

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 168 | [Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/) | Easy | ⬜ [`01-excel-sheet-column-title.go`](./Math-geometry/01-excel-sheet-column-title.go) |
| 1071 | [Greatest Common Divisor of Strings](https://leetcode.com/problems/greatest-common-divisor-of-strings/) | Easy | ⬜ [`02-greatest-common-divisor-of-strings.go`](./Math-geometry/02-greatest-common-divisor-of-strings.go) |
| 2807 | [Insert Greatest Common Divisors in Linked List](https://leetcode.com/problems/insert-greatest-common-divisors-in-linked-list/) | Medium | ⬜ [`03-insert-greatest-common-divisors-in-linked-list.go`](./Math-geometry/03-insert-greatest-common-divisors-in-linked-list.go) |
| 867 | [Transpose Matrix](https://leetcode.com/problems/transpose-matrix/) | Easy | ⬜ [`04-transpose-matrix.go`](./Math-geometry/04-transpose-matrix.go) |
| 48 | [Rotate Image](https://leetcode.com/problems/rotate-image/) | Medium | ⬜ [`05-rotate-image.go`](./Math-geometry/05-rotate-image.go) |
| 54 | [Spiral Matrix](https://leetcode.com/problems/spiral-matrix/) | Medium | ⬜ [`06-spiral-matrix.go`](./Math-geometry/06-spiral-matrix.go) |
| 73 | [Set Matrix Zeroes](https://leetcode.com/problems/set-matrix-zeroes/) | Medium | ⬜ [`07-set-matrix-zeroes.go`](./Math-geometry/07-set-matrix-zeroes.go) |
| 202 | [Happy Number](https://leetcode.com/problems/happy-number/) | Easy | ⬜ [`08-happy-number.go`](./Math-geometry/08-happy-number.go) |
| 66 | [Plus One](https://leetcode.com/problems/plus-one/) | Easy | ⬜ [`09-plus-one.go`](./Math-geometry/09-plus-one.go) |
| 13 | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) | Easy | ⬜ [`10-roman-to-integer.go`](./Math-geometry/10-roman-to-integer.go) |
| 50 | [Pow(x, n)](https://leetcode.com/problems/powx-n/) | Medium | ⬜ [`11-powx-n.go`](./Math-geometry/11-powx-n.go) |
| 43 | [Multiply Strings](https://leetcode.com/problems/multiply-strings/) | Medium | ⬜ [`12-multiply-strings.go`](./Math-geometry/12-multiply-strings.go) |
| 2013 | [Detect Squares](https://leetcode.com/problems/detect-squares/) | Medium | ⬜ [`13-detect-squares.go`](./Math-geometry/13-detect-squares.go) |

## Bit Manipulation

| # | Problem | Difficulty | Solution |
| --- | --- | --- | --- |
| 136 | [Single Number](https://leetcode.com/problems/single-number/) | Easy | ⬜ [`01-single-number.go`](./Bit-manipulation/01-single-number.go) |
| 191 | [Number of 1 Bits](https://leetcode.com/problems/number-of-1-bits/) | Easy | ⬜ [`02-number-of-1-bits.go`](./Bit-manipulation/02-number-of-1-bits.go) |
| 338 | [Counting Bits](https://leetcode.com/problems/counting-bits/) | Easy | ⬜ [`03-counting-bits.go`](./Bit-manipulation/03-counting-bits.go) |
| 67 | [Add Binary](https://leetcode.com/problems/add-binary/) | Easy | ⬜ [`04-add-binary.go`](./Bit-manipulation/04-add-binary.go) |
| 190 | [Reverse Bits](https://leetcode.com/problems/reverse-bits/) | Easy | ⬜ [`05-reverse-bits.go`](./Bit-manipulation/05-reverse-bits.go) |
| 268 | [Missing Number](https://leetcode.com/problems/missing-number/) | Easy | ⬜ [`06-missing-number.go`](./Bit-manipulation/06-missing-number.go) |
| 371 | [Sum of Two Integers](https://leetcode.com/problems/sum-of-two-integers/) | Medium | ⬜ [`07-sum-of-two-integers.go`](./Bit-manipulation/07-sum-of-two-integers.go) |
| 7 | [Reverse Integer](https://leetcode.com/problems/reverse-integer/) | Medium | ⬜ [`08-reverse-integer.go`](./Bit-manipulation/08-reverse-integer.go) |
| 201 | [Bitwise AND of Numbers Range](https://leetcode.com/problems/bitwise-and-of-numbers-range/) | Medium | ⬜ [`09-bitwise-and-of-numbers-range.go`](./Bit-manipulation/09-bitwise-and-of-numbers-range.go) |
| 3133 | [Minimum Array End](https://leetcode.com/problems/minimum-array-end/) | Medium | ⬜ [`10-minimum-array-end.go`](./Bit-manipulation/10-minimum-array-end.go) |
