package list

type elem struct {
	belongsTo *List
	prev      *elem
	next      *elem
	Value     interface{}
}

type List struct {
	first  *elem
	last   *elem
	length int
}

type Node interface {
	// Gets pointer to the prev Elem
	Prev() *elem
	// Gets pointer to the next elem
	Next() *elem
	// Sets a new elem in prev position
	setBefore(newNode *elem)
	// Sets a new elem in next position
	setAfter(newNode *elem)
	// Deletes the elem. This method sets prev to the next elem and vice versa
	erase()
}

type DoubleLinkedList interface {
	// Returns length of the List
	Length() int
	// Returns first value of the List
	GetFirst() *elem
	// Returns last value of the List
	GetLast() *elem
	//Creates and sets a value as first in the List
	PushFront(value interface{})
	//Creates and sets a value as last in the List
	PushBack(value interface{})
	//Creates and sets a value as first in the List
	PushBefore(value interface{}, n *elem)
	//Creates and sets a value as last in the List
	PushAfter(value interface{}, n *elem)
	//Remove the first value in the List
	RemoveFirst()
	//Remove the last value in the List
	RemoveLast()
	//Swaps elements in List
	Swap(n1 *elem, n2 *elem)
	// Delete elem from List
	Delete(n *elem)
	// Initialize the List
	Init() *List
	//Creates a elem
	createNode(value interface{}) *elem
	// Increase List length
	increaseLength()
	// Decrease List length
	decreaseLength()
	// Initialize list if was not initialized before
	lateInit()
	// Swap elem that are neighbors
	swapNeighbors(n1 *elem, n2 *elem)
	// Swap elem that are not neighbors
	swapNotNeighbors(n1 *elem, n2 *elem)

	setNodeByFirstTime(newNode *elem)
}

func (l *List) Length() int {
	return l.length
}

func (l *List) GetFirst() *elem {
	return l.first
}

func (l *List) GetLast() *elem {
	return l.last
}

func (l *List) PushFront(value interface{}) {
	newNode := l.createNode(value)

	if l.length == 0 {
		l.setNodeByFirstTime(newNode)
	} else {
		oldFirstNode := l.first
		newNode.next = oldFirstNode
		oldFirstNode.prev = newNode

		l.first = newNode
		newNode.belongsTo = l
		l.increaseLength()
	}
}
func (l *List) PushBefore(value interface{}, n *elem) {
	newNode := l.createNode(value)

	if l.length == 0 {
		l.setNodeByFirstTime(newNode)
	} else if n.belongsTo == l {
		if n == l.first {
			l.first = newNode
		}
		n.setBefore(newNode)
		l.increaseLength()
	}
}
func (l *List) PushAfter(value interface{}, n *elem) {
	newNode := l.createNode(value)

	if l.length == 0 {
		l.setNodeByFirstTime(newNode)
	} else if n.belongsTo == l {
		n.setAfter(newNode)
		l.increaseLength()
	}
}

func (l *List) PushBack(value interface{}) {
	newNode := l.createNode(value)

	if l.length == 0 {
		l.setNodeByFirstTime(newNode)
	} else {
		oldLastNode := l.last
		newNode.prev = oldLastNode
		oldLastNode.next = newNode

		l.last = newNode
		newNode.belongsTo = l
		l.increaseLength()
	}
}

func (l *List) RemoveFirst() {
	elemToRemove := l.first
	if elemToRemove.belongsTo == l && l.length > 0 {
		nextNode := elemToRemove.next
		l.first = nextNode
		nextNode.prev = nil
		elemToRemove.next = nil
		l.decreaseLength()
	}
}

func (l *List) RemoveLast() {
	elemToRemove := l.last
	if elemToRemove.belongsTo == l && l.length > 0 {
		prevNode := elemToRemove.prev
		l.last = prevNode
		prevNode.next = nil
		elemToRemove.prev = nil

		l.decreaseLength()
	}
}

func (l *List) Swap(n1 *elem, n2 *elem) {
	if n1 != n2 && n1.belongsTo == l && n2.belongsTo == l {
		if l.first == n1 {
			l.first = n2
		} else if l.first == n2 {
			l.first = n1
		}

		if l.last == n1 {
			l.last = n2
		} else if l.last == n2 {
			l.last = n1
		}

		if n1.next == n2 {
			l.swapNeighbors(n1, n2)
		} else if n2.next == n1 {
			l.swapNeighbors(n2, n1)
		} else {
			l.swapNotNeighbors(n1, n2)
		}
	}
}

func (l *List) createNode(value interface{}) *elem {
	return &elem{Value: value}
}

func (l *List) increaseLength() {
	l.length += 1
}

func (l *List) decreaseLength() {
	l.length -= 1
}

func (l *List) setNodeByFirstTime(newNode *elem) {
	l.first = newNode
	l.last = newNode
	newNode.belongsTo = l
	l.increaseLength()
}

func (l *List) lateInit() {
	if l.first == nil {
		l.Init()
	}
}

func (l *List) swapNeighbors(n1 *elem, n2 *elem) {
	prev1 := n1.prev
	next2 := n2.next
	n2.next = n1
	n1.next = next2
	n2.prev = prev1
	n1.prev = n2
	if prev1 != nil{
		prev1.next = n2
	}
	if next2 != nil{
		next2.prev = n1
	}
}

func (l *List) swapNotNeighbors(n1 *elem, n2 *elem) {
	prev1 := n1.prev
	next1 := n1.next
	prev2 := n2.prev
	next2 := n2.next

	n1.setNext(next2)
	n1.setPrev(prev2)
	n2.setNext(next1)
	n2.setPrev(prev1)
	prev1.setNext(n2)
	next1.setPrev(n2)
	prev2.setNext(n1)
	next2.setPrev(n1)
}

func (l *List) Delete(n *elem) {
	if l == n.belongsTo {
		n.erase()
		l.decreaseLength()
	}
}

// Init initializes or clears list l.
func (l *List) Init() *List {
	l.length = 0
	return l
}

func (n *elem) setBefore(newNode *elem) {
	prevNode := n.prev

	newNode.prev = prevNode
	if prevNode != nil {
		prevNode.next = newNode
	}

	n.prev = newNode
	newNode.next = n

	newNode.belongsTo = n.belongsTo
}

func (n *elem) setAfter(newNode *elem) {
	nextNode := n.next

	newNode.next = nextNode
	nextNode.prev = newNode

	n.next = newNode
	newNode.prev = n

	newNode.belongsTo = n.belongsTo
}

func (n *elem) Prev() *elem {
	return n.prev
}

func (n *elem) Next() *elem {
	return n.next
}

func (n *elem) setNext(n1 *elem) {
	if n != nil {
		if n != n1 {
			n.next = n1
		}
	}
}

func (n *elem) setPrev(n1 *elem) {
	if n != nil {
		if n.prev != n1 {
			n.prev = n1
		}
	}
}

func (n *elem) erase() {
	nextNode := n.next
	prevNode := n.prev

	if nextNode != nil {
		nextNode.prev = prevNode
	}

	if prevNode != nil {
		prevNode.next = nextNode
	}

	if n.belongsTo.first == n {
		n.belongsTo.first = nextNode
	}

	if n.belongsTo.last == n {
		n.belongsTo.last = prevNode
	}

	n.next = nil
	n.prev = nil
	n.belongsTo = nil
}

// New returns an initialized list.
func New() *List {
	return new(List).Init()
}
