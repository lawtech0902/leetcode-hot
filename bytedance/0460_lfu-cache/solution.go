/*
 * Author: robin-luo
 * Created time: 2024-02-27 17:10:51
 * Last Modified by: robin-luo
 * Last Modified time: 2024-02-27 17:11:42
 */

package solution

import "container/list"

type LFUData struct {
	key   int
	value int
	count int
}

type LFUCache struct {
	capacity int
	findMap  map[int]*list.Element
	countMap map[int]*list.List
	minCount int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		findMap:  make(map[int]*list.Element, 0),
		countMap: make(map[int]*list.List, 0),
	}
}

func (l *LFUCache) Get(key int) int {
	val, ok := l.findMap[key]
	if !ok {
		return -1
	}

	return l.update(val)
}

func (l *LFUCache) update(ele *list.Element) int {
	data := ele.Value.(*LFUData)
	curList, ok := l.countMap[data.count]
	if !ok {
		return -1
	}

	curList.Remove(ele)
	if curList.Len() == 0 {
		if data.count == l.minCount {
			l.minCount++
		}
	}

	data.count++
	newList, ok := l.countMap[data.count]
	if !ok {
		newList = list.New()
		l.countMap[data.count] = newList
	}

	newE := newList.PushBack(data)
	l.findMap[data.key] = newE
	return data.value
}

func (l *LFUCache) Put(key int, value int) {
	existEle, ok := l.findMap[key]
	if ok {
		data := existEle.Value.(*LFUData)
		data.value = value
		l.update(existEle)
		return
	}

	if len(l.findMap) == l.capacity {
		minList, ok := l.countMap[l.minCount]
		if !ok {
			return
		}

		delE := minList.Front()
		minList.Remove(delE)
		data := delE.Value.(*LFUData)
		delete(l.findMap, data.key)
	}

	newData := &LFUData{
		key:   key,
		value: value,
		count: 1,
	}

	newList, ok := l.countMap[newData.count]
	if !ok {
		newList = list.New()
		l.countMap[newData.count] = newList
	}

	newEle := newList.PushBack(newData)
	l.findMap[key] = newEle
	l.minCount = newData.count
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
