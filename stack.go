// Valeria López Barcelata A00833578
// Desarrollo de aplicaciones avanzadas de ciencias computacionales
// Módulo 3 - Compiladores: Tarea 1
// 3 de abril de 2025
// Implementación de una pila de enteros en Go

package main

// Inicializa una pila de enteros
type Stack struct {
	items []int
}

// Agrega un elemento al final de la pila
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Elimina el último elemento de la pila y lo devuelve
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		panic("Error: No hay elementos en la pila")
	}
	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item
}

// Verifica si la pila está vacía
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Regresa el tamaño de la pila
func (s *Stack) Size() int {
	return len(s.items)
}

// Elimina todos los elementos de la pila
func (s *Stack) Clear() {
	s.items = []int{}
}

// Imprime los elementos de la pila en una línea
func (s *Stack) Print() {
	for _, item := range s.items {
		print(item, " ")
	}
	println()
}
