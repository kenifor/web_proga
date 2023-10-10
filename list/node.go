package list

// Node - элемент списка
type node struct {
	index int64
	value int64
	prev  *node
	next  *node
}
