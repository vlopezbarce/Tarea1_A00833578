// Valeria López Barcelata A00833578
// Desarrollo de aplicaciones avanzadas de ciencias computacionales
// Módulo 3 - Compiladores: Tarea 1
// 3 de abril de 2025
// Programa para implementar estructuras de datos en Go

package main

import "fmt"

// Casos de prueba para la implementación de la pila
func testStack() {
	stack := Stack{}

	// Agrega elementos a la pila
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Imprime el contenido de la pila
	// Resultado esperado: 1 2 3
	stack.Print()

	// Obtiene el tamaño de la pila
	// Resultado esperado: 3
	fmt.Println("Tamaño de la pila:", stack.Size())

	// Elimina un elemento de la pila
	// Resultado esperado: 3
	fmt.Println("Elemento eliminado:", stack.Pop())

	// Imprime el contenido de la pila
	// Resultado esperado: 1 2
	stack.Print()

	// Verifica si la pila está vacía
	// Resultado esperado: false
	fmt.Println("¿La pila está vacía?", stack.IsEmpty())

	// Elimina todos los elementos de la pila
	stack.Clear()

	// Verifica si la pila está vacía
	// Resultado esperado: true
	fmt.Println("¿La pila está vacía?", stack.IsEmpty())
}

// Casos de prueba para la implementación de la fila
func testQueue() {
	queue := Queue{}

	// Agrega elementos a la fila
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	// Imprime el contenido de la fila
	// Resultado esperado: 1 2 3
	queue.Print()

	// Obtiene el tamaño de la fila
	// Resultado esperado: 3
	fmt.Println("Tamaño de la fila:", queue.Size())

	// Elimina un elemento de la fila
	// Resultado esperado: 1
	fmt.Println("Elemento eliminado:", queue.Pop())

	// Imprime el contenido de la fila
	// Resultado esperado: 2 3
	queue.Print()

	// Verifica si la fila está vacía
	// Resultado esperado: false
	fmt.Println("¿La fila está vacía?", queue.IsEmpty())

	// Elimina todos los elementos de la fila
	queue.Clear()

	// Verifica si la fila está vacía
	// Resultado esperado: true
	fmt.Println("¿La fila está vacía?", queue.IsEmpty())
}

// Casos de prueba para la implementación del diccionario
func testDictionary() {
	dict := Dictionary{}

	// Agrega elementos al diccionario
	dict.Set("clave1", "valor1")
	dict.Set("clave2", "valor2")
	dict.Set("clave3", "valor3")

	// Imprime el contenido del diccionario
	// Resultado esperado: clave1: valor1, clave2: valor2, clave3: valor3
	for key, value := range dict.Items() {
		fmt.Println(key, ":", value)
	}

	// Obtiene el tamaño del diccionario
	// Resultado esperado: 3
	fmt.Println("Tamaño del diccionario:", dict.Size())

	// Verifica si una clave existe en el diccionario
	// Resultado esperado: true
	fmt.Println("¿La clave1 existe?", dict.Contains("clave1"))

	// Obtiene el valor de una clave
	// Resultado esperado: valor1
	fmt.Println("Valor de clave1:", dict.Get("clave1"))

	// Cambia el valor de una clave
	// Resultado esperado: nuevoValor1
	dict.Set("clave1", "nuevoValor1")

	// Obtiene el valor de una clave
	// Resultado esperado: nuevoValor1
	fmt.Println("Valor de clave1:", dict.Get("clave1"))

	// Elimina un elemento del diccionario
	dict.Remove("clave1")

	// Verifica si una clave existe en el diccionario
	// Resultado esperado: false
	fmt.Println("¿La clave1 existe?", dict.Contains("clave1"))

	// Imprime el contenido del diccionario
	// Resultado esperado: clave2: valor2, clave3: valor3
	for key, value := range dict.Items() {
		fmt.Println(key, ":", value)
	}

	// Imprime las llaves del diccionario
	// Resultado esperado: clave2, clave3
	fmt.Println("Llaves del diccionario:", dict.Keys())

	// Imprime los valores del diccionario
	// Resultado esperado: valor2, valor3
	fmt.Println("Valores del diccionario:", dict.Values())

	// Verifica si el diccionario está vacío
	// Resultado esperado: false
	fmt.Println("¿El diccionario está vacío?", dict.IsEmpty())

	// Elimina todos los elementos del diccionario
	dict.Clear()

	// Verifica si el diccionario está vacío
	// Resultado esperado: true
	fmt.Println("¿El diccionario está vacío?", dict.IsEmpty())
}

func main() {
	// Ejecuta los casos de prueba para cada estructura de datos
	fmt.Println("Pruebas de la pila:")
	testStack()
	fmt.Println()

	fmt.Println("Pruebas de la fila:")
	testQueue()
	fmt.Println()

	fmt.Println("Pruebas del diccionario:")
	testDictionary()
	fmt.Println()
}
