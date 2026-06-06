package main

import "fmt"

// === TRIE FUNDAMENTALS ===
// A Trie is a tree where each node represents a character.
// Each path from root to a marked node represents a word.

// TrieNode represents a single character in the trie
type TrieNode struct {
	children map[rune]*TrieNode // Maps character to child node
	isEnd    bool               // true if this node marks the end of a word
}

// NewTrieNode creates a new, empty TrieNode
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

// Trie is the root node of the entire data structure
type Trie struct {
	root *TrieNode
}

// NewTrie creates an empty Trie
func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode(),
	}
}

// === VISUALIZATION ===
// Let's build the trie for "cat" and "car":
//
//       root
//       /
//      c
//      |
//      a
//     / \
//    t   r
//   [E] [E]   <- [E] marks end of word
//
// When we insert "cat":
//   1. Go to root
//   2. Follow/create path: c -> a -> t
//   3. Mark 't' as end of word (isEnd = true)

// Insert adds a word to the Trie
func (t *Trie) Insert(word string) {
	node := t.root

	// For each character in the word
	for _, ch := range word {
		// If this character doesn't exist as a child, create it
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = NewTrieNode()
		}
		// Move to the child node
		node = node.children[ch]
	}

	// Mark the final node as the end of a word
	node.isEnd = true
}

// Search returns true if the exact word exists in the Trie
func (t *Trie) Search(word string) bool {
	node := t.root

	// Follow the path for each character
	for _, ch := range word {
		if child, exists := node.children[ch]; exists {
			node = child
		} else {
			// Character not found, word doesn't exist
			return false
		}
	}

	// We found the path, but is it marked as a word?
	return node.isEnd
}

// StartsWith returns true if any word in the Trie starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root

	// Follow the path for the prefix
	for _, ch := range prefix {
		if child, exists := node.children[ch]; exists {
			node = child
		} else {
			return false
		}
	}

	// We found the prefix path (doesn't matter if it's a word end)
	return true
}

// === HELPER: Print the trie structure ===
func (t *Trie) printHelper(node *TrieNode, prefix string, indent string) {
	if node.isEnd {
		fmt.Printf("%s[WORD: %s]\n", indent, prefix)
	}

	for ch, child := range node.children {
		fmt.Printf("%s%c\n", indent, ch)
		t.printHelper(child, prefix+string(ch), indent+"  ")
	}
}

// Print displays the entire Trie structure
func (t *Trie) Print() {
	fmt.Println("Trie structure:")
	fmt.Println("root")
	t.printHelper(t.root, "", "  ")
}

func main() {
	fmt.Println("=== TRIE BASICS ===\n")

	// Create a new Trie
	trie := NewTrie()

	fmt.Println("1. Insert words into Trie")
	words := []string{"cat", "car", "card", "care", "careful", "can", "dog", "dodge", "door"}
	for _, word := range words {
		trie.Insert(word)
		fmt.Printf("  Inserted: %s\n", word)
	}

	fmt.Println("\n2. Trie structure:")
	trie.Print()

	fmt.Println("\n3. Search for exact words")
	searchTerms := []string{"cat", "ca", "cat!", "car", "care", "dog", "doll"}
	for _, term := range searchTerms {
		found := trie.Search(term)
		fmt.Printf("  Search('%s'): %v\n", term, found)
	}

	fmt.Println("\n4. Check for prefixes (StartsWith)")
	prefixes := []string{"ca", "car", "do", "doo", "bat"}
	for _, prefix := range prefixes {
		exists := trie.StartsWith(prefix)
		fmt.Printf("  StartsWith('%s'): %v\n", prefix, exists)
	}

	fmt.Println("\n=== KEY INSIGHTS ===")
	fmt.Println("1. Insert adds a word by following/creating a path")
	fmt.Println("2. Search checks if exact word exists (requires isEnd=true)")
	fmt.Println("3. StartsWith checks if prefix exists (doesn't care about isEnd)")
	fmt.Println("4. 'ca' is a prefix of 'cat' and 'car', but not a word itself")
	fmt.Println("5. 'cat' is both a prefix and a word")

	fmt.Println("\n=== EXERCISE ===")
	fmt.Println("1. Create a new Trie and insert: 'hello', 'world', 'help', 'hell'")
	fmt.Println("2. Search for: 'hello', 'hell', 'he', 'help'")
	fmt.Println("3. Check prefixes: 'hel', 'wor', 'he', 'wo'")
	fmt.Println("4. Print the Trie structure")
	fmt.Println("\nUncomment and modify the code below to complete the exercise:")

	// Exercise: Build your own trie
	// myTrie := NewTrie()
	// myTrie.Insert("hello")
	// myTrie.Insert("world")
	// myTrie.Insert("help")
	// myTrie.Insert("hell")
	//
	// fmt.Println("\nMy Trie:")
	// myTrie.Print()
	//
	// fmt.Println("\nSearching in my trie:")
	// fmt.Printf("Search('hello'): %v\n", myTrie.Search("hello"))
	// fmt.Printf("Search('hell'): %v\n", myTrie.Search("hell"))
	// fmt.Printf("Search('he'): %v\n", myTrie.Search("he"))
	//
	// fmt.Println("\nPrefix checks:")
	// fmt.Printf("StartsWith('hel'): %v\n", myTrie.StartsWith("hel"))
	// fmt.Printf("StartsWith('he'): %v\n", myTrie.StartsWith("he"))
}

// To run:
// go run 01_trie_basics.go
