package main

func main() {

}

type TrieNode struct {
	Val       map[byte]*TrieNode
	EndOfWord bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{Val: make(map[byte]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
	curr := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := curr.Val[c]; !ok {
			curr.Val[c] = &TrieNode{Val: make(map[byte]*TrieNode)}
		}
		curr = curr.Val[c]
	}
	curr.EndOfWord = true
}

func (this *Trie) Search(word string) bool {
	curr := this.root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if _, ok := curr.Val[c]; !ok {
			return false
		}
		curr = curr.Val[c]
	}
	return curr.EndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	curr := this.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i]
		if _, ok := curr.Val[c]; !ok {
			return false
		}
		curr = curr.Val[c]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
