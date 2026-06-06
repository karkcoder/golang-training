package main

import (
	"fmt"
	"math/rand"
	"time"
)

// === TRIE PERFORMANCE ANALYSIS ===
// Compare Trie vs Hash Set for different operations

// --- HASH SET APPROACH ---
func buildHashSet(words []string) map[string]bool {
	set := make(map[string]bool)
	for _, word := range words {
		set[word] = true
	}
	return set
}

func searchHashSet(set map[string]bool, word string) bool {
	return set[word]
}

// --- TRIE APPROACH ---
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

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

// Prefix search (Trie advantage!)
func (t *Trie) FindWordsWithPrefix(prefix string) int {
	node := t.root
	count := 0

	for _, ch := range prefix {
		if child, exists := node.children[ch]; exists {
			node = child
		} else {
			return 0
		}
	}

	// Count all words from this node
	t.countWordsFromNode(node, &count)
	return count
}

func (t *Trie) countWordsFromNode(node *TrieNode, count *int) {
	if node.isEnd {
		*count++
	}
	for _, child := range node.children {
		t.countWordsFromNode(child, count)
	}
}

// --- BENCHMARKING ---

func generateRandomWords(count int, maxLen int) []string {
	words := make([]string, count)
	chars := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < count; i++ {
		len := rand.Intn(maxLen) + 1
		word := ""
		for j := 0; j < len; j++ {
			word += string(chars[rand.Intn(len(chars))])
		}
		words[i] = word
	}

	return words
}

func benchmarkExactSearch(name string, words []string, searchTerms []string) {
	start := time.Now()

	if name == "HashSet" {
		set := buildHashSet(words)
		for _, term := range searchTerms {
			searchHashSet(set, term)
		}
	} else {
		trie := NewTrie()
		for _, word := range words {
			trie.Insert(word)
		}
		for _, term := range searchTerms {
			trie.Search(term)
		}
	}

	elapsed := time.Since(start)
	fmt.Printf("  %s: %v\n", name, elapsed)
}

func benchmarkPrefixSearch(name string, words []string, prefixes []string) {
	if name == "HashSet" {
		start := time.Now()
		set := buildHashSet(words)
		for _, prefix := range prefixes {
			count := 0
			for word := range set {
				if len(word) >= len(prefix) && word[:len(prefix)] == prefix {
					count++
				}
			}
		}
		elapsed := time.Since(start)
		fmt.Printf("  %s: %v\n", name, elapsed)
	} else {
		start := time.Now()
		trie := NewTrie()
		for _, word := range words {
			trie.Insert(word)
		}
		for _, prefix := range prefixes {
			trie.FindWordsWithPrefix(prefix)
		}
		elapsed := time.Since(start)
		fmt.Printf("  %s: %v\n", name, elapsed)
	}
}

func main() {
	fmt.Println("=== TRIE PERFORMANCE ANALYSIS ===\n")

	// Generate test data
	wordCount := 10000
	maxWordLen := 20
	words := generateRandomWords(wordCount, maxWordLen)

	fmt.Printf("Test data: %d random words, max length %d\n\n", wordCount, maxWordLen)

	// --- TEST 1: BUILD TIME ---
	fmt.Println("1. BUILD TIME (creating structure):")

	start := time.Now()
	buildHashSet(words)
	hashSetBuildTime := time.Since(start)
	fmt.Printf("  HashSet: %v\n", hashSetBuildTime)

	start = time.Now()
	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}
	trieBuildTime := time.Since(start)
	fmt.Printf("  Trie: %v\n", trieBuildTime)

	// --- TEST 2: EXACT SEARCH ---
	fmt.Println("\n2. EXACT SEARCH (searching for existing words):")
	searchTerms := words[:100]
	benchmarkExactSearch("HashSet", words, searchTerms)
	benchmarkExactSearch("Trie", words, searchTerms)

	// --- TEST 3: PREFIX SEARCH (Trie shines here!) ---
	fmt.Println("\n3. PREFIX SEARCH (finding all words with prefix 'a'):")
	prefixes := []string{"a", "ab", "abc", "abcd"}
	benchmarkPrefixSearch("HashSet", words, prefixes)
	benchmarkPrefixSearch("Trie", words, prefixes)

	// --- ANALYSIS ---
	fmt.Println("\n=== COMPLEXITY COMPARISON ===")

	fmt.Println("\nExact Search:")
	fmt.Println("  HashSet: O(L) - hash the word, compare")
	fmt.Println("  Trie:    O(L) - traverse L characters")
	fmt.Println("  Same complexity, but Trie has constant overhead")

	fmt.Println("\nPrefix Search:")
	fmt.Println("  HashSet: O(N*L) - must check every word!")
	fmt.Println("  Trie:    O(P+K) - P=prefix len, K=results")
	fmt.Println("  Trie is MUCH faster for prefix search!")

	fmt.Println("\nSpace Usage:")
	fmt.Println("  HashSet: O(N*L) - store full strings")
	fmt.Println("  Trie:    O(N*L) worst, but shares prefixes")
	fmt.Println("  For English words, Trie uses ~40% less space")

	// --- PRACTICAL IMPLICATIONS ---
	fmt.Println("\n=== WHEN TO USE WHAT ===")

	fmt.Println("\nUse HashSet when:")
	fmt.Println("  - Only doing exact lookups")
	fmt.Println("  - Memory is very constrained")
	fmt.Println("  - Words are very long (Trie overhead increases)")
	fmt.Println("  - No prefix searching needed")

	fmt.Println("\nUse Trie when:")
	fmt.Println("  - Need prefix search (autocomplete, IP routing)")
	fmt.Println("  - Many queries share common prefixes")
	fmt.Println("  - Need to find all words with prefix")
	fmt.Println("  - Want to support wildcards")

	fmt.Println("\n=== REAL-WORLD SCALE ===")
	fmt.Println("Google Search:")
	fmt.Println("  - ~3.5 billion searches/day")
	fmt.Println("  - Must handle in < 100ms")
	fmt.Println("  - Prefix search is essential")
	fmt.Println("  - Uses optimized Trie variants with caching")

	fmt.Println("\nIDEs (VS Code, IntelliJ):")
	fmt.Println("  - Autocomplete on every keystroke")
	fmt.Println("  - Instant suggestions")
	fmt.Println("  - Uses Trie for symbol names")

	fmt.Println("\nIP Routing:")
	fmt.Println("  - Routers need longest prefix match")
	fmt.Println("  - Trie enables O(L) matching vs O(N) linear search")
	fmt.Println("  - Critical for network performance")
}

// To run:
// go run 05_performance_analysis.go

// === KEY TAKEAWAY ===
// Trie is not always faster, but it excels at prefix operations.
// Choose your data structure based on your access patterns!
