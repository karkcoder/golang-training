package main

import (
	"fmt"
	"sort"
	"strings"
)

// === AUTOCOMPLETE SYSTEM ===
// Real-world application: "search as you type"
// Requirements:
// 1. Quick response (milliseconds)
// 2. Ranked by relevance (frequency, recency)
// 3. Handle typos (optional)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
	freq     int // How many times this word was searched
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
		freq:     0,
	}
}

type AutocompleteSystem struct {
	root        *TrieNode
	maxResults  int // Limit number of suggestions
}

func NewAutocompleteSystem(maxResults int) *AutocompleteSystem {
	return &AutocompleteSystem{
		root:       NewTrieNode(),
		maxResults: maxResults,
	}
}

// AddSearchTerm adds a query with a frequency count
func (as *AutocompleteSystem) AddSearchTerm(term string) {
	term = strings.ToLower(term)
	node := as.root

	for _, ch := range term {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}

	node.isEnd = true
	node.freq++
}

// Suggestion represents a suggested search term
type Suggestion struct {
	Term  string
	Score int
}

// getSuggestionsHelper collects all suggestions from a node
func (as *AutocompleteSystem) getSuggestionsHelper(node *TrieNode, prefix string, suggestions *[]Suggestion) {
	if node.isEnd {
		*suggestions = append(*suggestions, Suggestion{
			Term:  prefix,
			Score: node.freq,
		})
	}

	for ch, child := range node.children {
		as.getSuggestionsHelper(child, prefix+string(ch), suggestions)
	}
}

// GetSuggestions returns top N suggestions for a prefix
func (as *AutocompleteSystem) GetSuggestions(prefix string) []Suggestion {
	prefix = strings.ToLower(prefix)
	node := as.root
	var suggestions []Suggestion

	// Navigate to prefix node
	for _, ch := range prefix {
		if child, exists := node.children[ch]; exists {
			node = child
		} else {
			return suggestions
		}
	}

	// Collect all suggestions
	as.getSuggestionsHelper(node, prefix, &suggestions)

	// Sort by frequency (descending), then alphabetically
	sort.Slice(suggestions, func(i, j int) bool {
		if suggestions[i].Score != suggestions[j].Score {
			return suggestions[i].Score > suggestions[j].Score
		}
		return suggestions[i].Term < suggestions[j].Term
	})

	// Limit results
	if len(suggestions) > as.maxResults {
		suggestions = suggestions[:as.maxResults]
	}

	return suggestions
}

// === SIMULATING REAL SEARCHES ===

type SearchLog struct {
	query string
	count int
}

func main() {
	fmt.Println("=== AUTOCOMPLETE SYSTEM ===\n")

	// Create autocomplete system (max 5 suggestions)
	ac := NewAutocompleteSystem(5)

	// Simulate real search history with frequencies
	searchHistory := []SearchLog{
		{"golang tutorials", 150},
		{"golang concurrency", 120},
		{"golang best practices", 85},
		{"go language", 200},
		{"golang docker", 90},
		{"google search", 1000},
		{"google maps", 800},
		{"google drive", 500},
		{"good restaurants", 200},
		{"good movies", 150},
	}

	fmt.Println("1. Building search history:")
	for _, log := range searchHistory {
		for i := 0; i < log.count; i++ {
			ac.AddSearchTerm(log.query)
		}
		fmt.Printf("  Added '%s' (frequency: %d)\n", log.query, log.count)
	}

	fmt.Println("\n2. Get suggestions as user types 'go':")
	suggestions := ac.GetSuggestions("go")
	for i, sugg := range suggestions {
		fmt.Printf("  %d. '%s' (score: %d)\n", i+1, sugg.Term, sugg.Score)
	}

	fmt.Println("\n3. Get suggestions as user types 'goo':")
	suggestions = ac.GetSuggestions("goo")
	for i, sugg := range suggestions {
		fmt.Printf("  %d. '%s' (score: %d)\n", i+1, sugg.Term, sugg.Score)
	}

	fmt.Println("\n4. Get suggestions for 'google':")
	suggestions = ac.GetSuggestions("google")
	for i, sugg := range suggestions {
		fmt.Printf("  %d. '%s' (score: %d)\n", i+1, sugg.Term, sugg.Score)
	}

	fmt.Println("\n5. Get suggestions for 'golang':")
	suggestions = ac.GetSuggestions("golang")
	for i, sugg := range suggestions {
		fmt.Printf("  %d. '%s' (score: %d)\n", i+1, sugg.Term, sugg.Score)
	}

	fmt.Println("\n6. Non-existent prefix 'xyz':")
	suggestions = ac.GetSuggestions("xyz")
	fmt.Printf("  Suggestions: %v\n", suggestions)

	fmt.Println("\n=== COMPLEXITY ANALYSIS ===")
	fmt.Println("Building the trie:")
	fmt.Println("  - Time: O(N * L) where N = number of searches, L = avg query length")
	fmt.Println("  - Space: O(N * L) for storing all unique prefixes")

	fmt.Println("\nGetting suggestions:")
	fmt.Println("  - Time: O(P + K) where P = prefix length, K = matching results")
	fmt.Println("  - Much faster than O(N * L) for hash table lookup!")
	fmt.Println("  - Prefix search is instant with Trie")

	fmt.Println("\n=== REAL WORLD: GOOGLE SEARCH ===")
	fmt.Println("When you search 'how to build a ':
	fmt.Println("  1. Google's Trie suggests thousands of completions")
	fmt.Println("  2. Top ones ranked by:")
	fmt.Println("     - Frequency (how often people search it)")
	fmt.Println("     - Relevance (personalized by location, history)")
	fmt.Println("     - Freshness (recent trending searches)")
	fmt.Println("  3. Results returned in < 100ms")

	fmt.Println("\n=== EXERCISE ===")
	fmt.Println("1. Create autocomplete for your favorite searches")
	fmt.Println("2. Add different frequency counts")
	fmt.Println("3. Implement case-insensitive matching (already done!)")
	fmt.Println("4. Add a 'recent' bonus to frequently accessed recent items")
	fmt.Println("5. Implement prefix boost: if prefix is exact match, rank higher")
	fmt.Println("\nBonus: Compare speed with hash table lookup on large datasets")
}

// To run:
// go run 04_autocomplete_system.go

// === INTERVIEW QUESTION ===
// "Design an autocomplete system that handles 1M searches/day"
//
// Solution:
// 1. Use Trie for O(P + K) prefix search
// 2. Cache top K results for each popular prefix
// 3. Use async updates to avoid blocking user
// 4. Track frequency and recency with weighted scoring
// 5. Use database for persistence
// 6. Add load balancing if traffic > single machine capacity
