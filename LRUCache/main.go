package LRUCache

type listNode struct {
	prev *listNode
	next *listNode
	val  int
	key  int
}

type LRUCache struct {
	m    map[int]*listNode
	head *listNode
	tail *listNode
	cap  int
	len  int
}

func Constructor(capacity int) LRUCache {
	fakeLast := &listNode{val: -1, prev: nil, next: nil}
	fakeHead := &listNode{val: -2, prev: nil, next: fakeLast}
	fakeLast.prev = fakeHead
	return LRUCache{
		cap:  capacity,
		len:  0,
		m:    make(map[int]*listNode, capacity),
		head: fakeHead,
		tail: fakeLast,
	}
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.m[key]
	if !ok {
		return -1
	}

	if c.head.next == v {
		return v.val
	}

	c.swapFirst(v)

	return v.val
}

func (c *LRUCache) swapFirst(v *listNode) {
	frst := c.head.next
	next := v.next
	prev := v.prev

	c.head.next = v
	v.prev = c.head
	prev.next = next
	next.prev = prev
	frst.prev = v
	v.next = frst
}

func (c *LRUCache) Put(key int, value int) {
	if c.Get(key) != -1 {
		c.head.next.val = value
		return
	}

	var cur *listNode
	if c.len >= c.cap {
		c.len--
		cur = c.tail.prev
		cur.prev.next = c.tail
		c.tail.prev = cur.prev
		delete(c.m, cur.key)
	} else {
		cur = &listNode{}
	}

	c.len++
	next := c.head.next
	cur.val = value
	cur.key = key
	cur.next = next
	cur.prev = c.head
	c.head.next = cur
	next.prev = cur
	c.m[key] = cur
}
