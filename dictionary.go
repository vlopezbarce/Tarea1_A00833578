// Valeria López Barcelata A00833578
// Desarrollo de aplicaciones avanzadas de ciencias computacionales
// Módulo 3 - Compiladores: Tarea 1
// 3 de abril de 2025
// Implementación de un diccionario en Go

package main

// Inicializa la estructura de datos para el diccionario
type Dictionary struct {
	words map[string]string
}

// Agrega una palabra y su valor al diccionario
func (d *Dictionary) Set(word string, value string) {
	if d.words == nil {
		d.words = make(map[string]string) // Inicializa el mapa si es nil
	}
	d.words[word] = value
}

// Obtiene el valor de una palabra en el diccionario
func (d *Dictionary) Get(word string) string {
	if value, exists := d.words[word]; exists {
		return value
	}
	panic("Error: La palabra no existe en el diccionario")
}

// Elimina una palabra del diccionario
func (d *Dictionary) Remove(word string) {
	delete(d.words, word)
}

// Verifica si el diccionario está vacío
func (d *Dictionary) IsEmpty() bool {
	return len(d.words) == 0
}

// Verifica si una palabra existe en el diccionario
func (d *Dictionary) Contains(word string) bool {
	_, exists := d.words[word]
	return exists
}

// Regresa el número de palabras en el diccionario
func (d *Dictionary) Size() int {
	return len(d.words)
}

// Elimina todas las palabras del diccionario
func (d *Dictionary) Clear() {
	d.words = make(map[string]string)
}

// Regresa solo las llaves del diccionario
func (d *Dictionary) Keys() []string {
	keys := make([]string, 0, len(d.words))
	for word := range d.words {
		keys = append(keys, word)
	}
	return keys
}

// Regresa solo los valores del diccionario
func (d *Dictionary) Values() []string {
	values := make([]string, 0, len(d.words))
	for _, value := range d.words {
		values = append(values, value)
	}
	return values
}

// Regresa todos los pares de palabras y valores del diccionario
func (d *Dictionary) Items() map[string]string {
	return d.words
}
