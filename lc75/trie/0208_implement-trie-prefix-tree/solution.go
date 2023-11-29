/*
__author__ = 'robin-luo'
__date__ = '2023/11/28 15:43'
*/

package solution

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, i := range word {
		i -= 'a'
		if node.child[i] == nil {
			node.child[i] = &Trie{}
		}

		node = node.child[i]
	}

	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, i := range word {
		i -= 'a'
		if node.child[i] == nil {
			return false
		}

		node = node.child[i]
	}

	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, i := range prefix {
		i -= 'a'
		if node.child[i] == nil {
			return false
		}

		node = node.child[i]
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
