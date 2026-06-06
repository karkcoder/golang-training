package main

import "fmt"

// === TRIE DELETION (THE HARD PART!) ===
// Deleting a word is tricky because we need to clean up unused nodes
// while keeping nodes that are part of other words.
//
// Example: Trie has "cat" and "car"
//   Deleting "cat" should:
//   - Mark 't' node's isEnd as false
//   - Don't delete 't' node (yet) because path c-a is shared with "car"

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

// === DELETION (COMPLEX!) ===
// Strategy: Use recursion with backtracking
//
// Rules:
// 1. Only delete if the node is not part of another word
// 2. Only delete if the node has no children
// 3. Mark isEnd = false, then bubble up to clean empty nodes

// deleteHelper removes a word and cleans up nodes
// Returns true if this node should be deleted by parent
func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		// Reached end of word
		if !node.isEnd {
			return false // Word doesn't exist, nothing to delete
		}

		// Mark this node as not the end of a word
		node.isEnd = false

		// This node can be deleted by parent only if it has no children
		return len(node.children) == 0
	}

	ch := rune(word[index])

	// Recursively delete in child
	if child, exists := node.children[ch]; !exists {
		return false // Path doesn't exist
	} else {
		shouldDeleteChild := t.deleteHelper(child, word, index+1)

		// If child should be deleted, remove it from our children map
		if shouldDeleteChild {
			delete(node.children, ch)
		}

		// This node can be deleted only if:
		// 1. It's not the end of a word
		// 2. It has no children
		return !node.isEnd && len(node.children) == 0
	}
}

// Delete removes a word from the Trie
func (t *Trie) Delete(word string) bool {
	return t.deleteHelper(t.root, word, 0)
}

// === HELPER: Count total words ===
func (t *Trie) countWords(node *TrieNode) int {
	count := 0
	if node.isEnd {
		count = 1
	}
	for _, child := range node.children {
		count += t.countWords(child)
	}
	return count
}

// GetWordCount returns total words in Trie
func (t *Trie) GetWordCount() int {
	return t.countWords(t.root)
}

// === HELPER: Print all words ===
func (t *Trie) getAllWords(node *TrieNode, prefix string, words *[]string) {
	if node.isEnd {
		*words = append(*words, prefix)
	}
	for ch, child := range node.children {
		t.getAllWords(child, prefix+string(ch), words)
	}
}

// GetAllWords returns all words in the Trie
func (t *Trie) GetAllWords() []string {
	var words []string
	t.getAllWords(t.root, "", &words)
	return words
}

func main() {
	fmt.Println("=== TRIE DELETION ===\n")

	trie := NewTrie()

	// Build a trie
	words := []string{"cat", "car", "card", "care", "dog"}
	fmt.Println("1. Insert words:")
	for _, word := range words {
		trie.Insert(word)
		fmt.Printf("  Inserted: %s\n", word)
	}

	fmt.Printf("\nTotal words: %d\n", trie.GetWordCount())
	fmt.Printf("All words: %v\n", trie.GetAllWords())

	fmt.Println("\n2. Delete 'cat' (leaf node, safe to delete)")
	deleted := trie.Delete("cat")
	fmt.Printf("  Deleted: %v\n", deleted)
	fmt.Printf("  Search 'cat': %v (should be false)\n", trie.Search("cat"))
	fmt.Printf("  Search 'car': %v (should be true)\n", trie.Search("car"))
	fmt.Printf("  Total words: %d\n", trie.GetWordCount())
	fmt.Printf("  All words: %v\n", trie.GetAllWords())

	fmt.Println("\n3. Try to delete non-existent word 'bat'")
	deleted = trie.Delete("bat")
	fmt.Printf("  Deleted: %v (should be false)\n", deleted)
	fmt.Printf("  All words still: %v\n", trie.GetAllWords())

	fmt.Println("\n4. Delete 'car' (internal node, shared prefix 'ca')")
	deleted = trie.Delete("car")
	fmt.Printf("  Deleted: %v\n", deleted)
	fmt.Printf("  Search 'car': %v\n", trie.Search("car"))
	fmt.Printf("  Search 'card': %v (should still exist)\n", trie.Search("card"))
	fmt.Printf("  Search 'care': %v (should still exist)\n", trie.Search("care"))
	fmt.Printf("  Total words: %d\n", trie.GetWordCount())
	fmt.Printf("  All words: %v\n", trie.GetAllWords())

	fmt.Println("\n5. Delete 'card'")
	deleted = trie.Delete("card")
	fmt.Printf("  Deleted: %v\n", deleted)
	fmt.Printf("  All words: %v\n", trie.GetAllWords())

	fmt.Println("\n6. Delete 'care'")
	deleted = trie.Delete("care")
	fmt.Printf("  Deleted: %v\n", deleted)
	fmt.Printf("  All words: %v\n", trie.GetAllWords())

	fmt.Println("\n7. Now 'ca' prefix is completely gone, let's verify")
	fmt.Printf("  Search 'ca': %v (should be false)\n", trie.Search("ca"))
	fmt.Printf("  Search 'dog': %v (should be true)\n", trie.Search("dog"))

	fmt.Println("\n=== DELETION CHALLENGES ===")
	fmt.Println("Challenge 1: What happens when you delete a word that's a prefix of another?")
	fmt.Println("  Example: Trie has 'cat' and 'car', delete 'ca' (never inserted)")
	fmt.Println("           The nodes still exist, but 'ca' is not marked as a word")

	fmt.Println("\nChallenge 2: What about shared prefixes?")
	fmt.Println("  When you delete 'cat' from {'cat', 'car'}:")
	fmt.Println("  - The 't' node becomes unused (no children, not an end)")
	fmt.Println("  - But 'a' and 'c' are still used by 'car'")
	fmt.Println("  - Only 't' gets deleted, 'c' and 'a' remain")

	fmt.Println("\n=== EXERCISE ===")
	fmt.Println("1. Build a Trie with: 'apple', 'app', 'apply', 'banana'")
	fmt.Println("2. Delete 'app' and verify 'apple' and 'apply' still exist")
	fmt.Println("3. Delete 'apple' and verify 'app' and 'apply' still exist")
	fmt.Println("4. Delete 'apply' and check remaining words")
	fmt.Println("5. Test deleting non-existent words")
	fmt.Println("6. Try: Delete all words one by one, verify each step")
	fmt.Println("\nBonus: Implement a function that cleans up empty branches")
}

// To run:
// go run 03_trie_delete.go
