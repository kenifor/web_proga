package list

import "fmt"

// List представляет собой двусвязный список целых чисел.
type List struct {
	len  int64
	head *node
	tail *node
}

// NewList создает и возвращает новый пустой список.
func NewList() *List {
	return &List{}
}

// Len возвращает текущую длину списка.
func (l *List) Len() int64 {
	return l.len
}

// Add добавляет элемент в конец списка и возвращает его индекс.
func (l *List) Add(value int64) int64 {
	n := &node{value: value, index: l.len}

	if l.len == 0 {
		l.head = n
		l.tail = n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}

	l.len++
	return n.index
}

// RemoveByIndex удаляет элемент по заданному индексу.
func (l *List) RemoveByIndex(index int64) {
	n := l.findNodeByIndex(index)
	if n != nil {
		if n.prev != nil {
			n.prev.next = n.next
		}
		if n.next != nil {
			n.next.prev = n.prev
		}
		if n == l.head {
			l.head = n.next
		}
		if n == l.tail {
			l.tail = n.prev
		}
		l.len--
	}
}

// RemoveByValue удаляет первый найденный элемент с заданным значением.
func (l *List) RemoveByValue(value int64) {
	n := l.findNodeByValue(value)
	if n != nil {
		l.removeNode(n)
	}
}

// RemoveAllByValue удаляет все элементы с заданным значением.
func (l *List) RemoveAllByValue(value int64) {
	for {
		n := l.findNodeByValue(value)
		if n == nil {
			break
		}
		l.removeNode(n)
	}
}

// GetByIndex возвращает значение элемента по заданному индексу и флаг, указывающий, найден ли элемент.
func (l *List) GetByIndex(index int64) (int64, bool) {
	n := l.findNodeByIndex(index)
	if n == nil {
		return 0, false
	}
	return n.value, true
}

// GetByValue возвращает индекс первого найденного элемента с заданным значением и флаг, указывающий, найден ли элемент.
func (l *List) GetByValue(value int64) (int64, bool) {
	n := l.findNodeByValue(value)
	if n == nil {
		return 0, false
	}
	return n.index, true
}

// GetAllByValue возвращает индексы всех элементов с заданным значением и флаг, указывающий, найдены ли элементы.
func (l *List) GetAllByValue(value int64) ([]int64, bool) {
	var ids []int64
	n := l.head
	for n != nil {
		if n.value == value {
			ids = append(ids, n.index)
		}
		n = n.next
	}
	if len(ids) == 0 {
		return nil, false
	}
	return ids, true
}

// GetAll возвращает все элементы списка и флаг, указывающий, пуст ли список.
func (l *List) GetAll() ([]int64, bool) {
	var values []int64
	n := l.head
	for n != nil {
		values = append(values, n.value)
		n = n.next
	}
	if len(values) == 0 {
		return nil, false
	}
	return values, true
}

// Clear очищает весь список.
func (l *List) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

// Print выводит элементы списка в консоль.
func (l *List) Print() {
	n := l.head
	for n != nil {
		fmt.Print(n.value, " ")
		n = n.next
	}
	fmt.Println()
}

// removeNode удаляет указанный узел из списка.
func (l *List) removeNode(n *node) {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n == l.head {
		l.head = n.next
	}
	if n == l.tail {
		l.tail = n.prev
	}
	l.len--
}

// findNodeByIndex ищет узел с заданным индексом в списке и возвращает его, если найден
func (l *List) findNodeByIndex(index int64) *node {
	n := l.head
	for n != nil {
		if n.index == index {
			return n
		}
		n = n.next
	}
	return nil
}

// findNodeByValue ищет узел с заданным значением в списке и возвращает его, если найден.
func (l *List) findNodeByValue(value int64) *node {
	n := l.head
	for n != nil {
		if n.value == value {
			return n
		}
		n = n.next
	}
	return nil
}
