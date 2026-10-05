# Data Structures Basics — Go

Implementación de la especificación [06_Data_Structures_Basics](https://yorche3.github.io/programming_languages/core/algorithms/06_Data_Structures_Basics/) en **Go**, con gestión de dependencias mediante **Go Modules** y el framework de pruebas **testify**.

Cuatro estructuras de datos construidas desde cero sobre un único tipo `Node` compartido: **Node** (celda enlazada), **LinkedList** (lista enlazada con punteros a cabeza y cola), **Stack** (pila LIFO) y **Queue** (cola FIFO). Cada ADT gestiona independientemente sus punteros y contador; no hay delegación de unas estructuras en otras ni uso de colecciones de la biblioteca estándar.

---

## 📂 Archivos y estructura / Files & Structure

| Archivo / Directorio | Propósito |
|----------------------|-----------|
| `src/data_structures_basics.go` | Implementación de `Node`, `LinkedList`, `Stack` y `Queue` sobre el mismo tipo de celda enlazada. |
| `tests/data_structures_basics_test.go` | Pruebas unitarias: 4 tests que cubren los casos de la especificación para cada estructura. |
| `go.mod` | Definición del módulo Go y dependencias. |
| `go.sum` | Checksums de las dependencias descargadas. |
| `README.md` | Este archivo. |

Con relación a la implementación base en **Ada**, que separa especificación (`data_structures_basics.ads`) e implementación (`data_structures_basics.adb`), Go reúne contrato y código en un único archivo. El motor `go test` descubre automáticamente las funciones con firma `func TestXxx(t *testing.T)`, sin suites ni registros manuales, así que no hace falta el `run_tests` de la especificación.

**Estructura de directorios esperada:**

```text
data_structures_basics/              # Módulo Go
├── src/
│   └── data_structures_basics.go    # Node, LinkedList, Stack, Queue
├── tests/
│   └── data_structures_basics_test.go # 4 tests × casos de la especificación
├── go.mod                           # Módulo y dependencias
├── go.sum                           # Checksums de dependencias
└── README.md                        # Este archivo
```

---

## 🛠️ Enfoque y construcción / Approach & Build

**ES:** El proyecto se creó manualmente, sin herramientas de scaffolding. A diferencia de **Ada** (que usa `alr init --lib`), en Go basta con `go mod init` y escribir el código fuente.

**EN:** The project was created manually, without scaffolding tools. Unlike **Ada** (which uses `alr init --lib`), in Go it's enough to run `go mod init` and write the source code.

### Inicialización / Initialization

```bash
# 1. Inicializar el módulo / Initialize the module
go mod init example.com/data_structures_basics

# 2. Agregar testify como dependencia / Add testify dependency
go get github.com/stretchr/testify/assert
```

---

## 📄 Configuración clave / Key Configuration

| Archivo / File | Propósito / Purpose |
|----------------|---------------------|
| `go.mod` | Declara el módulo (`example.com/data_structures_basics`) y la versión de Go (1.21). |
| `go.sum` | Checksums de las dependencias descargadas (`testify` y sus transitivas). |

No se usan archivos de configuración adicionales. El código fuente vive en `src/` (paquete `src`) y las pruebas en `tests/` (paquete `tests` con *dot-import* para acceder a los símbolos exportados).

---

## 🚀 Compilación y ejecución / Build & Run

```bash
# Verificar el código / Vet the code
go vet ./...

# Ejecutar las pruebas / Run tests
go test ./... -v
```

**Salida real / Actual output:**

```text
?       example.com/data_structures_basics/src  [no test files]
=== RUN   TestNode
=== RUN   TestNode/initialize_and_observe_value_and_link
=== RUN   TestNode/initialize_another_node,_link_and_traverse
--- PASS: TestNode (0.00s)
    --- PASS: TestNode/initialize_and_observe_value_and_link (0.00s)
    --- PASS: TestNode/initialize_another_node,_link_and_traverse (0.00s)
=== RUN   TestLinkedList
=== RUN   TestLinkedList/empty_state
=== RUN   TestLinkedList/insert_at_both_ends
=== RUN   TestLinkedList/delete_first_occurrence
=== RUN   TestLinkedList/absent_value
=== RUN   TestLinkedList/empty_the_list
--- PASS: TestLinkedList (0.00s)
    --- PASS: TestLinkedList/empty_state (0.00s)
    --- PASS: TestLinkedList/insert_at_both_ends (0.00s)
    --- PASS: TestLinkedList/delete_first_occurrence (0.00s)
    --- PASS: TestLinkedList/absent_value (0.00s)
    --- PASS: TestLinkedList/empty_the_list (0.00s)
=== RUN   TestStack
=== RUN   TestStack/empty_state_and_failed_removal
=== RUN   TestStack/LIFO_and_non-mutating_peek
=== RUN   TestStack/removal_and_reuse
=== RUN   TestStack/empty_after_removal
--- PASS: TestStack (0.00s)
    --- PASS: TestStack/empty_state_and_failed_removal (0.00s)
    --- PASS: TestStack/LIFO_and_non-mutating_peek (0.00s)
    --- PASS: TestStack/removal_and_reuse (0.00s)
    --- PASS: TestStack/empty_after_removal (0.00s)
=== RUN   TestQueue
=== RUN   TestQueue/empty_state_and_failed_removal
=== RUN   TestQueue/FIFO_and_non-mutating_peek
=== RUN   TestQueue/removal_and_reuse
=== RUN   TestQueue/empty_after_removal
--- PASS: TestQueue (0.00s)
    --- PASS: TestQueue/empty_state_and_failed_removal (0.00s)
    --- PASS: TestQueue/FIFO_and_non-mutating_peek (0.00s)
    --- PASS: TestQueue/removal_and_reuse (0.00s)
    --- PASS: TestQueue/empty_after_removal (0.00s)
PASS
ok      example.com/data_structures_basics/tests        (cached)
```

---

## 🧠 Algoritmos y operaciones / Algorithms & Operations

### Node

| Operación / Operation | Entrada → salida / Input → output | Complejidad / Complexity | Notas / Notes |
|---|---|---|---|
| `NewNode(v int) *Node` | `int → *Node` | `O(1)` | Crea un nodo con valor `v` y enlace `nil`. Equivalente a `init(value)`. |
| `Value() int` | `*Node → int` | `O(1)` | Devuelve el valor del nodo. Equivalente a `get_value()`. |
| `Next() *Node` | `*Node → *Node` | `O(1)` | Devuelve el nodo enlazado o `nil` si no hay enlace. Equivalente a `get_next()`. |
| `SetNext(next *Node)` | `*Node, *Node → void` | `O(1)` | Actualiza el enlace del nodo. Equivalente a `set_next(next)`. |

### LinkedList

| Operación / Operation | Entrada → salida / Input → output | Complejidad / Complexity | Notas / Notes |
|---|---|---|---|
| `Head() int` | `*LinkedList → int` | `O(1)` | Devuelve el valor de la cabeza o `-1` si la lista está vacía. Equivalente a `get_head()`. |
| `InsertHead(v int)` | `*LinkedList, int → void` | `O(1)` | Inserta al principio. Equivalente a `insert_head(value)`. |
| `InsertTail(v int)` | `*LinkedList, int → void` | `O(1)` | Inserta al final. Equivalente a `insert_tail(value)`. |
| `Remove(v int) bool` | `*LinkedList, int → bool` | `O(n)` | Elimina la primera aparición de `v`. Devuelve `true` si lo eliminó, `false` si `v` no está. Equivalente a `delete(value)`. |
| `IsEmpty() bool` | `*LinkedList → bool` | `O(1)` | Devuelve `true` si la lista está vacía. Equivalente a `is_empty()`. |
| `Len() int` | `*LinkedList → int` | `O(1)` | Devuelve el número de nodos. Equivalente a `size()`. |

### Stack

| Operación / Operation | Entrada → salida / Input → output | Complejidad / Complexity | Notas / Notes |
|---|---|---|---|
| `Push(v int)` | `*Stack, int → void` | `O(1)` | Apila un valor. Equivalente a `push(value)`. |
| `Pop() int` | `*Stack → int` | `O(1)` | Desapila y devuelve el tope, o `-1` si la pila está vacía. Equivalente a `pop()`. |
| `Peek() int` | `*Stack → int` | `O(1)` | Devuelve el tope sin extraerlo, o `-1` si la pila está vacía. Equivalente a `peek()`. |
| `IsEmpty() bool` | `*Stack → bool` | `O(1)` | Devuelve `true` si la pila está vacía. Equivalente a `is_empty()`. |
| `Len() int` | `*Stack → int` | `O(1)` | Devuelve el número de elementos. Equivalente a `size()`. |

### Queue

| Operación / Operation | Entrada → salida / Input → output | Complejidad / Complexity | Notas / Notes |
|---|---|---|---|
| `Enqueue(v int)` | `*Queue, int → void` | `O(1)` | Encola un valor al final. Equivalente a `enqueue(value)`. |
| `Dequeue() int` | `*Queue → int` | `O(1)` | Desencola y devuelve el frente, o `-1` si la cola está vacía. Equivalente a `dequeue()`. |
| `Peek() int` | `*Queue → int` | `O(1)` | Devuelve el frente sin extraerlo, o `-1` si la cola está vacía. Equivalente a `peek()`. |
| `IsEmpty() bool` | `*Queue → bool` | `O(1)` | Devuelve `true` si la cola está vacía. Equivalente a `is_empty()`. |
| `Len() int` | `*Queue → int` | `O(1)` | Devuelve el número de elementos. Equivalente a `size()`. |

---

## 🧩 Decisiones de diseño / Design decisions

| Decisión / Decision | Alternativa considerada / Alternative | Razón / Reason |
|---|---|---|
| Un único tipo `Node` compartido por las tres estructuras | Tipos de nodo separados para cada ADT | La especificación exige un único `Node` compartido; cada ADT gestiona sus propios punteros (`head`/`tail`, `top`, `front`/`rear`) y contador. |
| El valor cero de `LinkedList`, `Stack` y `Queue` es la instancia vacía | Requerir una función `Init()` explícita | Go permite que el valor cero de un struct sea usable inmediatamente. Los punteros `nil` y el contador `0` ya representan el estado vacío, así que no hace falta llamar a `init()` antes de usar las estructuras. |
| `NewNode(v int)` como constructor explícito | Usar el valor cero de `Node` | `Node` necesita un valor inicial; el valor cero (`value=0`, `next=nil`) no es significativo. `NewNode` es el equivalente idiomático de `init(value)`. |
| Mutaciones sin valor de retorno (`InsertHead`, `Push`, `Enqueue`, etc.) | Devolver la estructura modificada | Go es mutable; las operaciones mutan la instancia recibida. No hace falta devolver la estructura ni un indicador de éxito (las inserciones no fallan por capacidad). |
| Contador `count` en cada estructura | Calcular el tamaño recorriendo los nodos | Mantener un contador `O(1)` para `Len()` es más eficiente que recorrer la estructura `O(n)`. |

---

## 🔀 Adaptaciones idiomáticas / Idiomatic adaptations

| Especificación / Specification | Adaptación / Adaptation | Justificación / Justification |
|---|---|---|
| `init()` para inicializar cada estructura | El valor cero de `LinkedList`, `Stack` y `Queue` ya es la instancia vacía | Go permite que el valor cero de un struct sea usable. Los punteros `nil` y el contador `0` representan el estado vacío sin necesidad de una función `Init()` explícita. |
| `Node.init(value)` | `NewNode(v int) *Node` | Go no tiene constructores; las funciones `New*` son el patrón idiomático para crear instancias. |
| `get_value()`, `get_next()`, `get_head()` | `Value()`, `Next()`, `Head()` | Go usa nombres sin prefijos `get_`; los métodos acceden directamente al valor. |
| `set_next(next)` | `SetNext(next *Node)` | Go usa `Set` para mutadores, sin prefijo `set_` en el nombre del método. |
| `delete(value)` | `Remove(v int) bool` | Go prefiere `Remove` para eliminar elementos de una colección. |
| `size()` | `Len() int` | Go usa `Len` para obtener la longitud de una colección (convención de la biblioteca estándar). |
| `is_empty()` | `IsEmpty() bool` | Go usa `IsEmpty` como método booleano (convención común). |
| Entrada nula o inválida | No se modela explícitamente | La especificación no exige manejar entradas nulas para este módulo; las estructuras se inicializan con el valor cero o con `NewNode`. |

---

## 🚨 Indicadores de fallo / Failure indicators

| Operación / Operation | Situación de fallo / Failure situation | Indicador / Indicator | Ejemplo / Example |
|---|---|---|---|
| `Head()` | Lista vacía | `-1` | `Head()` devuelve `-1` cuando `IsEmpty()` es `true`. |
| `Pop()` | Pila vacía | `-1` | `Pop()` devuelve `-1` cuando `IsEmpty()` es `true`. |
| `Peek()` (Stack) | Pila vacía | `-1` | `Peek()` devuelve `-1` cuando `IsEmpty()` es `true`. |
| `Dequeue()` | Cola vacía | `-1` | `Dequeue()` devuelve `-1` cuando `IsEmpty()` es `true`. |
| `Peek()` (Queue) | Cola vacía | `-1` | `Peek()` devuelve `-1` cuando `IsEmpty()` es `true`. |
| `Remove(v)` | Valor `v` no está en la lista | `false` | `Remove(99)` devuelve `false` si `99` no existe en la lista. |
| `Next()` | Enlace ausente | `nil` | `Next()` devuelve `nil` cuando el nodo no tiene enlace. |

**ES:** El indicador de fallo es `-1` para operaciones que devuelven `int` y `nil` para operaciones que devuelven punteros. Go no tiene `null`; usa `nil` como la representación de ausencia para punteros, slices, mapas, canales y funciones.

**EN:** The failure indicator is `-1` for operations returning `int` and `nil` for operations returning pointers. Go has no `null`; it uses `nil` as the absent representation for pointers, slices, maps, channels and functions.

---

## ✅ Cobertura de pruebas / Test coverage

| Caso de la especificación / Specification case | Cubierto / Covered | Prueba / Test | Notas / Notes |
|---|---|:--:|---|
| **Node**: Inicializar y observar valor/enlace | Sí | `TestNode/initialize_and_observe_value_and_link` | Crea `a` con `NewNode(10)`, verifica `Value()=10` y `Next()==nil`. |
| **Node**: Inicializar otro nodo, enlazar y recorrer | Sí | `TestNode/initialize_another_node,_link_and_traverse` | Crea `b` con `NewNode(20)`, enlaza `a→b`, verifica `Next().Value()=20` y `b.Next()==nil`. |
| **LinkedList**: Estado vacío | Sí | `TestLinkedList/empty_state` | Verifica `Head()=-1`, `IsEmpty()=true`, `Len()=0`. |
| **LinkedList**: Insertar por ambos extremos | Sí | `TestLinkedList/insert_at_both_ends` | Inserta `10, 20` al final y `5` al principio; verifica `Head()=5`, `Len()=4`. |
| **LinkedList**: Eliminar primera aparición | Sí | `TestLinkedList/delete_first_occurrence` | Elimina `10` (primera aparición); verifica `Head()=5`, `Len()=3`, `Remove()=true`. |
| **LinkedList**: Valor ausente | Sí | `TestLinkedList/absent_value` | Intenta eliminar `99`; verifica `Head()=5`, `Len()=3`, `Remove()=false`. |
| **LinkedList**: Vaciar | Sí | `TestLinkedList/empty_the_list` | Elimina `5, 20, 10`; verifica `IsEmpty()=true`, `Len()=0`, `Head()=-1`. |
| **Stack**: Estado vacío y extracción fallida | Sí | `TestStack/empty_state_and_failed_removal` | Verifica `Peek()=-1`, `Pop()=-1`, `IsEmpty()=true`, `Len()=0`. |
| **Stack**: LIFO y `peek` no mutante | Sí | `TestStack/LIFO_and_non-mutating_peek` | Apila `10, 20, 30`; verifica `Peek()=30`, `Len()=3`. |
| **Stack**: Extracción y reutilización | Sí | `TestStack/removal_and_reuse` | Desapila `30`, apila `40`, desapila `40, 20, 10`; verifica `IsEmpty()=true`, `Len()=0`. |
| **Stack**: Vacío tras extracción | Sí | `TestStack/empty_after_removal` | Desapila de pila vacía; verifica `Pop()=-1`, `IsEmpty()=true`. |
| **Queue**: Estado vacío y extracción fallida | Sí | `TestQueue/empty_state_and_failed_removal` | Verifica `Peek()=-1`, `Dequeue()=-1`, `IsEmpty()=true`, `Len()=0`. |
| **Queue**: FIFO y `peek` no mutante | Sí | `TestQueue/FIFO_and_non-mutating_peek` | Encola `10, 20, 30`; verifica `Peek()=10`, `Len()=3`. |
| **Queue**: Extracción y reutilización | Sí | `TestQueue/removal_and_reuse` | Desencola `10`, encola `40`, desencola `20, 30, 40`; verifica `IsEmpty()=true`, `Len()=0`. |
| **Queue**: Vacío tras extracción | Sí | `TestQueue/empty_after_removal` | Desencola de cola vacía; verifica `Dequeue()=-1`, `IsEmpty()=true`. |

**Total de pruebas:** 4 tests con 15 subcasos (2 para `Node`, 5 para `LinkedList`, 4 para `Stack`, 4 para `Queue`).

---

## ⚠️ Limitaciones conocidas / Known limitations

| Limitación / Limitation | Impacto / Impact | Alternativa o plan / Workaround or plan |
|---|---|---|
| Ninguna | No hay limitaciones conocidas | La implementación cumple el contrato de la especificación y todas las pruebas pasan. |

---

## 📝 Notas de implementación / Implementation Notes

**ES:** Go es un lenguaje mutable con recolección de basura, así que las operaciones mutan la instancia recibida en lugar de devolver una estructura nueva. El valor cero de `LinkedList`, `Stack` y `Queue` ya es la instancia vacía (punteros `nil`, contador `0`), lo que elimina la necesidad de una función `Init()` explícita. Solo `Node` requiere un constructor explícito (`NewNode`) porque su valor cero no es significativo.

Los punteros en Go se declaran con `*` y la ausencia de enlace se representa con `nil`. No hay aritmética de punteros como en C, pero los punteros permiten mutar la instancia recibida sin copiarla.

La separación de paquetes `src` y `tests` con *dot-import* permite acceder a los símbolos exportados de `src` desde `tests` sin prefijo, simulando la estructura de la especificación. Los símbolos deben empezar con mayúscula (`PascalCase`) para ser exportados.

**EN:** Go is a mutable language with garbage collection, so operations mutate the received instance instead of returning a new structure. The zero value of `LinkedList`, `Stack` and `Queue` is already the empty instance (pointers `nil`, count `0`), which eliminates the need for an explicit `Init()` function. Only `Node` requires an explicit constructor (`NewNode`) because its zero value is not meaningful.

Pointers in Go are declared with `*` and link absence is represented with `nil`. There is no pointer arithmetic as in C, but pointers allow mutating the received instance without copying it.

The separation of packages `src` and `tests` with *dot-import* allows accessing exported symbols from `src` in `tests` without a prefix, simulating the specification's structure. Symbols must start with an uppercase letter (`PascalCase`) to be exported.

**ES:** Este proyecto también está implementado en otros lenguajes. Explora el repositorio principal para consultar las demás versiones.

**EN:** This project is also implemented in other languages. Explore the main repository to see the other versions.

---

## 🔍 Checklist de validación / Validation checklist

- [x] La suite nativa se ejecutó y su salida real está copiada en este README.
- [x] Cada caso de la especificación tiene su fila en _Cobertura de pruebas_ (o `Omitido` con razón).
- [x] Cada desviación del pseudocódigo o de la ubicación esperada está en _Adaptaciones idiomáticas_.
- [x] Cada operación con fallo posible está en _Indicadores de fallo_.
- [x] No hay rutas absolutas del autor, credenciales ni salidas inventadas.
- [x] Los enlaces relativos resuelven dentro del repositorio y el documento es bilingüe.
- [x] Ninguna sección repite lo que ya dice la especificación.

---

*[← Volver a Algorithms](../README.md)*

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
