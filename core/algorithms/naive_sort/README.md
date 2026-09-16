# Naive Sort — Go

Implementación de la especificación [05_Naive_Sort](https://yorche3.github.io/programming_languages/core/algorithms/05_Naive_Sort/) en **Go**, con gestión de dependencias mediante **Go Modules** y el framework de pruebas **testify**.

Los tres algoritmos elementales de ordenamiento ($O(n^2)$) — **Selection Sort**, **Bubble Sort** e **Insertion Sort** — trabajan **in-place** sobre el slice de entrada `[]int` y no invocan ninguna biblioteca de ordenamiento del sistema.

---

## 📂 Archivos y estructura / Files & Structure

| Archivo / Directorio | Propósito |
|----------------------|-----------|
| `src/naive_sort.go` | Implementación de los 3 algoritmos del contrato (`SelectionSort`, `BubbleSort`, `InsertionSort`). |
| `tests/naive_sort_test.go` | Pruebas unitarias: 3 tests × (7 casos de la especificación + caso nulo). |
| `go.mod` | Definición del módulo Go y dependencias. |
| `go.sum` | Checksums de las dependencias descargadas. |
| `README.md` | Este archivo. |

Con relación a la implementación base en **Ada**, que separa especificación (`naive_sort.ads`) e implementación (`naive_sort.adb`), Go reúne contrato y código en un único archivo. El motor `go test` descubre automáticamente las funciones con firma `func TestXxx(t *testing.T)`, sin suites ni registros manuales, así que no hace falta el `run_tests` de la especificación.

En **C#** la implementación devuelve `int[]?` porque los *nullable reference types* permiten expresar la entrada inválida; en **Go** el indicador de fallo es el **slice `nil`**, que se propaga sin comprobación explícita. Ambos son los únicos hermanos de la Fase 1 que sí modelan el caso nulo: en **Elm**, **Erlang**, **Forth**, **F#** y **Gleam** el tipo de colección no admite entradas inválidas.

**Estructura de directorios esperada:**

```text
naive_sort/                        # Módulo Go
├── src/
│   └── naive_sort.go              # SelectionSort, BubbleSort, InsertionSort
├── tests/
│   └── naive_sort_test.go         # 3 tests × (7 casos + caso nulo) = 24 aserciones
├── go.mod                         # Módulo y dependencias
├── go.sum                         # Checksums de dependencias
└── README.md                      # Este archivo
```

---

## 🛠️ Enfoque y construcción / Approach & Build

**ES:** El proyecto se creó manualmente, sin herramientas de scaffolding. A diferencia de **Ada** (que usa `alr init --lib`), en Go basta con `go mod init` y escribir el código fuente.

**EN:** The project was created manually, without scaffolding tools. Unlike **Ada** (which uses `alr init --lib`), in Go it's enough to run `go mod init` and write the source code.

### Inicialización / Initialization

```bash
# 1. Inicializar el módulo / Initialize the module
go mod init example.com/naive_sort

# 2. Agregar testify como dependencia / Add testify dependency
go get github.com/stretchr/testify
```

> **ES:** Una vez escritos los archivos fuente, `go mod tidy` sincroniza las dependencias automáticamente.
>
> **EN:** Once source files are written, `go mod tidy` synchronizes dependencies automatically.

---

## 📄 Configuración clave / Key Configuration

### `go.mod` – Definición del módulo

**ES:** Declara el módulo y fija la versión de Go y de `testify`. La versión de `testify` está fijada a `v1.11.1` para coincidir con el módulo hermano `foundations/numbers`, y `go.sum` registra los checksums de las tres dependencias indirectas.

**EN:** Declares the module and pins the Go and `testify` versions. `testify` is pinned to `v1.11.1` to match the sibling module `foundations/numbers`, and `go.sum` records the checksums of the three indirect dependencies.

```go.mod
module example.com/naive_sort

go 1.26.5

require github.com/stretchr/testify v1.11.1

require (
        github.com/davecgh/go-spew v1.1.1 // indirect
        github.com/pmezard/go-difflib v1.0.0 // indirect
        gopkg.in/yaml.v3 v3.0.1 // indirect
)
```

### `.gitignore` – Archivos ignorados

**ES:** El módulo no necesita `.gitignore` propio: el de la raíz del submódulo `go/` ya cubre los binarios de prueba (`*.test`), los perfiles de cobertura (`*.out`, `coverage.*`, `*.coverprofile`) y los binarios compilados (`*.exe`, `*.dll`, `*.so`, `*.dylib`). Verificado con `git check-ignore -v tests/tests.test` → `.gitignore:12:*.test`.

**EN:** The module needs no `.gitignore` of its own: the one at the root of the `go/` submodule already covers test binaries (`*.test`), coverage profiles (`*.out`, `coverage.*`, `*.coverprofile`) and compiled binaries (`*.exe`, `*.dll`, `*.so`, `*.dylib`). Verified with `git check-ignore -v tests/tests.test` → `.gitignore:12:*.test`.

---

## 🚀 Compilación y ejecución / Build & Run

### Compilar / Build

```bash
go build ./...
```

### Analizar / Analyze

```bash
go vet ./...
gofmt -l .
```

### Ejecutar pruebas / Run tests

```bash
go test ./... -v
```

**Salida real / Actual output:**

```text
$ go vet ./...
$ gofmt -l .
$ go build ./...
$ go test ./... -v -count=1
?       example.com/naive_sort/src      [no test files]
=== RUN   TestSelectionSort
--- PASS: TestSelectionSort (0.00s)
=== RUN   TestBubbleSort
--- PASS: TestBubbleSort (0.00s)
=== RUN   TestInsertionSort
--- PASS: TestInsertionSort (0.00s)
PASS
ok      example.com/naive_sort/tests    0.003s
```

> **ES:** `go vet` no reporta hallazgos, `gofmt -l .` no lista ningún archivo y la compilación termina sin salida (exit 0 en los tres casos). `go test` ejecuta **3 tests y 24 aserciones**: 7 casos de la especificación + el caso nulo, por cada uno de los 3 algoritmos.
>
> **EN:** `go vet` reports no findings, `gofmt -l .` lists no files and the build finishes with no output (exit 0 for all three). `go test` runs **3 tests and 24 assertions**: the 7 specification cases + the null case, for each of the 3 algorithms.

---

## 🧠 Algoritmos y operaciones / Algorithms & Operations

| Algoritmo | Estrategia | Complejidad | In-place |
|-----------|------------|-------------|:--------:|
| `SelectionSort` | Busca el mínimo del tramo no ordenado y lo ubica al inicio | $O(n^2)$ siempre | ✅ |
| `BubbleSort` | Compara e intercambia adyacentes, con **salida temprana** mediante la bandera `swapped` | $O(n^2)$ peor/promedio, $O(n)$ mejor | ✅ |
| `InsertionSort` | Inserta cada elemento en el sub-array ya ordenado desplazando los mayores | $O(n^2)$ peor/promedio, $O(n)$ mejor | ✅ |

Los tres algoritmos comparten el mismo contrato `func([]int) []int`, lo que permite que el helper de la suite los recorra con un único bucle.

### Casos cubiertos / Covered cases

| Caso | Entrada | Salida esperada |
|------|---------|-----------------|
| Array estándar desordenado | `[5, 2, 9, 1, 5, 6]` | `[1, 2, 5, 5, 6, 9]` |
| Array ya ordenado | `[1, 2, 3, 4, 5]` | `[1, 2, 3, 4, 5]` |
| Array en orden inverso | `[5, 4, 3, 2, 1]` | `[1, 2, 3, 4, 5]` |
| Elementos idénticos | `[7, 7, 7, 7]` | `[7, 7, 7, 7]` |
| Con números negativos | `[3, -1, 4, -5, 0]` | `[-5, -1, 0, 3, 4]` |
| Un solo elemento | `[42]` | `[42]` |
| Array vacío | `[]` | `[]` |
| **Entrada nula** | `nil` | `nil` |

---

## 📝 Notas de implementación / Implementation Notes

### 🧬 Ordenamiento *in-place* / In-place sorting

**ES:** El pseudocódigo ya contempla ordenar el propio array (`swap(arr, i, j)`), así que los tres algoritmos mutan el slice recibido y lo devuelven, sin asignar ninguna estructura auxiliar. Go pasa los slices por referencia a un array subyacente, de modo que la mutación es visible para el llamador.

Como los *fixtures* de la suite son variables compartidas a nivel de paquete, cada caso ordena una **copia** creada con `copyInput` (`make` + `copy`). Sin ese aislamiento, el primer algoritmo dejaría los datos ordenados y los siguientes pasarían con entradas ya resueltas.

**EN:** The pseudocode already sorts the array itself (`swap(arr, i, j)`), so all three algorithms mutate the received slice and return it, without allocating any auxiliary structure. Go passes slices by reference to an underlying array, so the mutation is visible to the caller.

Since the suite's fixtures are package-level shared variables, each case sorts a **copy** created by `copyInput` (`make` + `copy`). Without that isolation, the first algorithm would leave the data sorted and the following ones would pass with already-solved inputs.

### 🆗 Indicador de fallo con slice `nil` / Failure indicator via `nil` slice

**ES:** La especificación exige devolver el indicador de fallo del lenguaje ante una entrada nula o inválida, sin lanzar excepciones. En Go el slice `nil` es un valor válido y utilizable: `len(nil) == 0`, y la guarda `len(arr) <= 1` devuelve el propio slice sin tocarlo, de modo que `nil` entra y `nil` sale. No hay ninguna comprobación especial ni `panic`.

El caso nulo se prueba de forma explícita y separada del caso vacío, porque en Go `nil` y `[]int{}` son valores distintos: ambos tienen longitud 0, pero sólo `nil` representa la ausencia de array.

**EN:** The specification requires returning the language's failure indicator for a null or invalid input, without throwing exceptions. In Go a `nil` slice is a valid, usable value: `len(nil) == 0`, and the `len(arr) <= 1` guard returns the slice untouched, so `nil` goes in and `nil` comes out. There is no special check and no `panic`.

The null case is tested explicitly and separately from the empty case, because in Go `nil` and `[]int{}` are different values: both have length 0, but only `nil` represents the absence of an array.

### 🏷️ Naming exportado y `package src` / Exported naming and `package src`

**ES:** La especificación nombra las funciones en `snake_case` (`selection_sort`), pero Go sólo exporta identificadores en `PascalCase`; como los tests viven en un paquete distinto (`tests`) y deben poder llamarlas, se usa `SelectionSort`, `BubbleSort` e `InsertionSort`. Además, el paquete se llama `src` y no `naive_sort`, siguiendo el layout del módulo hermano `foundations/numbers`. Los helpers del algoritmo no se extraen: cada función es autocontenida.

**EN:** The specification names the functions in `snake_case` (`selection_sort`), but Go only exports `PascalCase` identifiers; since the tests live in a separate package (`tests`) and must be able to call them, `SelectionSort`, `BubbleSort` and `InsertionSort` are used. The package is also named `src` rather than `naive_sort`, following the layout of the sibling `foundations/numbers` module. Algorithm helpers are not extracted: each function is self-contained.

### 🔁 Salida temprana de `BubbleSort` / `BubbleSort` early exit

**ES:** El criterio de aceptación exige conservar la optimización con bandera de intercambio. La versión Go declara `swapped := true` y usa `for swapped` como condición del bucle externo, de modo que el algoritmo termina en cuanto una pasada completa no realiza ningún intercambio. Es la misma semántica que el `swapped` + `break` del pseudocódigo, sin necesidad de contar pasadas con un índice externo: el caso `[1, 2, 3, 4, 5]` sale en una sola pasada.

**EN:** The acceptance criteria require preserving the swap-flag optimization. The Go version declares `swapped := true` and uses `for swapped` as the outer loop condition, so the algorithm stops as soon as a full pass performs no swaps. This is the same semantics as the pseudocode's `swapped` + `break`, without counting passes with an outer index: the `[1, 2, 3, 4, 5]` case exits in a single pass.

### ➿ Guarda de un solo elemento / Single-element guard

**ES:** Las tres funciones empiezan con `if len(arr) <= 1 { return arr }`, que cubre a la vez los casos de array vacío, array de un solo elemento y entrada nula.

**EN:** All three functions start with `if len(arr) <= 1 { return arr }`, which covers the empty array, the single-element array and the null input at once.

### 🧪 Estructura de los tests / Test structure

**ES:** La suite usa la convención de Go (`func TestXxx(t *testing.T)`) con el framework `testify`:

- **Constantes con nombre** para cada entrada y salida esperada (`standardInput`, `standardOutput`, `reverseInput`, …), sin duplicar literales.
- Una **tabla de casos** (`[]testCase`) con el campo descriptivo, la entrada y la salida esperada.
- Un **helper compartido** `assertSortsAllCases(t, sort, algorithm)` que recibe la función a probar y el nombre del algoritmo, de modo que un único bucle recorre los 7 casos y añade el caso nulo.
- Un **test por función** del contrato: `TestSelectionSort`, `TestBubbleSort` y `TestInsertionSort`.
- El mensaje de cada aserción incluye el nombre del algoritmo y el caso (`"selection_sort should sort an unsorted array"`), por lo que un fallo identifica de inmediato el algoritmo y la entrada responsable.

`testify` ya es dependencia del módulo `foundations/numbers`, así que su uso mantiene la coherencia del submódulo.

**EN:** The suite uses Go's convention (`func TestXxx(t *testing.T)`) with the `testify` framework:

- **Named constants** for every input and expected output (`standardInput`, `standardOutput`, `reverseInput`, …), with no duplicated literals.
- A **case table** (`[]testCase`) with the descriptive field, the input and the expected output.
- A **shared helper** `assertSortsAllCases(t, sort, algorithm)` that receives the function under test and the algorithm name, so a single loop walks the 7 cases and appends the null case.
- One **test per contract function**: `TestSelectionSort`, `TestBubbleSort` and `TestInsertionSort`.
- Each assertion message includes the algorithm name and the case (`"selection_sort should sort an unsorted array"`), so a failure immediately identifies the responsible algorithm and input.

`testify` is already a dependency of the `foundations/numbers` module, so using it keeps the submodule coherent.

### 📍 Desviaciones respecto a la ubicación esperada / Deviations from the expected location

| Especificación | Implementación | Motivo |
|----------------|----------------|--------|
| `src/naive_sort.ext` | `src/naive_sort.go` | La extensión de Go es `.go`; el nombre coincide con el esperado. |
| `test/naive_sort_test.ext` | `tests/naive_sort_test.go` | El módulo hermano `foundations/numbers` usa `tests/` en plural; la convención existente prevalece. |
| `test/run_tests.ext` | *(no existe)* | `go test ./...` descubre automáticamente los tests `TestXxx` de los paquetes del módulo; no se necesita punto de entrada. |

**ES:** El paquete se declara como `src` para que la ruta de importación sea `example.com/naive_sort/src`, igual que en `numbers/`.

**EN:** The package is declared as `src` so that the import path is `example.com/naive_sort/src`, as in `numbers/`.

**ES:** Este proyecto también está implementado en otros lenguajes. Explora el repositorio principal para consultar las demás versiones.

**EN:** This project is also implemented in other languages. Explore the main repository to see the other versions.

---

*[← Volver a Algoritmos Puros](../README.md) · [↑ Volver a Core](../../README.md)*

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
