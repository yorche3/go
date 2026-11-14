# Data Structures Advanced — Go

Implementación de la especificación [07_Data_Structures_Advanced](https://yorche3.github.io/programming_languages/core/algorithms/07_Data_Structures_Advanced/) en **Go**, con gestión de dependencias mediante **Go Modules** y el paquete estándar de pruebas `testing`.

Estructuras de datos avanzadas con múltiples enlaces por nodo construidas desde cero: **TreeNode** (nodo con enlaces izquierdo y derecho), **BinaryTree** (árbol binario de búsqueda / ABB con inserción y búsqueda `O(h)`), **GraphNode** (nodo de vértice con valor embebido y lista de adyacencia ordenada) y **Graph** (grafo dirigido sobre registro lineal de vértices, adyacencias manuales y sin bibliotecas auxiliares de mapas o conjuntos).

---

## 📂 Archivos y estructura / Files & Structure

| Archivo / Directorio | Propósito |
|---|---|
| `src/data_structures_advanced.go` | Implementación de `TreeNode`, `BinaryTree`, `GraphNode` y `Graph`. |
| `tests/data_structures_advanced_test.go` | Pruebas unitarias nativas (`testing`): subtests que cubren los casos de la especificación. |
| `go.mod` | Manifiesto del módulo Go (`example.com/data_structures_advanced`). |
| `.gitignore` | Archivos temporales o de compilación excluidos. |
| `README.md` | Este archivo / This file. |

**Desviación de layout / Layout deviation:**
Respecto a la ubicación sugerida en la especificación (`test/data_structures_advanced_test.ext` y `test/run_tests.ext`), la implementación ubica las pruebas en `tests/` (paquete `tests`), siguiendo la convención homogénea del submódulo Go (`data_structures_basics`, `numbers`). En Go no se requiere un script `run_tests`, ya que la herramienta estándar `go test ./...` descubre y ejecuta automáticamente todas las funciones `TestXxx`.

**Estructura de directorios:**

```text
data_structures_advanced/
├── src/
│   └── data_structures_advanced.go
├── tests/
│   └── data_structures_advanced_test.go
├── .gitignore
├── go.mod
└── README.md
```

---

## 🛠️ Enfoque y construcción / Approach & Build

**ES:** El proyecto se creó manualmente, sin herramientas pesadas de scaffolding. En Go basta con inicializar el módulo (`go mod init`) y estructurar el código fuente y las pruebas en sus respectivos paquetes. A diferencia de `data_structures_basics`, aquí se utiliza exclusivamente la biblioteca estándar `testing` (con `reflect.DeepEqual` y helpers locales) sin dependencias externas como `testify`.

**EN:** The project was created manually without scaffolding generators. In Go, initializing the module (`go mod init`) and structuring source code and tests into packages is sufficient. Unlike `data_structures_basics`, tests rely strictly on standard `testing` and `reflect.DeepEqual` with local assertions, avoiding third-party dependencies such as `testify`.

### Inicialización / Initialization

```bash
# Inicializar el módulo Go / Initialize the Go module
go mod init example.com/data_structures_advanced
```

---

## 📄 Configuración clave / Key Configuration

| Archivo / File | Propósito / Purpose |
|---|---|
| `go.mod` | Declara el módulo (`example.com/data_structures_advanced`) y la directiva de toolchain Go (`go 1.26.5`). |

No se requieren archivos de configuración adicionales ni dependencias externas en `go.sum`. El código de producción vive en `src/` (paquete `src`) y las pruebas en `tests/` (paquete `tests` importando con dot-import `. "example.com/data_structures_advanced/src"`).

---

## 🚀 Compilación y ejecución / Build & Run

```bash
# Compilar y verificar estáticamente / Build and vet code
go vet ./...

# Ejecutar las pruebas unitarias / Run unit tests
go test -v ./...
```

**Salida real / Actual output:**

```text
?   	example.com/data_structures_advanced/src	[no test files]
=== RUN   TestTreeNode
=== RUN   TestTreeNode/Create_a_tree_node
=== RUN   TestTreeNode/Absent_initial_links
=== RUN   TestTreeNode/Link_two_children
=== RUN   TestTreeNode/Traverse_each_link
=== RUN   TestTreeNode/Unlink_children
--- PASS: TestTreeNode (0.00s)
    --- PASS: TestTreeNode/Create_a_tree_node (0.00s)
    --- PASS: TestTreeNode/Absent_initial_links (0.00s)
    --- PASS: TestTreeNode/Link_two_children (0.00s)
    --- PASS: TestTreeNode/Traverse_each_link (0.00s)
    --- PASS: TestTreeNode/Unlink_children (0.00s)
=== RUN   TestBinaryTree
=== RUN   TestBinaryTree/Create_empty
=== RUN   TestBinaryTree/Insert_the_root
=== RUN   TestBinaryTree/Insert_lower_and_higher
=== RUN   TestBinaryTree/Multilevel_branches
=== RUN   TestBinaryTree/Missing_value_in_nonempty_tree
=== RUN   TestBinaryTree/Idempotent_duplicate
=== RUN   TestBinaryTree/Ordered_insertions
--- PASS: TestBinaryTree (0.00s)
    --- PASS: TestBinaryTree/Create_empty (0.00s)
    --- PASS: TestBinaryTree/Insert_the_root (0.00s)
    --- PASS: TestBinaryTree/Insert_lower_and_higher (0.00s)
    --- PASS: TestBinaryTree/Multilevel_branches (0.00s)
    --- PASS: TestBinaryTree/Missing_value_in_nonempty_tree (0.00s)
    --- PASS: TestBinaryTree/Idempotent_duplicate (0.00s)
    --- PASS: TestBinaryTree/Ordered_insertions (0.00s)
=== RUN   TestGraph
=== RUN   TestGraph/Create_empty
=== RUN   TestGraph/Isolated_vertex
=== RUN   TestGraph/Nonconsecutive_labels_and_directed_edge
=== RUN   TestGraph/Duplicate_vertex_preserves_adjacency
=== RUN   TestGraph/Duplicate_edge
=== RUN   TestGraph/Missing_endpoint_without_mutation
=== RUN   TestGraph/Sorted_neighbours_and_counts
=== RUN   TestGraph/Missing-edge_query
=== RUN   TestGraph/Neighbours_of_missing_vertex
=== RUN   TestGraph/Self-loop_and_duplicate
=== RUN   TestGraph/Neighbour-result_isolation
--- PASS: TestGraph (0.00s)
    --- PASS: TestGraph/Create_empty (0.00s)
    --- PASS: TestGraph/Isolated_vertex (0.00s)
    --- PASS: TestGraph/Nonconsecutive_labels_and_directed_edge (0.00s)
    --- PASS: TestGraph/Duplicate_vertex_preserves_adjacency (0.00s)
    --- PASS: TestGraph/Duplicate_edge (0.00s)
    --- PASS: TestGraph/Missing_endpoint_without_mutation (0.00s)
    --- PASS: TestGraph/Sorted_neighbours_and_counts (0.00s)
    --- PASS: TestGraph/Missing-edge_query (0.00s)
    --- PASS: TestGraph/Neighbours_of_missing_vertex (0.00s)
    --- PASS: TestGraph/Self-loop_and_duplicate (0.00s)
    --- PASS: TestGraph/Neighbour-result_isolation (0.00s)
PASS
ok  	example.com/data_structures_advanced/tests	0.003s
```

---

## 🧠 Algoritmos y operaciones / Algorithms & Operations

### TreeNode

| Operación / Operation | Entrada → salida / Input → output | Complejidad / Complexity | Notas / Notes |
|---|---|---|---|
| `NewTreeNode(value int)` | `int → *TreeNode` | `O(1)` | Constructor con valor embebido; enlaces iniciales a `nil`. |
| `Value()` | `*TreeNode → int` | `O(1)` | Consulta del valor embebido del nodo. Equivalente a `get_value()`. |
| `Left()` | `*TreeNode → *TreeNode` | `O(1)` | Devuelve puntero al hijo izquierdo o `nil` si ausente. Equivalente a `get_left()`. |
| `Right()` | `*TreeNode → *TreeNode` | `O(1)` | Devuelve puntero al hijo derecho o `nil` si ausente. Equivalente a `get_right()`. |
| `SetLeft(node *TreeNode)` | `*TreeNode, *TreeNode → void` | `O(1)` | Modifica o desenlaza (`nil`) el hijo izquierdo. Equivalente a `set_left(node)`. |
| `SetRight(node *TreeNode)` | `*TreeNode, *TreeNode → void` | `O(1)` | Modifica o desenlaza (`nil`) el hijo derecho. Equivalente a `set_right(node)`. |

### BinaryTree

| Operación / Operation | Entrada → salida / Input → output | Complejidad / Complexity | Notas / Notes |
|---|---|---|---|
| `NewBinaryTree()` | `void → *BinaryTree` | `O(1)` | Constructor del árbol binario vacío (`root=nil`, `count=0`). Equivalente a `init()`. |
| `Insert(value int)` | `*BinaryTree, int → bool` | `O(h)` | Inserción iterativa siguiendo la propiedad ABB; duplicado no altera tamaño (`true`). |
| `Contains(value int)` | `*BinaryTree, int → bool` | `O(h)` | Búsqueda iterativa en ABB; devuelve `true` si existe, `false` si no. |
| `IsEmpty()` | `*BinaryTree → bool` | `O(1)` | Devuelve `true` si `count == 0`. |
| `Len()` | `*BinaryTree → int` | `O(1)` | Retorna el número de elementos únicos en el árbol. Equivalente a `size()`. |

### GraphNode & Graph

| Operación / Operation | Entrada → salida / Input → output | Complejidad / Complexity | Notas / Notes |
|---|---|---|---|
| `NewGraph()` | `void → *Graph` | `O(1)` | Inicializa grafo vacío sin vértices ni aristas. Equivalente a `init()`. |
| `FindVertex(label int)` | `*Graph, int → *GraphNode` | `O(V)` | Búsqueda manual lineal sobre la lista de vértices; `nil` si no existe. |
| `AddVertex(label int)` | `*Graph, int → bool` | `O(V)` | Registra vértice si no existe; idempotente si ya existe (`true`). |
| `AddEdge(u, v int)` | `*Graph, int, int → bool` | `O(V + d(u))` | Requiere extremos existentes (`false` si falta alguno); idempotente; orden creciente en adyacencia. |
| `HasEdge(u, v int)` | `*Graph, int, int → bool` | `O(V + d(u))` | Recorre adyacencia lineal; `false` si arista o `u` no existen. |
| `Neighbors(u int)` | `*Graph, int → []int` | `O(V + d(u))` | Copia independiente de vecinos; `nil` si `u` no existe, slice vacío `[]int{}` si está aislado. |
| `IsEmpty()` | `*Graph → bool` | `O(1)` | `true` si `vertexCount == 0`. |
| `VertexCount()` | `*Graph → int` | `O(1)` | Contador directo `O(1)` de vértices registrados. |
| `EdgeCount()` | `*Graph → int` | `O(1)` | Contador directo `O(1)` de aristas dirigidas únicas. |

---

## 🧩 Decisiones de diseño / Design decisions

| Decisión / Decision | Alternativa considerada / Alternative | Razón / Reason |
|---|---|---|
| Valor embebido en `TreeNode` y `GraphNode` | Herencia o composición de un struct `Node` común | Go favorece composición simple; al tratarse de un único campo `value int` sin comportamiento compartido, el valor embebido cumple la especificación directamente sin sobrecarga sintáctica. |
| Registro lineal manual de vértices y lista enlazada de vecinos | Mapas hash `map[int]*GraphNode` o `map[int][]int` | La especificación prohíbe delegar el registro e índices en tablas hash o árboles de biblioteca; el encadenamiento explícito cumple con el requisito pedagógico de implementación manual. |
| Devolución de slice recién asignado en `Neighbors(u)` | Devolver puntero a la lista enlazada interna de adyacencia | Garantiza aislamiento del resultado: modificaciones externas sobre el slice no corrompen la estructura interna del grafo. |
| Inserción recursiva simple `insertNeighbor` para adyacencia ordenada | Ordenar tras inserción con `sort.Slice` | Mantiene la lista enlazada de adyacencia siempre ordenada e impide duplicados en tiempo `O(d(u))` sin bibliotecas externas. |

---

## 🔀 Adaptaciones idiomáticas / Idiomatic adaptations

| Especificación / Specification | Adaptación / Adaptation | Justificación / Justification |
|---|---|---|
| Layout `test/data_structures_advanced_test.ext` y `run_tests` | Carpeta `tests/` y runner estándar `go test` | En Go, `go test ./...` es la herramienta estándar universal; no se emplean scripts de ejecución de pruebas ad-hoc. |
| `size()` | `Len()` | En Go, la convención idiomática para colecciones y tipos con tamaño es `Len()`. |
| `get_value()`, `get_left()`, `get_right()`, `get_neighbors()` | `Value()`, `Left()`, `Right()`, `Neighbors()` | En Go no se usa el prefijo `get_` en getters exportados (Go Style Guide). |
| `Node` base | Valor embebido (`value int`) en nodos | Go no tiene herencia de clases; la especificación permite explícitamente el valor embebido. |
| Representación de ausencia en enlaces | Puntero `nil` | En Go, `nil` es la representación canónica y unificada para punteros nulos a nodos. |

---

## 🚨 Indicadores de fallo / Failure indicators

| Operación / Operation | Situación de fallo / Failure situation | Indicador / Indicator | Ejemplo / Example |
|---|---|---|---|
| `Left()`, `Right()` | Enlace hijo ausente | `nil` | `n.Left() == nil` |
| `SetLeft(nil)`, `SetRight(nil)` | Desenlace explícito | `nil` | `n.SetLeft(nil)` deja el enlace ausente |
| `Contains(v)` | Valor ausente en el árbol | `false` | `tree.Contains(99) == false` |
| `FindVertex(label)` | Vértice no registrado | `nil` | `g.FindVertex(99) == nil` |
| `AddEdge(u, v)` | `u` o `v` inexistente | `false` | `g.AddEdge(10, 99) == false` |
| `HasEdge(u, v)` | Arista ausente o `u` inexistente | `false` | `g.HasEdge(10, 30) == false` |
| `Neighbors(u)` | Vértice `u` inexistente en el grafo | `nil` | `g.Neighbors(99) == nil` (a diferencia de `[]int{}` en vértice aislado) |

---

## ✅ Cobertura de pruebas / Test coverage

| Caso de la especificación / Specification case | Cubierto / Covered | Prueba / Test | Notas / Notes |
|---|---|:--:|---|
| **TreeNode**: Crear un nodo de árbol | Sí | `TestTreeNode/Create_a_tree_node` | Comprueba `Value() == 10`. |
| **TreeNode**: Enlaces iniciales ausentes | Sí | `TestTreeNode/Absent_initial_links` | Comprueba `Left() == nil` y `Right() == nil`. |
| **TreeNode**: Enlazar dos hijos | Sí | `TestTreeNode/Link_two_children` | Verifica que `Left()` y `Right()` apunten a los hijos con sus valores. |
| **TreeNode**: Recorrer cada enlace | Sí | `TestTreeNode/Traverse_each_link` | Navega a hijos 20 y 30 y verifica que sus enlaces sean `nil`. |
| **TreeNode**: Desenlazar hijos | Sí | `TestTreeNode/Unlink_children` | Desvincula hijo izquierdo conservando el derecho y luego desvincula el derecho. |
| **TreeNode**: Adaptación inmutable | Omitido | — | No aplica: Go maneja estructuras mutables por puntero. |
| **BinaryTree**: Crear vacío | Sí | `TestBinaryTree/Create_empty` | `IsEmpty() == true`, `Len() == 0`, `Contains(50) == false`. |
| **BinaryTree**: Insertar la raíz | Sí | `TestBinaryTree/Insert_the_root` | Inserta 50; `Len() == 1`, `IsEmpty() == false`, `Contains(50) == true`. |
| **BinaryTree**: Insertar menor y mayor | Sí | `TestBinaryTree/Insert_lower_and_higher` | Inserta 50, 30, 70; `Len() == 3`, los tres presentes. |
| **BinaryTree**: Ramas de varios niveles | Sí | `TestBinaryTree/Multilevel_branches` | Inserta 50, 30, 70, 20, 40; `Len() == 5`, todos presentes. |
| **BinaryTree**: Valor ausente en árbol no vacío | Sí | `TestBinaryTree/Missing_value_in_nonempty_tree` | `Contains(99) == false` tras insertar 50. |
| **BinaryTree**: Duplicado idempotente | Sí | `TestBinaryTree/Idempotent_duplicate` | Inserción repetida de 30 no altera `Len() == 3`. |
| **BinaryTree**: Inserciones ordenadas | Sí | `TestBinaryTree/Ordered_insertions` | Inserta 20, 30, 40, 50, 70; verifica presencia y fallo de 99. |
| **Graph**: Crear vacío | Sí | `TestGraph/Create_empty` | `IsEmpty() == true`, contadores en 0, `HasEdge` y `Neighbors` fallan. |
| **Graph**: Vértice aislado | Sí | `TestGraph/Isolated_vertex` | Vértice 10; `VertexCount() == 1`, `EdgeCount() == 0`, `Neighbors(10) == []int{}`. |
| **Graph**: Etiquetas no consecutivas y arista dirigida | Sí | `TestGraph/Nonconsecutive_labels_and_directed_edge` | Arista 10→70 dirigida; 70→10 ausente; 2 vértices, 1 arista. |
| **Graph**: Vértice duplicado conserva adyacencia | Sí | `TestGraph/Duplicate_vertex_preserves_adjacency` | Re-registrar 10 conserva adyacencia previa `[30]`. |
| **Graph**: Arista duplicada | Sí | `TestGraph/Duplicate_edge` | `AddEdge(10,30)` dos veces mantiene `EdgeCount() == 1`. |
| **Graph**: Extremo inexistente sin mutación | Sí | `TestGraph/Missing_endpoint_without_mutation` | Aristas con extremos 98 o 99 fallan (`false`) y no alteran grafo. |
| **Graph**: Vecinos ordenados y contadores | Sí | `TestGraph/Sorted_neighbours_and_counts` | Inserta aristas desordenadas; verifica orden `[30, 50, 70]`. |
| **Graph**: Consulta de arista ausente | Sí | `TestGraph/Missing-edge_query` | Vértices 10 y 30 sin arista devuelven `HasEdge == false`. |
| **Graph**: Vecinos de vértice inexistente | Sí | `TestGraph/Neighbours_of_missing_vertex` | `Neighbors(99) == nil`, distinguiéndose de `[]int{}` de vértice aislado. |
| **Graph**: Bucle y duplicado | Sí | `TestGraph/Self-loop_and_duplicate` | `AddEdge(10,10)` duplicado registra 1 arista; vecinos `[10]`. |
| **Graph**: Aislamiento del resultado de vecinos | Sí | `TestGraph/Neighbour-result_isolation` | Mutar copia devuelta o agregar aristas posteriores no altera slice previo. |

---

## ⚠️ Limitaciones conocidas / Known limitations

| Limitación / Limitation | Impacto / Impact | Alternativa o plan / Workaround or plan |
|---|---|---|
| Búsqueda lineal en grafo `O(V)` | El tiempo de búsqueda de vértices escala con el número de vértices | Requisito pedagógico deliberado del módulo para practicar estructuras enlazadas puras sin tablas hash. |
| Árbol binario sin balanceo automático (no AVL / Rojo-Negro) | Inserciones ordenadas degradan la altura a `O(n)` | Cumple con la especificación de árbol binario de búsqueda básico; estructuras auto-balanceadas corresponden a módulos posteriores. |

---

## 📝 Notas de implementación / Implementation Notes

**ES:** Go gestiona memoria mediante recolección de basura (*garbage collection*), permitiendo enlazar nodos dinámicamente mediante punteros `*TreeNode` y `*GraphNode`. La representación de ausencia en referencias es naturalmente `nil`.

Para `BinaryTree`, las operaciones de inserción y consulta se implementaron de forma iterativa, garantizando un consumo de pila `O(1)` y tiempo `O(h)`. Para `Graph`, la especificación exige construir el registro de vértices y las listas de adyacencia de forma manual sin recurrir a mapas (`map`) ni conjuntos de la biblioteca estándar. La distinción entre un vértice aislado (cuyo método `Neighbors` devuelve un slice vacío `[]int{}`) y un vértice no existente (que devuelve `nil` como indicador de fallo) se modela con fidelidad en Go aprovechando la diferencia entre slice nulo y slice vacío no nulo.

**EN:** Go handles memory using garbage collection, allowing nodes to be linked dynamically through `*TreeNode` and `*GraphNode` pointers. Absence of a link is represented canonically by `nil`.

For `BinaryTree`, insertion and membership operations are implemented iteratively, ensuring `O(1)` stack space and `O(h)` time complexity. For `Graph`, the specification demands building vertex registries and adjacency sequences manually without delegating to standard library hash maps or sets. The distinction between an isolated vertex (whose `Neighbors` returns an empty slice `[]int{}`) and a missing vertex (which returns `nil` as a failure indicator) is represented accurately in Go through the semantic distinction between nil slices and empty non-nil slices.

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

## 📚 Referencias / References

| Tipo / Kind | Referencia / Reference |
|---|---|
| Especificación / Specification | [07_Data_Structures_Advanced.md](../../../../docs/core/algorithms/07_Data_Structures_Advanced.md) |
| Primera base del lenguaje / First language reference | `data_structures_basics` — [data_structures_basics/README.md](../data_structures_basics/README.md) |
| Segunda base del lenguaje / Second language reference | `numbers` — [numbers/README.md](../../foundations/numbers/README.md) |
| Guía de inicialización / Initialisation guide | [00_Project_Initialization_Guide.md](../../../../docs/core/00_Project_Initialization_Guide.md) |
| Adaptaciones idiomáticas / Idiomatic adaptations | [AGENT_Template.md](../../../../docs/AGENT_Template.md) |
| Validación de la documentación / Documentation validation | [WORKFLOW.md](../../../../docs/WORKFLOW.md) |
| Documentación oficial del lenguaje / Language official docs | [The Go Programming Language Specification](https://go.dev/ref/spec) |

---

*[← Volver a Algorithms](../README.md)*

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
