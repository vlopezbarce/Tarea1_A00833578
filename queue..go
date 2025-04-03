// Valeria López Barcelata A00833578
// Desarrollo de aplicaciones avanzadas de ciencias computacionales
// Módulo 3 - Compiladores: Tarea 1
// 3 de abril de 2025
// Implementación de una fila de enteros en Go

package main

// Inicializa una fila de enteros
type Queue struct {
	items []int
}

// Agrega un elemento al final de la fila
func (q *Queue) Push(item int) {
	q.items = append(q.items, item)
}

// Elimina y devuelve el primer elemento de la fila
func (q *Queue) Pop() int {
	if len(q.items) == 0 {
		panic("Error: No hay elementos en la fila")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Verifica si la fila está vacía
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Regresa el tamaño de la fila
func (q *Queue) Size() int {
	return len(q.items)
}

// Elimina todos los elementos de la fila
func (q *Queue) Clear() {
	q.items = []int{}
}

// Imprime los elementos de la fila en una línea
func (q *Queue) Print() {
	for _, item := range q.items {
		print(item, " ")
	}
	println()
}
