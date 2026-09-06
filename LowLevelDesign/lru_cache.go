package main

import "fmt"

/*
LRU Cache - Low Level Design

Functional Requirements:
1. Cache size should be configurable.
2. Put(key, value) should add/update an element.
3. Get(key) should return the value if present, otherwise -1.
4. Recently accessed or written elements should become Most Recently Used (MRU).
5. When capacity is full, the Least Recently Used (LRU) element should be evicted.



Design:
We use two data structures together:

1. Hash Map
   key -> *Node

   Provides O(1) lookup for a key.

2. Doubly Linked List
   Maintains the usage order.

   head.next -> Most Recently Used element
   tail.prev -> Least Recently Used element

Why Doubly Linked List?
Because after finding a node using the map, we need to remove it from its
current position and move it to the front in O(1).

Overall:
Get -> O(1)
Put -> O(1)
Space -> O(capacity)
*/

// Node represents one entry stored inside the cache.
//
// Apart from key and value, it stores pointers to the previous and next
// nodes so that it can be removed from the linked list in O(1).
type Node struct {
	key   int
	value int

	prev *Node
	next *Node
}

// LRUCache owns both data structures used by the implementation.
//
// cache provides fast key lookup.
// head and tail are dummy/sentinel nodes used to simplify linked-list
// insertion and deletion.
//
// List structure:
//
// head <-> MRU <-> ... <-> LRU <-> tail
type LRUCache struct {
	capacity int

	// Maps a key directly to the same Node that exists in the linked list.
	cache map[int]*Node

	head *Node
	tail *Node
}

// LRUCacheConstructor initializes an empty cache.
//
// Dummy head and tail nodes are used so that we don't need separate logic
// for inserting/removing the first or last real element.
func LRUCacheConstructor(capacity int) LRUCache {
	cache := make(map[int]*Node, capacity)

	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    cache,
		head:     head,
		tail:     tail,
	}
}

// addToFront inserts a node immediately after head.
//
// The front of the list represents the Most Recently Used position.
func (l *LRUCache) addToFront(node *Node) {
	node.next = l.head.next
	node.prev = l.head

	l.head.next.prev = node
	l.head.next = node
}

// remove disconnects a node from its current position.
//
// Because this is a doubly linked list, both neighbouring nodes
// are directly available, so removal takes O(1).
func (l *LRUCache) remove(node *Node) {
	prev := node.prev
	next := node.next

	prev.next = next
	next.prev = prev
}

// moveToFront marks an existing node as recently used.
//
// First remove it from its current position, then insert it
// immediately after head.
func (l *LRUCache) moveToFront(node *Node) {
	l.remove(node)
	l.addToFront(node)
}

// removeLRUElement removes the Least Recently Used element.
//
// Since tail.prev always points to the LRU node, we can locate it
// directly without traversing the list.
//
// The node must be removed from both:
// 1. the linked list
// 2. the hash map
func (l *LRUCache) removeLRUElement() {
	lastElement := l.tail.prev

	l.remove(lastElement)
	delete(l.cache, lastElement.key)
}

// Get returns the value for a key.
//
// If the key exists:
// - Find it in O(1) using the map.
// - Move it to the front because it was just accessed.
//
// If it does not exist, return -1.
func (l *LRUCache) Get(key int) int {
	node, exists := l.cache[key]

	if !exists {
		return -1
	}

	l.moveToFront(node)

	return node.value
}

// Put adds a new key-value pair or updates an existing one.
//
// Existing key:
// - Update its value.
// - Move it to the front because it was recently written.
//
// New key:
// - If capacity is full, evict the LRU element.
// - Create the new node.
// - Add it to the front.
// - Store its reference in the map.
func (l *LRUCache) Put(key int, value int) {
	if node, exists := l.cache[key]; exists {
		node.value = value
		l.moveToFront(node)
		return
	}

	if len(l.cache) == l.capacity {
		l.removeLRUElement()
	}

	node := &Node{
		key:   key,
		value: value,
	}

	l.addToFront(node)
	l.cache[key] = node
}

func main() {
	cache := LRUCacheConstructor(2)

	cache.Put(1, 10)
	cache.Put(2, 20)

	fmt.Println(cache.Get(1)) // 10

	cache.Put(3, 30)          // Capacity full -> key 2 is evicted.
	fmt.Println(cache.Get(2)) // -1

	cache.Put(4, 40)          // Key 1 is now the LRU -> evicted.
	fmt.Println(cache.Get(1)) // -1

	fmt.Println(cache.Get(3)) // 30
	fmt.Println(cache.Get(4)) // 40
}
