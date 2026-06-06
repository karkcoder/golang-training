package main

import (
	"fmt"
	"sort"
)

// === TRIE WITH PREFIX SEARCH ===
// This is where Tries shine! Finding all words with a given prefix
// in O(P + K) time where P is prefix length and K is result count.

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

// === PREFIX SEARCH (NEW!) ===
// Find all words that start with a given prefix

// findWordsHelper is a DFS that collects all words from a given node
func (t *Trie) findWordsHelper(node *TrieNode, prefix string, words *[]string) {
	// If this node marks the end of a word, add the full word
	if node.isEnd {
		*words = append(*words, prefix)
	}

	// Recursively visit all children
	for ch, child := range node.children {
		t.findWordsHelper(child, prefix+string(ch), words)
	}
}

// FindWordsWithPrefix returns all words that start with the given prefix
func (t *Trie) FindWordsWithPrefix(prefix string) []string {
	node := t.root
	var words []string

	// First, navigate to the prefix node
	for _, ch := range prefix {
		if child, exists := node.children[ch]; exists {
			node = child
		} else {
			// Prefix not found, return empty
			return words
		}
	}

	// Now collect all words from this node
	t.findWordsHelper(node, prefix, &words)

	// Return sorted for consistent results
	sort.Strings(words)
	return words
}

// === WILDCARD SEARCH ===
// Support wildcards where '.' matches any character

// searchWildcardHelper does DFS with wildcard support
func (t *Trie) searchWildcardHelper(node *TrieNode, pattern string, index int) bool {
	// Reached end of pattern and end of a word
	if index == len(pattern) {
		return node.isEnd
	}

	ch := rune(pattern[index])

	if ch == '.' {
		// Wildcard: try all children
		for _, child := range node.children {
			if t.searchWildcardHelper(child, pattern, index+1) {
				return true
			}
		}
		return false
	} else {
		// Regular character: must match
		if child, exists := node.children[ch]; exists {
			return t.searchWildcardHelper(child, pattern, index+1)
		}
		return false
	}
}

// SearchWithWildcard searches for a pattern with '.' as wildcard (matches any char)
func (t *Trie) SearchWithWildcard(pattern string) bool {
	return t.searchWildcardHelper(t.root, pattern, 0)
}

// === AUTOCOMPLETE RANKING ===
// Track frequency to rank suggestions

type TrieNodeWithFreq struct {
	children map[rune]*TrieNodeWithFreq
	isEnd    bool
	freq     int // How many times this word was inserted
}

func NewTrieNodeWithFreq() *TrieNodeWithFreq {
	return &TrieNodeWithFreq{
		children: make(map[rune]*TrieNodeWithFreq),
		isEnd:    false,
		freq:     0,
	}
}

type TrieWithFreq struct {
	root *TrieNodeWithFreq
}

func NewTrieWithFreq() *TrieWithFreq {
	return &TrieWithFreq{
		root: NewTrieNodeWithFreq(),
	}
}

func (t *TrieWithFreq) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = NewTrieNodeWithFreq()
		}
		node = node.children[ch]
	}
	node.isEnd = true
	node.freq++ // Increment frequency
}

type SuggestedWord struct {
	Word  string
	Freq  int
}

func (t *TrieWithFreq) findSuggestionsHelper(node *TrieNodeWithFreq, prefix string, suggestions *[]SuggestedWord) {
	if node.isEnd {
		*suggestions = append(*suggestions, SuggestedWord{Word: prefix, Freq: node.freq})
	}
	for ch, child := range node.children {
		t.findSuggestionsHelper(child, prefix+string(ch), suggestions)
	}
}

// GetSuggestions returns words with prefix, ranked by frequency
func (t *TrieWithFreq) GetSuggestions(prefix string) []SuggestedWord {
	node := t.root
	var suggestions []SuggestedWord

	for _, ch := range prefix {
		if child, exists := node.children[ch]; exists {
			node = child
		} else {
			return suggestions
		}
	}

	t.findSuggestionsHelper(node, prefix, &suggestions)

	// Sort by frequency (descending), then alphabetically
	sort.Slice(suggestions, func(i, j int) bool {
		if suggestions[i].Freq != suggestions[j].Freq {
			return suggestions[i].Freq > suggestions[j].Freq
		}
		return suggestions[i].Word < suggestions[j].Word
	})

	return suggestions
}

func main() {
	fmt.Println("=== TRIE PREFIX SEARCH ===\n")

	trie := NewTrie()
	words := []string{"apple", "app", "application", "apply", "apricot", "apt", "banana", "band", "bandana", "bat"}
	for _, word := range words {
		trie.Insert(word)
	}

	fmt.Println("1. Find all words with prefix 'app':")
	results := trie.FindWordsWithPrefix("app")
	for _, word := range results {
		fmt.Printf("  %s\n", word)
	}

	fmt.Println("\n2. Find all words with prefix 'ba':")
	results = trie.FindWordsWithPrefix("ba")
	for _, word := range results {
		fmt.Printf("  %s\n", word)
	}

	fmt.Println("\n3. Prefix that doesn't exist 'xyz':")
	results = trie.FindWordsWithPrefix("xyz")
	fmt.Printf("  Results: %v\n", results)

	fmt.Println("\n=== WILDCARD SEARCH ===")
	fmt.Println("Pattern 'ap.' matches: app (a-p-p), apt (a-p-t)")
	result := trie.SearchWithWildcard("ap.")
	fmt.Printf("  SearchWithWildcard('ap.'): %v\n", result)

	result = trie.SearchWithWildcard("ba...")
	fmt.Printf("  SearchWithWildcard('ba...'): %v (matches banana, bandana, etc)\n", result)

	result = trie.SearchWithWildcard("ba")
	fmt.Printf("  SearchWithWildcard('ba'): %v (must be exact word)\n", result)

	fmt.Println("\n=== FREQUENCY-BASED SUGGESTIONS ===")
	trieFq := NewTrieWithFreq()

	// Insert with different frequencies
	searches := []string{"apple", "apple", "apple", "app", "app", "application", "apricot", "apricot"}
	for _, search := range searches {
		trieFq.Insert(search)
	}

	fmt.Println("Suggestions for 'app' (ranked by frequency):")
	suggestions := trieFq.GetSuggestions("app")
	for i, sugg := range suggestions {
		fmt.Printf("  %d. %s (freq: %d)\n", i+1, sugg.Word, sugg.Freq)
	}

	fmt.Println("\n=== EXERCISE ===")
	fmt.Println("1. Create a Trie with words: 'cat', 'car', 'card', 'care', 'dog', 'dodge'")
	fmt.Println("2. Find all words with prefix 'ca'")
	fmt.Println("3. Find all words with prefix 'do'")
	fmt.Println("4. Use wildcard search:")
	fmt.Println("   - 'ca.' should match 'car' (if you convert it to exact search)")
	fmt.Println("   - '...' should match 'cat' and 'dog' (3-char words)")
	fmt.Println("5. Implement a spellchecker-like function that finds words")
	fmt.Println("   similar to a typo (e.g., 'cot' is close to 'cat')")
}

// To run:
// go run 02_trie_prefix_search.go
