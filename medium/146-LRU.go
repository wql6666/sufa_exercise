package medium


type LRUCache struct {
	Head, Tail *Node
	Cap        int
	NodeMap    map[int]*Node
}

type Node struct {
	Key, Val  int
	Pre, Next *Node
}

func NewLRUCache(cap int) LRUCache {
	head := Node{
		Key:  0,
		Val:  0,
		Pre:  nil,
		Next: nil,
	}
	tail := Node{
		Key:  0,
		Val:  0,
		Pre:  nil,
		Next: nil,
	}
	head.Next = &tail
	tail.Pre = &head
	return LRUCache{
		Head:    &head,
		Tail:    &tail,
		Cap:     cap,
		NodeMap: make(map[int]*Node, 0),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.NodeMap[key]; ok {
		// 节点放到链表最前边
		// 删除节点
		this.DelNode(node)
		// 放到链表首部
		this.AddNode(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key, val int) {
	if node, ok := this.NodeMap[key]; ok {
		this.DelNode(node)
	}

	if len(this.NodeMap) == this.Cap {
		this.DelNode(this.Tail.Pre)
	}

	this.AddNode(&Node{
		Key:  key,
		Val:  val,
		Pre:  nil,
		Next: nil,
	})
}

func (this *LRUCache) DelNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
	delete(this.NodeMap, node.Key)
}

func (this *LRUCache) AddNode(node *Node) {
	this.NodeMap[node.Key] = node
	next := this.Head.Next
	node.Next = next
	node.Pre = this.Head
	this.Head.Next = node
	next.Pre = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
