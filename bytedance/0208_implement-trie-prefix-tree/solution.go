/*
 * Author: robin-luo
 * Created time: 2024-03-19 17:13:56
 * Last Modified by: robin-luo
 * Last Modified time: 2024-03-19 17:20:18
 */

package main

import "fmt"

type Trie struct {
	Children [26]*Trie
	IsEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	for _, ch := range word {
		if this.Children[ch-'a'] == nil {
			this.Children[ch-'a'] = &Trie{}
		}

		this = this.Children[ch-'a']
	}

	this.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	this = this.searchPrefix(word)
	return this != nil && this.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	this = this.searchPrefix(prefix)
	return this != nil
}

func (this *Trie) searchPrefix(prefix string) *Trie {
	for _, ch := range prefix {
		if this.Children[ch-'a'] == nil {
			return nil
		}

		this = this.Children[ch-'a']
	}

	return this
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
}
