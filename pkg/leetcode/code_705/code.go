package code_705

// 链地址法
type MyHashSet struct {
	base   int
	bucket map[int]*NodeList
}

type NodeList struct {
	value int
	next  *NodeList
	pre   *NodeList
}

func add(key int, nodeList *NodeList)  *NodeList {
	if nodeList == nil {
		return &NodeList{
			value: key,
			next:  nil,
			pre:   nil,
		}
	}
	pNode := nodeList
	for {
		if pNode.value == key {
			return nodeList
		}
		if pNode.next == nil {
			pNode.next =  &NodeList{
				value: key,
				next:  nil,
				pre:   pNode,
			}
			return nodeList
		}
		pNode = pNode.next
	}
	return nodeList
}
func remove(key int, nodeList *NodeList) (newNodeList *NodeList) {
	if nodeList == nil {
		return nil
	}
	pNode := nodeList
	for {
		if pNode == nil {
			return nodeList
		}
		if pNode.value == key {
			if pNode.pre == nil {
				return pNode.next
			}else {
				pNode.pre.next = pNode.next
				return nodeList
			}
		}
		pNode = pNode.next
	}
}
func contains(key int, nodeList *NodeList) bool {
	pNode := nodeList
	for pNode != nil {
		if pNode.value == key{
			return true
		}
		pNode = pNode.next
	}
	return false
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		base:   1000,
		bucket: make(map[int]*NodeList, 0),
	}
}

func (this *MyHashSet) Add(key int) {
	t := key % this.base
	nodeList := this.bucket[t]
	this.bucket[t] = add(key, nodeList)
}

func (this *MyHashSet) Remove(key int) {
	t := key % this.base
	nodeList := this.bucket[t]
	this.bucket[t] = remove(key, nodeList)
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	t := key % this.base
	nodeList := this.bucket[t]
	return contains(key,nodeList)
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
