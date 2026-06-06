# Trie Training: Get Started

A comprehensive, hands-on guide to mastering the Trie data structure in Go.

## Prerequisites

- Basic Go knowledge (you just learned this!)
- Understanding of pointers and recursion
- ~3-4 hours for all exercises

## What You'll Learn

### Core Concepts
- ✅ Trie structure and terminology
- ✅ Insert, search, and delete operations
- ✅ Prefix searching (Trie's superpower!)
- ✅ Wildcard matching
- ✅ Performance optimization

### Real Applications
- ✅ Autocomplete systems
- ✅ Spell checkers
- ✅ Word games (Boggle, Wordle)
- ✅ IP routing
- ✅ Dictionary implementations

## Learning Path

### Phase 1: Fundamentals (1.5 hours)

**Exercise 1:** [01_trie_basics.go](exercises/01_trie_basics.go)
- Understand the structure
- Implement basic insert and search
- See how nodes are connected
- Time: 20 minutes

```bash
go run exercises/01_trie_basics.go
```

**Exercise 2:** [02_trie_prefix_search.go](exercises/02_trie_prefix_search.go)
- Find all words with a prefix
- Implement wildcard matching
- Add frequency-based ranking
- Time: 25 minutes

```bash
go run exercises/02_trie_prefix_search.go
```

**Exercise 3:** [03_trie_delete.go](exercises/03_trie_delete.go)
- Delete words from Trie
- Clean up unused nodes
- Handle shared prefixes
- Time: 20 minutes (most complex!)

```bash
go run exercises/03_trie_delete.go
```

### Phase 2: Applications (1.5 hours)

**Exercise 4:** [04_autocomplete_system.go](exercises/04_autocomplete_system.go)
- Build a real autocomplete system
- Rank suggestions by frequency
- Handle case insensitivity
- Time: 25 minutes

```bash
go run exercises/04_autocomplete_system.go
```

**Exercise 5:** [05_performance_analysis.go](exercises/05_performance_analysis.go)
- Benchmark Trie vs Hash Table
- Understand complexity trade-offs
- See where Trie excels
- Time: 20 minutes

```bash
go run exercises/05_performance_analysis.go
```

## Quick Start

```bash
# Navigate to exercises
cd trie/exercises

# Run the first exercise
go run 01_trie_basics.go

# Read the code
code 01_trie_basics.go  # VS Code

# Try the exercises at the bottom of each file
# Move to the next one
go run 02_trie_prefix_search.go
```

## How to Learn Each Exercise

### 1. Read Comments
Every file has detailed explanations. Read them carefully.

### 2. Understand the Code
Look at the working examples before the exercises.

### 3. Run It
```bash
go run XX_name.go
```

### 4. Modify & Experiment
Change values, test edge cases, break things on purpose.

### 5. Complete Exercises
Each file ends with an exercise section. Implement what it asks.

### 6. Move On
Once confident, go to the next file.

## Key Concepts Overview

### The Structure
```
       root
       /  \
      h    c
      |    |
      e    a
      |    |
      l    t
      |
      l
      |
      o
```
- Each node represents a character
- `isEnd` marks the end of a word
- `children` map points to next characters

### Operations

**Insert:** O(L) where L = word length
```
Insert "hello":
1. Go to root
2. Create/navigate h -> e -> l -> l -> o
3. Mark 'o' as end of word
```

**Search:** O(L)
```
Search "hello":
1. Follow path h -> e -> l -> l -> o
2. Check if 'o' is marked as end
```

**Prefix Search:** O(P + K) where P = prefix length, K = results
```
Find words starting with "hel":
1. Navigate to 'l' in h -> e -> l
2. Collect all words from that node
3. Return "hello", "help", etc.
```

**Delete:** O(L) with cleanup
```
Delete "hello":
1. Navigate to 'o'
2. Unmark as end of word
3. Clean up empty nodes
```

## Common Mistakes

### ❌ Forgetting to Mark Word End
```go
// Wrong: forgot node.isEnd = true
for _, ch := range word {
    node = node.children[ch]
}
// node.isEnd = true  <- MISSING!
```

### ❌ Not Handling Non-existent Paths
```go
// Wrong: panics if path doesn't exist
node = node.children[ch]

// Right: check first
if child, exists := node.children[ch]; exists {
    node = child
} else {
    return false
}
```

### ❌ Inefficient Prefix Search
```go
// Wrong: checks every character of every word
for word := range allWords {
    if strings.HasPrefix(word, prefix) { count++ }
}

// Right: navigate to prefix node, then collect
node := navigateToPrefix(prefix)
collectWords(node)
```

## Testing Your Implementation

### Test Cases to Try

**Basic:**
```
Insert: "cat", "car"
Search: "cat" -> true, "ca" -> false, "car" -> true
```

**Prefix:**
```
Insert: "apple", "app", "apply"
FindPrefix("app") -> ["app", "apple", "apply"]
```

**Delete:**
```
Insert: "cat", "car"
Delete: "cat"
Search: "cat" -> false, "car" -> true
```

**Edge Cases:**
- Empty string
- Single character
- Very long words
- Unicode characters
- Overlapping words

## Resources

### Key References
- [QUICK_REFERENCE.md](QUICK_REFERENCE.md) — Syntax and patterns
- [README.md](README.md) — Conceptual overview
- https://www.cs.usfca.edu/~galles/visualization/Trie.html — Visual Trie builder

### Practice
- [LeetCode #208](https://leetcode.com/problems/implement-trie-prefix-tree/) — Implement Trie
- [LeetCode #211](https://leetcode.com/problems/add-and-search-word-data-structure-design/) — Wildcard matching
- [LeetCode #642](https://leetcode.com/problems/design-search-autocomplete-system/) — Autocomplete

### Interview Prep
- Know the complexity of each operation
- Be able to explain when Trie is better than hash table
- Implement delete (it's tricky!)
- Discuss variations (TST, Patricia Tree)

## After Exercises

### Level 1: Consolidate
1. Rewrite Trie from scratch (no copying)
2. Test with different datasets
3. Add error handling

### Level 2: Extend
1. Implement deletion cleanup properly
2. Add case-insensitive search
3. Add wildcard matching

### Level 3: Optimize
1. Use array-based children for English
2. Implement Patricia tree compression
3. Add frequency tracking

### Level 4: Real Project
1. Build autocomplete CLI
2. Implement spell checker
3. Solve LeetCode problems

## Time Estimates

```
Understanding Trie structure:          10 min
Implementing insert/search:            15 min
Prefix search:                         20 min
Deletion:                              25 min
Building autocomplete:                 20 min
Performance analysis:                  20 min
Bonus: variations and optimization:    30-60 min
─────────────────────────────────────────────
TOTAL:                                 3-4 hours
```

## Debugging Tips

### Print the Trie Structure
```go
func (t *Trie) Print() {
    t.printHelper(t.root, "", "")
}
```

### Check Intermediate States
```go
fmt.Printf("At node for 'cat': isEnd=%v, children=%v\n", node.isEnd, len(node.children))
```

### Trace Execution
```go
for _, ch := range word {
    fmt.Printf("Processing '%c'\n", ch)
    if child, exists := node.children[ch]; exists {
        fmt.Printf("  Found child\n")
        node = child
    } else {
        fmt.Printf("  No child, returning\n")
        return false
    }
}
```

## Common Questions

**Q: Should children be a map or array?**
A: Map for flexibility, array for speed. For English, array[26] is optimal.

**Q: How do I handle case sensitivity?**
A: Convert all words to lowercase on insert and search.

**Q: How do I support Unicode?**
A: Use `map[rune]*TrieNode` instead of array.

**Q: When is Trie better than hash table?**
A: When you need prefix search. Hash table for exact lookup.

**Q: How do I optimize for space?**
A: Use Patricia tree or compress common branches.

## Next Steps

1. ✅ Start with Exercise 1: `go run exercises/01_trie_basics.go`
2. ✅ Complete all exercises in order
3. ✅ Rewrite from scratch (no reference)
4. ✅ Solve LeetCode problems
5. ✅ Build a real project (autocomplete CLI, spell checker)

---

**Ready?** Run your first exercise:

```bash
cd exercises
go run 01_trie_basics.go
```

Happy learning! 🌳

