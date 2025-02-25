package main

import "fmt"

// doubly linked list
type List struct {
	head *Node
	tail *Node
	size int
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Remove(node *Node) {
	// fmt.Println("called ", l.size)

	// set tail and next
	if node.next == nil {
		l.tail = node.prev
	}
	if node.prev == nil {
		l.head = node.next
	}

	node.KickMe()

	l.size -= 1
}

// for debugging purposes
func (l *List) PrintAll() {
	n := l.First()
	for {
		fmt.Printf("%d, ", n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}
	fmt.Println("----")
}

func (l *List) PushHead(k, v int) *Node {
	node := &Node{key: k, value: v}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		node.next.prev = node
	}

	l.head = node
	l.size += 1
	return node
}

type Node struct {
	value int
	key   int
	prev  *Node
	next  *Node
}

func (n *Node) KickMe() {
	node := n
	n = node.prev
	n.next = node.next
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type LRUCache struct {
	keyData  map[int]*Node //hashtable
	listData *List
	size     int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keyData:  make(map[int]*Node),
		listData: &List{},
		size:     capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.listData.head == nil {
		return -1
	}

	node, exist := this.keyData[key]
	if !exist {
		return -1
	}

	// push recently use to head
	this.listData.PushHead(node.key, node.value)

	// remove node from list
	this.listData.Remove(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node := this.listData.PushHead(key, value)
	this.keyData[key] = node

	// remove last node
	// fmt.Println(this.listData.size, "|", this.size)
	if this.listData.size > this.size {
		// fmt.Println("here")
		delete(this.keyData, this.listData.Last().key)
		this.listData.Remove(this.listData.Last())
	}
}

func main() {
	// [[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

	lRUCache := Constructor(2)
	lRUCache.Put(1, 0)
	lRUCache.Put(2, 2)
	fmt.Println(lRUCache.Get(1))
	lRUCache.listData.PrintAll()
	lRUCache.Put(3, 3)
	lRUCache.listData.PrintAll()
	fmt.Println(lRUCache.Get(2))
	lRUCache.listData.PrintAll()
	lRUCache.Put(4, 4)
	lRUCache.listData.PrintAll()
	fmt.Println(lRUCache.Get(1))
	fmt.Println(lRUCache.Get(3))
	fmt.Println(lRUCache.Get(4))
}

func main1() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	fmt.Println("mysize ", lRUCache.listData.size)
	fmt.Println("mylast ", lRUCache.listData.Last().value)
	lRUCache.listData.PrintAll()

	lRUCache.Get(1) // return 1
	fmt.Println("mysize ", lRUCache.listData.size)
	fmt.Println("mylast ", lRUCache.listData.Last().value)
	lRUCache.listData.PrintAll()

	lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println("mysize ", lRUCache.listData.size)
	fmt.Println("mylast ", lRUCache.listData.Last().value)
	lRUCache.listData.PrintAll()

	fmt.Println(lRUCache.Get(2)) // returns -1 (not found)

	lRUCache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	fmt.Println(lRUCache.Get(3)) // return 3
	fmt.Println(lRUCache.Get(4)) // return 4

	// list := &List{}
	// list.PushHead(1)
	// list.PushHead(2)
	// list.PushHead(3)
	// list.PushHead(4)
	// list.PushHead(5)

	// n := list.First()
	// for {
	// 	fmt.Println(n.value)
	// 	n = n.Next()
	// 	if n == nil {
	// 		break
	// 	}
	// }

	// fmt.Println("-----")
	// list.Last().Prev().KickMe()

	// n2 := list.First()
	// for {
	// 	fmt.Println(n2.value)
	// 	n2 = n2.Next()
	// 	if n2 == nil {
	// 		break
	// 	}
	// }
}

// Notes,
// - use hashtable to map key, with node
// - put, push to head
// - get, push to head, remove current
// - if put meet capacity, remove tail
// - get, return -1 if not found
