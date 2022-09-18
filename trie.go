package main

// Trie node to data structure
type Node struct {
	Childrens [26]*Node
	WordEnds  bool
}

// Trie data structure
type Trie struct {
	Root *Node
}

// Insert word to Trie
func (t *Trie) Insert(word string) {
	current := t.Root

	for _, letter := range word {
		index := letter - 'a'
		if current.Childrens[index] == nil {
			current.Childrens[index] = &Node{}
		}
		current = current.Childrens[index]
	}

	current.WordEnds = true
}

// Depth First Search for Trie
// Inputs Trie.Root, Current String, Out array
func dfs(root *Node, currentStr string, out *[]string) {
	if root == nil {
		return
	}

	if root.WordEnds {
		(*out) = append((*out), currentStr)
	}

	for index, child := range root.Childrens {
		dfs(child, currentStr+string(rune('a'+index)), out)
	}
}

// Finds Closest string to input (word)
// Returns an array of strings with all the closest words to the input in the dictionary
// Could use a LinkedList in order to improve performance
func (t *Trie) FindClosestStrings(word string) []string {
	out := []string{}
	current := t.Root

	for _, wr := range word {
		index := wr - 'a'
		if current.Childrens[index] == nil {
			return nil
		}
		current = current.Childrens[index]
	}

	dfs(current, word, &out)

	return out
}
