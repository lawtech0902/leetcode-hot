/*
__author__ = 'robin-luo'
__date__ = '2023/11/28 16:05'
*/

package solution

import "fmt"

func suggestedProducts(products []string, searchWord string) [][]string {
	root := Constructor()
	for _, product := range products {
		root.Insert(product)
	}

	res := make([][]string, 0)
	for i := 1; i <= len(searchWord); i++ {
		tempStr := searchWord[:i]
		tempRes := root.Search(tempStr)
		res = append(res, tempRes)
	}

	return res
}

type Trie struct {
	child [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	for _, ch := range word {
		ch -= 'a'
		if this.child[ch] == nil {
			this.child[ch] = &Trie{}
		}

		this = this.child[ch]
	}

	this.isEnd = true
}

func (this *Trie) Search(word string) []string {
	var (
		res  []string
		isOk = true
	)

	for _, ch := range word {
		ch -= 'a'
		if this.child[ch] == nil {
			isOk = false
			break
		}

		this = this.child[ch]
	}

	if !isOk {
		return res
	}

	dfs(word, this, &res)
	return res
}

func dfs(curWord string, this *Trie, res *[]string) {
	if len(*res) == 3 {
		return
	}

	if this.isEnd {
		*res = append(*res, curWord)
	}

	for i := 0; i < 26; i++ {
		if this.child[i] != nil {
			num := i + 'a'
			str := fmt.Sprintf("%c", num)
			dfs(curWord+str, this.child[i], res)
		}
	}
}
