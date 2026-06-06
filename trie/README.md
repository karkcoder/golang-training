# Trie Data Structure Training

A comprehensive guide to understanding and implementing **Tries** (prefix trees) with hands-on exercises in Go.

## What is a Trie?

A **Trie** (pronounced "try" or "tree") is a tree-like data structure that stores strings efficiently. It's particularly useful for:

- **Autocomplete** — "search as you type"
- **Spell checking** — find words with common prefixes
- **IP routing** — longest prefix matching
- **Dictionary lookups** — efficiently store and search words
- **Word games** — Boggle, Scrabble word validation

### Visual Example

```
         root
        /    \
       h      c
       |      |
       e      a
       |      |
       l      t
       |
       l
       |
       o
```

This Trie stores: "hello", "cat"

### Why Use a Trie?

| Operation | Hash Table | Trie |
|-----------|-----------|------|
| Insert | O(n) | O(n) |
| Search | O(n) | O(n) |
| Delete | O(n) | O(n) |
| Prefix search | O(m*n) | O(m) |
| Space | O(n) | O(n*k) |

Where:
- n = length of word
- m = number of words with prefix
- k = alphabet size (26 for lowercase English)

**Tries shine at prefix operations**, which hash tables cannot do efficiently.

## Learning Path

### Phase 1: Fundamentals (1 hour)
- **Lesson 1** — Trie structure and terminology
- **Lesson 2** — Building a basic Trie
- **Lesson 3** — Insert and search operations

### Phase 2: Core Operations (1 hour)
- **Lesson 4** — Prefix search and wildcard matching
- **Lesson 5** — Delete operations (complex!)
- **Lesson 6** — Space optimization

### Phase 3: Applications (1-2 hours)
- **Lesson 7** — Autocomplete system
- **Lesson 8** — Spell checker
- **Lesson 9** — Word games and puzzles

### Phase 4: Optimization (30 min)
- **Lesson 10** — Performance analysis
- **Lesson 11** — Advanced Trie variants (Ternary, Suffix)

## Exercises

All exercises are in the `exercises/` directory:

```
exercises/
├── 01_trie_basics.go         # Understand the structure
├── 02_trie_insert.go         # Implement insert
├── 03_trie_search.go         # Implement search
├── 04_prefix_search.go       # Find all words with prefix
├── 05_trie_delete.go         # Remove words (hard!)
├── 06_autocomplete.go        # Build autocomplete system
├── 07_spell_checker.go       # Implement spell checker
├── 08_word_games.go          # Word validation
└── 09_performance.go         # Benchmark vs hash tables
```

**Total time:** 3-4 hours hands-on

## Quick Start

```bash
# Start with basics
go run exercises/01_trie_basics.go

# Then progress through each exercise
go run exercises/02_trie_insert.go
```

## Key Concepts to Learn

### 1. Node Structure
```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool                // Marks end of a word
}
```

### 2. Insertion
Add characters one by one, creating nodes as needed.

### 3. Search
Follow the path. If all characters exist and `isEnd` is true, word exists.

### 4. Prefix Search
Follow the prefix path, then collect all words from that node.

### 5. Deletion
Remove word, then cleanup empty nodes (backtracking).

## Real-World Examples

### Autocomplete
- **Google Search** — billions of queries, needs sub-millisecond response
- **IDE autocomplete** — suggest variable names, function names
- **Mobile keyboards** — predict next word

### Spell Checking
- **Microsoft Word** — red squiggly underlines
- **Browser spellcheck** — context-aware suggestions

### IP Routing
- **Network routers** — longest prefix match for routing tables
- **Firewalls** — IP blocking/whitelisting

### Spelling Bee / Wordle
- **Word validation** — is this a valid English word?
- **Dictionary lookups** — return all valid words

## Performance Characteristics

### Space
```
Dictionary of N words, average length L:
- Hash table: O(N*L) characters stored
- Trie: O(N*L) worst case, but shares common prefixes
```

For English words, Trie is often smaller because it shares prefixes.

### Time
```
Insert/Search/Delete: O(L) where L = word length
Prefix search: O(P + M) where:
  P = length of prefix
  M = number of words with that prefix (collecting results)
```

## Comparison: Trie vs Hash Table vs Sorted Array

### Use Case: "Given a word list, find all words starting with 'pre'"

```
Hash Table: O(N) — must check every word
Sorted Array: O(log N + K) — binary search, then scan
Trie: O(P + K) — P = prefix length, K = results (optimal!)
```

## Resources

- **Visualization** — https://www.cs.usfca.edu/~galles/visualization/Trie.html
- **LeetCode** — Search "Trie" for practice problems
- **Interview Questions** — Common in FAANG interviews

## Exercises Breakdown

| # | Exercise | Focus | Difficulty |
|---|----------|-------|------------|
| 1 | basics | Structure, terminology | ⭐ Easy |
| 2 | insert | Adding words | ⭐ Easy |
| 3 | search | Finding words | ⭐ Easy |
| 4 | prefix_search | All words with prefix | ⭐⭐ Medium |
| 5 | delete | Removing words | ⭐⭐⭐ Hard |
| 6 | autocomplete | Ranking suggestions | ⭐⭐⭐ Hard |
| 7 | spell_checker | Edit distance | ⭐⭐⭐ Hard |
| 8 | word_games | Validation | ⭐⭐ Medium |
| 9 | performance | Benchmarking | ⭐⭐ Medium |

## Next Steps After Training

1. **Implement a complete Trie package** — Make it reusable
2. **Solve LeetCode problems** — Cement your understanding
3. **Build a real project** — Autocomplete for your CLI tool
4. **Interview prep** — Common in system design and coding rounds

## Common Interview Questions

1. Implement a Trie (insert, search, startsWith)
2. Word search in a 2D grid using Trie
3. Autocomplete system design
4. Spell checker with edit distance
5. Longest word in dictionary
6. Alien dictionary sort order

All of these are covered in our exercises!

---

**Ready?** Start with [Exercise 1](exercises/01_trie_basics.go).

