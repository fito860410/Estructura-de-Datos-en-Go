package main

type Stacker interface {
	add(element int)
	delete()
	getFirst() int
}

type Node struct {
	value    int
	NextNode *Node
}

type Stack struct {
	firstNode *Node
	size      int
}

//Add method
func (s *Stack) add(element int) {
	s.firstNode = &Node{
		value:    element,
		NextNode: s.firstNode,
	}

	s.size++
}

func (s *Stack) delete(position int) {
	s.firstNode = s.firstNode.NextNode
	s.size--
}

func (s *Stack) getFirst() int {
	return s.firstNode.value
}

func (s *Stack) getSize() int {
	return s.size
}
