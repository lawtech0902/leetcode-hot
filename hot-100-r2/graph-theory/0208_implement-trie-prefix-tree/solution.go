/*
__author__ = 'robin-luo'
__date__ = '2024/01/25 16:05'
*/

package solution

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	for _, ch := range word {
		ch -= 'a'
		if this.children[ch] == nil {
			this.children[ch] = &Trie{}
		}

		this = this.children[ch]
	}

	this.isEnd = true
}

func (this *Trie) SearchPrefix(prefix string) *Trie {
	for _, ch := range prefix {
		ch -= 'a'
		if this.children[ch] == nil {
			return nil
		}

		this = this.children[ch]
	}

	return this
}

func (this *Trie) Search(word string) bool {
	this = this.SearchPrefix(word)
	return this != nil && this.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.SearchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
