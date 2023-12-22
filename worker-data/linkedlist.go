package main

type T interface{} // Define T como cualquier tipo

type NodeList struct {
	Value 		T
	Next  		*NodeList
	Previous 	*NodeList
}

type LinkedList struct {
	Head  	*NodeList	
	Tail  	*NodeList
	Length  int
}

func (l* LinkedList) Add(value T) {
	node := &NodeList{Value: value}
	if l.Head == nil {
		l.Head = node
		l.Tail = node	
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.Length++
}

func (l* LinkedList) TopDeque() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
	l.Length--
}

func (l* LinkedList) BottomDeque() {
	if l.Head == l.Tail{
		l.Head = nil
		l.Tail = nil
	} else {
		var aux = l.Tail
		l.Tail = l.Tail.Previous
		l.Tail.Previous = aux.Previous
	}
}