# Numbers — Go

Implementación de la especificación [04_Numbers](https://yorche3.github.io/programming_languages/core/foundations/04_Numbers/) en **Go**, con gestión de dependencias mediante **Go Modules** y el framework de pruebas **testify**.

Tres enfoques de implementación para los mismos 5 algoritmos: **recursivo directo**, **recursivo con acumulador** e **iterativo**.

---

## 📂 Archivos y estructura / Files & Structure

| Archivo / Directorio | Propósito |
|----------------------|-----------|
| `src/numbers.go` | Implementación de los 5 algoritmos en 3 enfoques (15 funciones públicas + 3 helpers privados). |
| `tests/numbers_rec_test.go` | Pruebas unitarias para el enfoque recursivo directo (`Rec`). |
| `tests/numbers_iter_test.go` | Pruebas unitarias para el enfoque iterativo (`Iter`). |
| `go.mod` | Definición del módulo Go y dependencias. |
| `go.sum` | Checksums de las dependencias descargadas. |

Con relación a la implementación base en **Ada**, que separa especificación (`numbers.ads`) e implementación (`numbers.adb`) y requiere un subproyecto `tests/` completo con 3 suites (una por enfoque), cada una con su propia especificación y cuerpo, además de un punto de entrada `tests.adb` que las ejecuta todas, Go utiliza un **único módulo** con un solo archivo de código fuente y dos archivos de pruebas. El motor `go test` descubre automáticamente las funciones con firma `func TestXxx(t *testing.T)`, sin necesidad de suites ni registros manuales.

En **C**, el proyecto sigue una estructura con header (`numbers.h`), implementación (`numbers.c`) y tests separados (`numbers_rec_test.c`, `numbers_ite_test.c`). Los helpers `_help` se declaran `static` para mantenerlos privados. En **C++**, se usa Bazel con Google Test, separando también en header (`numbers.h`), implementación (`numbers.cpp`) y dos archivos de tests. Go logra la misma separación de responsabilidades pero con menos archivos de configuración, ya que no requiere ni Makefile (C), ni proyecto GPRbuild (Ada), ni BUILD + MODULE.bazel (C++).

**Estructura de directorios esperada:**

```text
numbers/                          # Módulo Go
├── src/
│   └── numbers.go                # 15 funciones públicas + 3 helpers privados
├── tests/
│   ├── numbers_rec_test.go       # Tests: enfoque recursivo directo (5)
│   └── numbers_iter_test.go      # Tests: enfoque iterativo (5)
├── go.mod                        # Módulo y dependencias
├── go.sum                        # Checksums de dependencias
└── README.md                     # Este archivo
```

---

## 🛠️ Enfoque y construcción / Approach & Build

**ES:** El proyecto se creó manualmente, sin herramientas de scaffolding. A diferencia de **Ada** (que usa `alr init --lib` + `alr init --bin tests`), en Go basta con `go mod init` y escribir el código fuente.

**EN:** The project was created manually, without scaffolding tools. Unlike **Ada** (which uses `alr init --lib` + `alr init --bin tests`), in Go it's enough to run `go mod init` and write the source code.

### Inicialización / Initialization

```bash
# 1. Inicializar el módulo / Initialize the module
go mod init example.com/numbers

# 2. Agregar testify como dependencia / Add testify dependency
go get github.com/stretchr/testify
```

> **ES:** Una vez escritos los archivos fuente, `go mod tidy` sincroniza las dependencias automáticamente.
>
> **EN:** Once source files are written, `go mod tidy` synchronizes dependencies automatically.

---

## 📄 Archivos de configuración clave / Key Configuration Files

### `go.mod` – Definición del módulo

**ES:** Es el equivalente a `alire.toml` (Ada), `MODULE.bazel` (C++) o al Makefile (C). Go descarga y gestiona las dependencias automáticamente.

**EN:** Equivalent to `alire.toml` (Ada), `MODULE.bazel` (C++) or the Makefile (C). Go downloads and manages dependencies automatically.

```go.mod
module example.com/numbers

go 1.26.5

require (
	github.com/stretchr/testify v1.11.1
)
```

### `.gitignore` – Archivos ignorados

**ES:** Go genera pocos artefactos: binarios compilados con `go build` y archivos de cobertura. A diferencia de Ada (que ignora `alire/`, `obj/`, `lib/`, `bin/`, `config/`), C (que ignora `obj/`, `bin/`) y C++ (que ignora los directorios de Bazel), la configuración es mínima.

**EN:** Go generates few artifacts: binaries compiled with `go build` and coverage files. Unlike Ada (which ignores `alire/`, `obj/`, `lib/`, `bin/`, `config/`), C (which ignores `obj/`, `bin/`), and C++ (which ignores Bazel directories), the configuration is minimal.

```gitignore
# Binaries
*.exe
*.test
*.out
```

---

## 🚀 Compilación y ejecución / Build & Run

### Compilar / Build

```bash
go build ./...
```

### Ejecutar pruebas / Run tests

```bash
go test ./tests/ -v
```

> **ES:** La bandera `-v` (verbose) muestra el nombre y resultado de cada test individualmente. Sin ella, solo se muestra el resumen.
>
> **EN:** The `-v` (verbose) flag displays the name and result of each test individually. Without it, only the summary is shown.

**Salida esperada / Expected output:**

```
=== RUN   TestSumFirstNRec
--- PASS: TestSumFirstNRec (0.00s)
=== RUN   TestFactorialRec
--- PASS: TestFactorialRec (0.00s)
=== RUN   TestFibonacciRec
--- PASS: TestFibonacciRec (0.00s)
=== RUN   TestGreatestCommonDivisorRec
--- PASS: TestGreatestCommonDivisorRec (0.00s)
=== RUN   TestLeastCommonMultipleRec
--- PASS: TestLeastCommonMultipleRec (0.00s)
=== RUN   TestSumFirstNIter
--- PASS: TestSumFirstNIter (0.00s)
=== RUN   TestFactorialIter
--- FAIL: TestFactorialIter (0.00s)
=== RUN   TestFibonacciIter
--- PASS: TestFibonacciIter (0.00s)
=== RUN   TestGreatestCommonDivisorIter
--- PASS: TestGreatestCommonDivisorIter (0.00s)
=== RUN   TestLeastCommonMultipleIter
--- FAIL: TestLeastCommonMultipleIter (0.00s)
...
```

> **ES:** Actualmente 2 tests fallan debido a errores en los archivos de prueba (ver sección de notas). La implementación del código fuente es correcta.
>
> **EN:** Currently 2 tests fail due to errors in the test files (see notes section). The source code implementation is correct.

---

## 🧠 Algoritmos / operaciones (según el módulo)

### 3 enfoques × 5 algoritmos = 15 funciones / 10 tests directos

| Algoritmo | Casos de prueba | `Rec` | `Acc` | `Iter` |
|-----------|----------------|:-----:|:-----:|:------:|
| `SumFirstN` | `(0) = 0`, `(3) = 6` | ✅ | ✅¹ | ✅ |
| `Factorial` | `(0) = 1`, `(4) = 24` | ✅ | ✅¹ | ✅² |
| `Fibonacci` | `(0) = 0`, `(1) = 1`, `(6) = 8` | ✅ | ✅¹ | ✅ |
| `GreatestCommonDivisor` | `(12, 8) = 4`, `(7, 5) = 1` | ✅ | ✅¹ | ✅ |
| `LeastCommonMultiple` | `(4, 6) = 12`, `(6, 8) = 24` | ✅ | ✅¹ | ✅³ |

> ¹ Los acumuladores se prueban implícitamente al ejecutar `Acc(n)`, que internamente llama al helper privado. No tienen tests directos (igual que en C y C++).
> ² El test espera `FactorialIter(0) = 0` cuando el valor real es `1`.
> ³ El test llama a `GreatestCommonDivisorIter` en lugar de `LeastCommonMultipleIter`.

---

## 📝 Notas de implementación / Implementation Notes

### 🔁 Sobre recursión con acumulador y Tail Call Optimization (TCO) / On recursion with accumulator and Tail Call Optimization (TCO)

**ES:**

Tail recursion ocurre cuando la llamada recursiva es la última acción que ejecuta una función/método; después de la llamada no hay más instrucciones, la función devuelve el resultado de la llamada recursiva. La recursión con acumulador consigue esto pasando el estado previo como parámetro a cada llamada, sin dejar trabajo pendiente en la pila.

**Go no garantiza TCO.** Aunque el compilador de Go (gc) aplica ciertas optimizaciones, no implementa Tail Call Optimization de forma explícita. La especificación del lenguaje no exige TCO. Por lo tanto, las funciones con acumulador (`Acc`) consumen la misma pila que las recursivas directas.

La implementación con acumulador se conserva únicamente con fines educativos: sirve como puente conceptual entre la recursión directa (más cercana a la definición matemática) y la versión iterativa (más eficiente). Como en este contexto no hay un beneficio práctico de rendimiento, no se desarrollan pruebas unitarias específicas para los métodos con acumulador. La validación del comportamiento se cubre a través de las pruebas de los enfoques recursivo e iterativo, que juntos ejercitan los mismos resultados.

Este comportamiento es **idéntico al de C y C++**, que tampoco garantizan TCO y también conservan los acumuladores solo con fines educativos. En contraste, **Ada** con GNAT puede aplicar TCO con niveles de optimización adecuados, por lo que sus tests con acumulador tienen cobertura directa.

**EN:**

Tail recursion occurs when the recursive call is the last action that runs a function/method; after the call there are no more instructions, the function returns the result of the recursive call. Recursion with accumulator achieves this by passing the previous state as a parameter to each call, without leaving any pending work on the stack.

**Go does not guarantee TCO.** Although the Go compiler (gc) applies certain optimizations, it does not explicitly implement Tail Call Optimization. The language specification does not require TCO. Therefore, accumulator functions (`Acc`) consume the same stack as direct recursion.

The accumulator implementation is preserved only for educational purposes: it serves as a conceptual bridge between the direct recursive (closer to mathematical definition) and the iterative version (more efficient). Since there is no practical performance benefit, no specific unit tests are developed for the recursive methods with accumulator. The behavior validation is covered through the tests of recursive and iterative approaches, which together exercise the same results.

This behavior is **identical to C and C++**, which also do not guarantee TCO and also preserve accumulators for educational purposes only. In contrast, **Ada** with GNAT can apply TCO with appropriate optimization levels, so its accumulator tests have direct coverage.

### 🆇 Errores conocidos en los tests / Known issues in tests

**ES:**
Se han identificado **2 errores** en los archivos de prueba que provocan fallos en la ejecución:

1. **`tests/numbers_iter_test.go` — `TestFactorialIter`**: Espera `FactorialIter(0) = 0`, pero el valor correcto es `1` (0! = 1 por definición matemática). La implementación en `src/numbers.go` retorna `1` correctamente.

   ```go
   // ❌ Actual (falla):
   assert.Equal(t, 0, FactorialIter(0), "Should be equal to 1")
   // ✅ Debería ser:
   assert.Equal(t, 1, FactorialIter(0), "Should be equal to 1")
   ```

2. **`tests/numbers_iter_test.go` — `TestLeastCommonMultipleIter` y `tests/numbers_rec_test.go` — `TestLeastCommonMultipleRec`**: Ambos tests llaman a `GreatestCommonDivisorIter`/`GreatestCommonDivisorRec` en lugar de `LeastCommonMultipleIter`/`LeastCommonMultipleRec`.

   ```go
   // ❌ Actual (prueba GCD, no LCM):
   assert.Equal(t, 24, GreatestCommonDivisorIter(6, 8), "Should be equal to 24")
   // ✅ Debería ser:
   assert.Equal(t, 24, LeastCommonMultipleIter(6, 8), "Should be equal to 24")
   ```

Estos errores están documentados aquí para su corrección. La implementación del código fuente (`src/numbers.go`) es correcta en todos los casos.

**EN:**
**2 errors** have been identified in the test files that cause failures during execution:

1. **`tests/numbers_iter_test.go` — `TestFactorialIter`**: Expects `FactorialIter(0) = 0`, but the correct value is `1` (0! = 1 by mathematical definition). The implementation in `src/numbers.go` correctly returns `1`.

2. **`tests/numbers_iter_test.go` — `TestLeastCommonMultipleIter` and `tests/numbers_rec_test.go` — `TestLeastCommonMultipleRec`**: Both tests call `GreatestCommonDivisorIter`/`GreatestCommonDivisorRec` instead of `LeastCommonMultipleIter`/`LeastCommonMultipleRec`.

These errors are documented here for correction. The source code implementation (`src/numbers.go`) is correct in all cases.

### 🔄 Comparación con Ada, C y C++

**ES:**

| Aspecto | Ada | C | C++ | Go |
|---------|-----|---|---|----|
| **Gestor de paquetes** | Alire (`alire.toml`) | Sistema + Makefile | Bazelisk (`MODULE.bazel`) | Go Modules (`go.mod`) |
| **Framework de tests** | AUnit | Criterion | Google Test | testify |
| **Separación interfaz/impl** | Sí (`.ads` / `.adb`) | Sí (`.h` / `.c`) | Sí (`.h` / `.cpp`) | No (todo en `.go`) |
| **Funciones privadas** | `private` en package | `static` en `.c` | `private:` en clase | minúscula inicial |
| **Registro de tests** | Manual (suite + caller) | Automático (`Test()` macro) | Automático (`TEST()` macro) | Automático (`TestXxx`) |
| **Tests para `_Acc`** | ✅ Directos (5 tests) | ❌ Implícitos | ❌ Implícitos | ❌ Implícitos |
| **TCO garantizado** | Parcial (GNAT -O2) | ❌ No | ❌ No | ❌ No |
| **Ejecución de tests** | `alr -C tests run` | `make test` | `bazelisk test //...` | `go test ./...` |

**EN:**

| Aspect | Ada | C | C++ | Go |
|--------|-----|---|---|----|
| **Package manager** | Alire (`alire.toml`) | System + Makefile | Bazelisk (`MODULE.bazel`) | Go Modules (`go.mod`) |
| **Test framework** | AUnit | Criterion | Google Test | testify |
| **Interface/impl separation** | Yes (`.ads` / `.adb`) | Yes (`.h` / `.c`) | Yes (`.h` / `.cpp`) | No (all in `.go`) |
| **Private functions** | `private` in package | `static` in `.c` | `private:` in class | lowercase initial |
| **Test registration** | Manual (suite + caller) | Automatic (`Test()` macro) | Automatic (`TEST()` macro) | Automatic (`TestXxx`) |
| **Tests for `_Acc`** | ✅ Direct (5 tests) | ❌ Implicit | ❌ Implicit | ❌ Implicit |
| **TCO guaranteed** | Partial (GNAT -O2) | ❌ No | ❌ No | ❌ No |
| **Test execution** | `alr -C tests run` | `make test` | `bazelisk test //...` | `go test ./...` |

### 💡 Privacidad de helpers en Go

**ES:**
En Go, la visibilidad se controla por la primera letra del nombre:
- **Mayúscula**: función exportada (pública), visible fuera del paquete.
- **Minúscula**: función privada, visible solo dentro del mismo paquete.

Las funciones helper `sumFirstHelper`, `factorialHelper` y `fibonacciHelper` comienzan con minúscula, por lo que son privadas al paquete `src`. Esto es análogo a:
- **Ada**: funciones declaradas en el cuerpo (`.adb`) sin estar en la especificación (`.ads`).
- **C**: funciones marcadas como `static` en el archivo `.c`.
- **C++**: métodos marcados como `private:` en la clase.

A diferencia de C (que requiere `static` explícito) y C++ (que requiere la sección `private:`), en Go la convención de nomenclatura es la única regla, lo que simplifica la sintaxis.

**EN:**
In Go, visibility is controlled by the first letter of the name:
- **Uppercase**: exported function (public), visible outside the package.
- **Lowercase**: private function, visible only within the same package.

The helper functions `sumFirstHelper`, `factorialHelper` and `fibonacciHelper` start with lowercase, so they are private to the `src` package. This is analogous to:
- **Ada**: functions declared in the body (`.adb`) without being in the specification (`.ads`).
- **C**: functions marked as `static` in the `.c` file.
- **C++**: methods marked as `private:` in the class.

Unlike C (which requires explicit `static`) and C++ (which requires the `private:` section), in Go the naming convention is the only rule, which simplifies the syntax.

---

### 🌐 Otras implementaciones / Other implementations

Este proyecto también está implementado en otros lenguajes. Explora el [repositorio principal](https://github.com/yorche3/programming_languages) para ver todas las versiones.

---

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
