# Desarrollo de Aplicaciones Avanzadas de Ciencias Computacionales

## Módulo 3: Compiladores
### 📄 Tarea 1

📅 **Fecha de entrega:** 3 de abril de 2025

**Nombre:** Valeria López Barcelata

**Matrícula:** A00833578

---

## Descripción

Esta tarea consiste en la implementación y documentación de tres estructuras de datos fundamentales:

1️⃣ **STACK (LIFO)** - Estructura de pila que sigue la política *Last In, First Out*.

2️⃣ **QUEUE (FIFO)** - Cola que sigue la política *First In, First Out*.

3️⃣ **DICTIONARY (MAP)** - Diccionario o mapa que permite almacenamiento y acceso eficiente de datos en formato clave-valor.

---

## Casos de Prueba

Para validar el correcto funcionamiento de cada estructura, se realizaron diversas pruebas:

### 🔹 **STACK**
**Inserción y eliminación de elementos:** Se agregan tres elementos (`1, 2, 3`) y se eliminan en orden LIFO (*Last In, First Out*).

**Tamaño de la pila:** Se verifica que el tamaño refleje correctamente la cantidad de elementos.

**Verificación de elementos:** Se imprimen los elementos para asegurar el orden correcto.

**Eliminación total:** Se vacía la pila y se verifica que esté vacía.

### 🔹 **QUEUE**
**Inserción y eliminación de elementos:** Se agregan tres elementos (`1, 2, 3`) y se eliminan en orden FIFO (*First In, First Out*).

**Tamaño de la cola:** Se verifica que el tamaño refleje correctamente la cantidad de elementos.

**Verificación de elementos:** Se imprimen los elementos para asegurar el orden correcto.

**Eliminación total:** Se vacía la cola y se verifica que esté vacía.

### 🔹 **DICTIONARY**
**Inserción de datos:** Se agregan claves y valores (`clave1: valor1`, `clave2: valor2`, `clave3: valor3`).

**Consulta de valores:** Se obtienen valores por clave y se validan los resultados.

**Modificación de valores:** Se cambia un valor y se verifica la actualización.

**Eliminación de claves:** Se eliminan claves y se verifica que no existan.

**Recuperación de llaves y valores:** Se imprimen todas las llaves y valores disponibles.

**Verificación de diccionario vacío:** Se elimina todo y se verifica que el diccionario esté vacío.

---

## Estructura del Proyecto
📁 Tarea1_A00833578

│── 📜 main.go

├── 📜 stack.go

├── 📜 queue.go

├── 📜 dictionary.go

│── 📜 README.md

│── 📜 go.mod

---

## Instrucciones de Ejecución  

1️⃣ **Clonar este repositorio:**  
```sh
git clone https://github.com/tuusuario/Tarea1_A00833578.git
cd Tarea1_A00833578
```
2️⃣ **Inicializar el módulo de Go:**
```
go mod init Tarea1_A00833578
go mod tidy
```
3️⃣ **Ejecutar el programa:**
```
go run Tarea1_A00833578
```
