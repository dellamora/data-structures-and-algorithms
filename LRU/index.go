// Terms to know:
// LRU Cache: Least Recently Used Cache

// load factor: The amount of data poinst vc the amount of storage
// key: a value that is hashable and is used to look up data. The hash has to be consistent
// value: A value that is associated with a key
// collision: When two keys hash to the same value (map to the same cell)

package main

import "container/list"

type LRUCache struct {
    capacity int
    cache    map[int]*list.Element
    list     *list.List
}

type entry struct {
    key   int
    value int
}

func NewLRUCache(capacity int) *LRUCache {
    return &LRUCache{
        capacity: capacity,
        cache:    make(map[int]*list.Element),
        list:     list.New(),
    }
}

func (lru *LRUCache) Get(key int) int {
    if ele, ok := lru.cache[key]; ok {
        lru.list.MoveToFront(ele)
        return ele.Value.(*entry).value
    }
    return -1
}

func (lru *LRUCache) Put(key, value int) {
    if ele, ok := lru.cache[key]; ok {
        ele.Value.(*entry).value = value
        lru.list.MoveToFront(ele)
    } else {
        if len(lru.cache) >= lru.capacity {
            back := lru.list.Back()
            if back != nil {
                deletedEntry := lru.list.Remove(back).(*entry)
                delete(lru.cache, deletedEntry.key)
            }
        }

        newEntry := &entry{key, value}
        ele := lru.list.PushFront(newEntry)
        lru.cache[key] = ele
    }
}

func main() {
    lruCache := NewLRUCache(2)

    lruCache.Put(1, 1)
    lruCache.Put(2, 2)
    println(lruCache.Get(1)) 
    lruCache.Put(3, 3)
    println(lruCache.Get(2)) 
    lruCache.Put(4, 4)
    println(lruCache.Get(1)) 
    println(lruCache.Get(3)) 
    println(lruCache.Get(4)) 
}
