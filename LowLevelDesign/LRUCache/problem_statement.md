# LRU Cache

## Problem Statement
Design and implement a Least Recently Used (LRU) Cache — a fixed-capacity
key-value store that evicts the least recently used entry when it runs out
of space.

## Requirements
1. Cache size should be configurable.
2. `Put(key, value)` should add a new element or update an existing one.
3. `Get(key)` should return the value if present, otherwise `-1`.
4. Recently accessed or written elements should become Most Recently Used (MRU).
5. When capacity is full, the Least Recently Used (LRU) element should be evicted.
6. Both `Get` and `Put` must run in O(1) time.

## Entities / Classes

1. **Node**
   - Fields: `key`, `value`, `prev`, `next`
   - Represents a single cache entry. Holds pointers to its neighbours so it
     can be unlinked and relinked in O(1) as part of a doubly linked list.

2. **LRUCache**
   - Fields: `capacity`, `cache` (map of key → `*Node`), `head`, `tail`
     (sentinel nodes)
   - Main class exposing `Get(key)` and `Put(key, value)`.
   - Combines a hash map (O(1) lookup) with a doubly linked list (O(1)
     reordering/eviction) to satisfy the time complexity requirement.
