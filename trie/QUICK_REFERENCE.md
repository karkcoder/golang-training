# Trie Quick Reference

## Basic Structure

```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}
```

## Core Operations

### Insert
```go
func (t *Trie) Insert(word string) {
    node := t.root
    for _, ch := range word {
        if _, exists := node.children[ch]; !exists {
            node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        node = node.children[ch]
    }
    node.isEnd = true
}
```
**Time:** O(L) where L = word length  
**Space:** O(L) new nodes created

### Search (Exact Match)
```go
func (t *Trie) Search(word string) bool {
    node := t.root
    for _, ch := range word {
        if child, exists := node.children[ch]; exists {
            node = child
        } else {
            return false
        }
    }
    return node.isEnd
}
```
**Time:** O(L)  
**Space:** O(1)

### Prefix Search
```go
func (t *Trie) FindWordsWithPrefix(prefix string) []string {
    node := t.root
    var words []string
    
    for _, ch := range prefix {
        if child, exists := node.children[ch]; exists {
            node = child
        } else {
            return words
        }
    }
    
    t.collectWords(node, prefix, &words)
    return words
}
```
**Time:** O(P + K) where P = prefix length, K = matching results  
**Space:** O(K) for results

### Delete
```go
func (t *Trie) Delete(word string) bool {
    return t.deleteHelper(t.root, word, 0)
}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
    if index == len(word) {
        if !node.isEnd {
            return false
        }
        node.isEnd = false
        return len(node.children) == 0
    }
    
    ch := rune(word[index])
    if child, exists := node.children[ch]; !exists {
        return false
    } else {
        shouldDelete := t.deleteHelper(child, word, index+1)
        if shouldDelete {
            delete(node.children, ch)
        }
        return !node.isEnd && len(node.children) == 0
    }
}
```
**Time:** O(L)  
**Space:** O(L) recursion stack

## Complexity Summary

| Operation | Time | Space | Notes |
|-----------|------|-------|-------|
| Insert | O(L) | O(L) | L = word length |
| Search | O(L) | O(1) | Exact match only |
| Delete | O(L) | O(L) | With cleanup |
| Prefix Search | O(P+K) | O(K) | K = results, P = prefix |
| Wildcard Search | O(N*L) | O(L) | N = alphabet size |

## Comparison: Trie vs Hash Table vs Sorted Array

### Exact Word Search
```
HashSet:        O(L)
Trie:           O(L)
Sorted Array:   O(log N + L) for comparison
Winner:         HashSet (simpler, less overhead)
```

### Find All Words with Prefix
```
HashSet:        O(N*L)  - must check every word!
Trie:           O(P+K)  - navigate prefix, then collect
Sorted Array:   O(log N + K) - binary search, then scan
Winner:         Trie (unless K > array size)
```

### Space Usage
```
HashSet:        O(N*L)
Trie:           O(ALPHABET * depth * nodes) 
                ≈ O(N*L) worst, but shares prefixes
Sorted Array:   O(N*L)
Winner:         Trie for English (shares "the", "and", etc.)
```

## Use Cases

### Perfect for Trie:
- ✅ Autocomplete systems
- ✅ IP routing (longest prefix match)
- ✅ Spell checking with suggestions
- ✅ Boggle, crossword solvers
- ✅ Dictionary with prefix operations

### Better Use Hash Table:
- ✅ Just need exact lookups
- ✅ Very long words (high overhead)
- ✅ No prefix operations needed

## Implementation Variations

### 1. Array-Based (Fixed Alphabet)
```go
type TrieNode struct {
    children [26]*TrieNode  // For 'a'-'z'
    isEnd    bool
}
```
**Pro:** O(1) child lookup  
**Con:** Wasted space for sparse alphabets

### 2. Hash Map (Flexible Alphabet)
```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}
```
**Pro:** Works with any Unicode  
**Con:** O(1) average, O(n) worst case lookup

### 3. Ternary Search Tree (TST)
```go
type TSTNode struct {
    ch    rune
    lo    *TSTNode  // Characters < ch
    mid   *TSTNode  // Characters == ch
    hi    *TSTNode  // Characters > ch
    isEnd bool
}
```
**Pro:** Binary search within tree, less space  
**Con:** More complex, slightly slower

### 4. Patricia Tree (Compressed Trie)
Combines edges with common children into single nodes.
**Pro:** Uses ~50% less space  
**Con:** More complex implementation

## Common Patterns

### Word Frequency Ranking
```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
    freq     int  // Add this
}

// After collecting words, sort by freq
sort.Slice(words, func(i, j int) bool {
    return words[i].freq > words[j].freq
})
```

### Wildcard Matching (. = any char)
```go
func (t *Trie) searchWildcard(node *TrieNode, pattern string, idx int) bool {
    if idx == len(pattern) {
        return node.isEnd
    }
    
    if pattern[idx] == '.' {
        for _, child := range node.children {
            if t.searchWildcard(child, pattern, idx+1) {
                return true
            }
        }
    } else {
        if child, ok := node.children[rune(pattern[idx])]; ok {
            return t.searchWildcard(child, pattern, idx+1)
        }
    }
    return false
}
```

### Edit Distance (Spellchecker)
```go
// Find all words within distance K from query
// Uses DFS with distance tracking
func (t *Trie) findWordsWithDistance(word string, maxDist int) []string {
    var results []string
    t.searchWithDistance(t.root, word, 0, 0, maxDist, "", &results)
    return results
}
```

## Interview Tips

1. **Always clarify:** Exact search or prefix search?
2. **Space matters:** Consider alphabet size and word length
3. **Edge cases:** Empty strings, special characters, case sensitivity
4. **Optimization:** Consider caching frequent prefixes
5. **Variants:** Know TST and Patricia trees by name

## LeetCode Problems

- **208:** Implement Trie (Prefix Tree)
- **211:** Add and Search Word
- **212:** Word Search II
- **642:** Design Search Autocomplete System
- **472:** Concatenated Words
- **588:** Design In-Memory File System

## Performance Tips

1. Use array-based Trie for English letters (26 children)
2. Use map-based for Unicode or sparse alphabets
3. Cache frequently accessed nodes
4. For very large tries, use Patricia tree
5. Pre-build and serialize for repeated use
