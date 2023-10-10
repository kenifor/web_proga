package main

import (
	"fmt"
	"main/list" // Импортируем ваш пакет list
)

func main() {
	// Создаем новый список
	myList := list.NewList()

	// Добавляем элементы в список
	myList.Add(1)
	myList.Add(2)
	myList.Add(3)
	myList.Add(74)
	myList.Add(8)
	myList.Add(2)
	myList.Add(4)
	myList.Add(55)
	myList.Add(23)

	// Выводим длину списка
	fmt.Println("Длина списка:", myList.Len())

	// Выводим все элементы списка
	fmt.Println("Все элементы списка:")
	myList.Print()

	// Удаляем элемент по индексу
	myList.RemoveByIndex(2)

	// Удаляем элемент по значению
	myList.RemoveByValue(2)

	// Выводим все элементы списка после удаления
	fmt.Println("Все элементы списка после удаления:")
	myList.Print()

	// Получаем элемент по индексу
	value, ok := myList.GetByIndex(0)
	if ok {
		fmt.Println("Элемент по индексу 0:", value)
	}

	// Получаем индекс элемента по значению
	index, ok := myList.GetByValue(4)
	if ok {
		fmt.Println("Индекс элемента со значением 4:", index)
	}

	// Получаем все индексы элементов по значению
	indices, ok := myList.GetAllByValue(2)
	if ok {
		fmt.Println("Индексы элементов со значением 2:", indices)
	}

	// Получаем все элементы списка
	values, ok := myList.GetAll()
	if ok {
		fmt.Println("Все элементы списка:", values)
	}

	// Очищаем список
	myList.Clear()

	// Выводим длину списка после очистки
	fmt.Println("Длина списка после очистки:", myList.Len())
}
