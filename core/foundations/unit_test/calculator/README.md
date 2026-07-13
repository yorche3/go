# Calculator — Go

Implementación de la especificación [03_Unit_Test_Calculator](https://yorche3.github.io/programming_languages/core/foundations/03_Unit_Test_Calculator/) en **Go**, con gestión de dependencias mediante **Go Modules** y el framework de pruebas **testify**.

---

## 📂 Archivos y estructura / Files & Structure

| Archivo / Directorio | Propósito |
|----------------------|-----------|
| `src/calculator.go` | Implementación de las 5 operaciones aritméticas básicas en el paquete `src`. |
| `tests/calculator_test.go` | Suite de pruebas unitarias con testify (paquete `tests`). |
| `go.mod` | Definición del módulo Go y dependencias. |
| `go.sum` | Checksums de las dependencias descargadas. |

Con relación a la implementación base en **Ada**, que separa interfaz (`calculator.ads`) e implementación (`calculator.adb`) y requiere un subproyecto `tests/` con registro manual de cada caso con `AUnit.Test_Caller`, Go utiliza un **único módulo** sin separación de interfaz/implementación, y las pruebas se colocan en un paquete independiente (`package tests`) en un directorio `tests/`. El framework `testify` descubre automáticamente las funciones que sigan la firma `func TestXxx(t *testing.T)` mediante el motor de pruebas integrado de Go (`go test`), sin necesidad de suites ni registros manuales.

En **C**, el proyecto sigue una estructura con archivos separados (`include/calculator.h`, `src/calculator.c`, `test/calculator_test.c`) y usa Criterion con descubrimiento automático de tests mediante el macro `Test()`. En **C++**, se utiliza Bazel con Google Test y targets declarativos en `BUILD`. Go se alinea más con C++ en cuanto a gestión de dependencias mediante un manifiesto (`go.mod`), pero más con C en la simplicidad del punto de entrada de pruebas (sin necesidad de suites ni runners explícitos).

**Estructura de directorios esperada:**

```text
calculator/                       # Módulo Go
├── src/
│   └── calculator.go             # Implementación de las 5 operaciones
├── tests/
│   └── calculator_test.go        # Pruebas unitarias (5 tests)
├── go.mod                        # Módulo y dependencias
├── go.sum                        # Checksums de dependencias
└── README.md                     # Este archivo
```

---

## 🛠️ Enfoque y construcción / Approach & Build

**ES:** El proyecto se creó manualmente, sin herramientas de scaffolding, para controlar cada detalle. A diferencia de **Ada**, donde se usa `alr init --lib` para crear la biblioteca base y luego `alr init --bin tests` para el subproyecto de pruebas, en Go basta con inicializar un módulo simple con `go mod init` y escribir el código fuente y las pruebas.

El flujo de trabajo en Go es más directo que en Ada: no es necesario un paso de compilación separado ni registro manual de tests en una suite. A diferencia de **C**, que requiere un Makefile y Criterion instalado como dependencia del sistema, Go proporciona su propio gestor de dependencias y ejecutor de pruebas integrados desde el compilador estándar.

**EN:** The project was created manually, without scaffolding tools, to control every detail. Unlike **Ada**, where `alr init --lib` creates the base library and then `alr init --bin tests` creates the test subproject, in Go it's enough to initialize a simple module with `go mod init` and write the source code and tests.

The Go workflow is more straightforward than Ada's: no separate compilation step is needed, and no manual test registration in a suite is required. Unlike **C**, which requires a Makefile and Criterion installed as a system dependency, Go provides its own built-in dependency manager and test runner from the standard compiler.

### Inicialización / Initialization

```bash
# 1. Inicializar el módulo / Initialize the module
go mod init example.com/calculator

# 2. Agregar testify como dependencia / Add testify dependency
go get github.com/stretchr/testify
```

> **ES:** Una vez escritos los archivos fuente, `go mod tidy` sincroniza las dependencias automáticamente.
>
> **EN:** Once source files are written, `go mod tidy` synchronizes dependencies automatically.

---

## 📄 Archivos de configuración clave / Key Configuration Files

### `go.mod` – Definición del módulo

**ES:** Define el módulo Go, su versión y las dependencias externas. Es el equivalente a `alire.toml` en Ada (con dependencias como `aunit`) o `MODULE.bazel` en C++. En contraste con C, que requiere instalar Criterion como biblioteca del sistema, Go descarga y gestiona las dependencias automáticamente.

**EN:** Defines the Go module, its version, and external dependencies. It's the equivalent of `alire.toml` in Ada (with dependencies like `aunit`) or `MODULE.bazel` in C++. Unlike C, which requires installing Criterion as a system library, Go downloads and manages dependencies automatically.

```go.mod
module example.com/calculator

go 1.26.5

require (
	github.com/stretchr/testify v1.11.1
)
```

### `go.sum` – Checksums de dependencias

**ES:** Almacena los checksums de las dependencias para garantizar integridad y reproducibilidad. Es generado automáticamente por Go.

**EN:** Stores dependency checksums to ensure integrity and reproducibility. It's automatically generated by Go.

### `.gitignore` – Archivos ignorados

**ES:** Patrones para no versionar archivos generados (binarios, dependencias, etc.). A diferencia de Ada (que ignora `alire/`, `obj/`, `lib/`, `bin/`, `config/`), C (que ignora `obj/`, `bin/`) y C++ (que ignora los directorios de Bazel), Go genera pocos artefactos: binarios compilados con `go build` y archivos de cobertura.

**EN:** Patterns to avoid versioning generated files (binaries, dependencies, etc.). Unlike Ada (which ignores `alire/`, `obj/`, `lib/`, `bin/`, `config/`), C (which ignores `obj/`, `bin/`), and C++ (which ignores Bazel directories), Go generates few artifacts: binaries compiled with `go build` and coverage files.

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
=== RUN   TestAddition
--- PASS: TestAddition (0.00s)
=== RUN   TestSubtraction
--- PASS: TestSubtraction (0.00s)
=== RUN   TestMultiplication
--- PASS: TestMultiplication (0.00s)
=== RUN   TestDivision
--- PASS: TestDivision (0.00s)
=== RUN   TestModulus
--- PASS: TestModulus (0.00s)
PASS
ok      example.com/calculator/tests    0.123s
```

> **ES:** A diferencia de Ada (que usa `alr -C tests run` y muestra un resumen con "Submitted: 5 test case(s)"), C (que usa `make test` y muestra "Synthesis: Tested: 5 | Passing: 5 | Failing: 0") y C++ (que usa `bazelisk test //...`), Go integra el ejecutor de pruebas en el propio compilador sin necesidad de herramientas externas.

---

## 🧠 Algoritmos / operaciones (según el módulo)

| Función / Function | Implementación / Implementation | Cumple / Complies |
|--------------------|--------------------------------|-------------------|
| `Addition(a, b)` | `a + b` (suma directa / direct addition) | ✅ |
| `Subtraction(a, b)` | `a - b` (resta directa / direct subtraction) | ✅ |
| `Multiplication(a, b)` | Suma repetitiva llamando a `Addition()` / Repeated addition calling `Addition()` | ✅ No usa `*` |
| `Division(a, b)` | Resta repetitiva llamando a `Subtraction()` y `Addition()` / Repeated subtraction | ✅ No usa `/` |
| `Modulus(a, b)` | `Subtraction(a, Multiplication(b, Division(a, b)))` | ✅ No usa `%` |

> **ES:** A diferencia de la especificación que sugiere el pseudocódigo con bucles básicos, Go permite nombrar las operaciones básicas (`Addition`, `Subtraction`) y reutilizarlas dentro de las operaciones compuestas (`Multiplication`, `Division`, `Modulus`), manteniendo una coherencia similar a Ada, donde `Multiplication` llama a `Addition(Result, A)` y `Division` llama a `Subtraction(New_Dividend, B)`.

---

## 📝 Notas de implementación / Implementation Notes

### 🔁 Sobre el uso de testify (framework de aserciones) / On using testify (assertion framework)

**ES:**
La especificación de este proyecto indica que debe usarse únicamente la biblioteca estándar del lenguaje. Sin embargo, Go proporciona en su stdlib el paquete `testing` para la estructura de pruebas, pero no incluye un framework de aserciones con mensajes descriptivos. El uso de `testify` como dependencia externa es una decisión deliberada para enriquecer la legibilidad y el diagnóstico de fallos, siguiendo el mismo patrón que:

- **Ada** usa **AUnit** como dependencia externa (vía `alr with aunit`).
- **C** usa **Criterion** como dependencia externa (vía `pkg-config`).
- **C++** usa **Google Test** como dependencia externa (vía Bazel `MODULE.bazel`).

Si se desea cumplir estrictamente con la biblioteca estándar, se puede utilizar únicamente el método `t.Errorf` o `t.Fatal` del paquete `testing`.

**EN:**
The specification for this project indicates that only the language's standard library should be used. However, Go provides the `testing` package in its stdlib for test structure, but does not include an assertion framework with descriptive messages. Using `testify` as an external dependency is a deliberate decision to enhance readability and failure diagnosis, following the same pattern as:

- **Ada** uses **AUnit** as an external dependency (via `alr with aunit`).
- **C** uses **Criterion** as an external dependency (via `pkg-config`).
- **C++** uses **Google Test** as an external dependency (via Bazel `MODULE.bazel`).

To strictly comply with the standard library, only the `t.Errorf` or `t.Fatal` methods from the `testing` package can be used.

### 🔄 Comparación con Ada, C y C++

**ES:**

| Aspecto | Ada | C | C++ | Go |
|---------|-----|---|---|----|
| **Gestor de paquetes** | Alire (`alire.toml`) | Sistema (`apt`/`brew`) + Makefile | Bazelisk (`MODULE.bazel`) | Go Modules (`go.mod`) |
| **Framework de tests** | AUnit | Criterion | Google Test | testify |
| **Separación interfaz/impl** | Sí (`.ads` / `.adb`) | Sí (`.h` / `.c`) | Sí (`.h` / `.cpp`) | No (todo en `.go`) |
| **Registro de tests** | Manual (suite + caller) | Automático (macro `Test()`) | Automático (macro `TEST()`) | Automático (convención `TestXxx`) |
| **Punto de entrada** | Manual (`tests.adb`) | Automático (Criterion) | Automático (Google Test) | Automático (`go test`) |
| **Compilación** | `alr build` | `make` | `bazelisk build //...` | `go build ./...` |
| **Ejecución de tests** | `alr -C tests run` | `make test` | `bazelisk test //...` | `go test ./...` |
| **Llamadas a ops. básicas** | Sí (Addition/Subtraction) | Sí (addition/subtraction) | Sí (Addition/Subtraction) | Sí (Addition/Subtraction) |

**EN:**

| Aspect | Ada | C | C++ | Go |
|--------|-----|---|---|----|
| **Package manager** | Alire (`alire.toml`) | System (`apt`/`brew`) + Makefile | Bazelisk (`MODULE.bazel`) | Go Modules (`go.mod`) |
| **Test framework** | AUnit | Criterion | Google Test | testify |
| **Interface/impl separation** | Yes (`.ads` / `.adb`) | Yes (`.h` / `.c`) | Yes (`.h` / `.cpp`) | No (all in `.go`) |
| **Test registration** | Manual (suite + caller) | Automatic (`Test()` macro) | Automatic (`TEST()` macro) | Automatic (`TestXxx` convention) |
| **Entry point** | Manual (`tests.adb`) | Automatic (Criterion) | Automatic (Google Test) | Automatic (`go test`) |
| **Build** | `alr build` | `make` | `bazelisk build //...` | `go build ./...` |
| **Test execution** | `alr -C tests run` | `make test` | `bazelisk test //...` | `go test ./...` |
| **Basic op calls** | Yes (Addition/Subtraction) | Yes (addition/subtraction) | Yes (Addition/Subtraction) | Yes (Addition/Subtraction) |

### 💡 Multiplicación, división y módulo sin operadores nativos

**ES:**
`Multiplication` y `Division` se implementan con sumas/restas repetitivas para cumplir la especificación educativa (sin operadores `*`, `/` ni `%` directos). Go, al igual que Ada, C y C++, permite bucles `for`.

Al igual que **Ada** (que llama explícitamente a `Addition(Result, A)` y `Subtraction(New_Dividend, B)`), la implementación en Go mantiene coherencia llamando a `Addition(result, a)` y `Subtraction(a, b)` dentro de `Multiplication` y `Division`. Esto contrasta con versiones alternativas que usan operadores compuestos (`+=`, `-=`). `Modulus` se define en términos de `Division` y `Multiplication`, siguiendo el mismo patrón que Ada y la especificación original.

A diferencia de **Ada**, que usa `while New_Dividend >= B loop`, Go usa `for a >= b { ... }`. También a diferencia de Ada, Go no requiere un `return` explícito en cada rama — el lenguaje trata la última expresión como valor de retorno, pero aquí se usa `return` explícito para claridad.

### ♻️ Comparación con lenguajes similares alfabéticamente anteriores / Comparison with alphabetically prior similar languages

**ES:**

**Ada** (el lenguaje base): Utiliza el gestor de paquetes **Alire**, separa interfaz (`calculator.ads`) de implementación (`calculator.adb`), requiere un subproyecto `tests/` independiente con su propio `tests.gpr`, y las pruebas se registran manualmente en una suite utilizando `AUnit.Test_Caller`. La creación del proyecto sigue el patrón `alr init --lib` + `alr init --bin tests`. Las operaciones `Multiplication` y `Division` llaman explícitamente a `Addition` y `Subtraction`, y los bucles usan sintaxis `in 1 .. B` y `while ... loop`.

**C**: Utiliza un **Makefile** artesanal como sistema de construcción, separa header (`calculator.h`), implementación (`calculator.c`) y tests (`calculator_test.c`). Las pruebas se descubren automáticamente con el macro `Test()` de Criterion. Requiere instalar Criterion como biblioteca del sistema operativo. Al igual que Go, llama a `addition()` y `subtraction()` dentro de `multiplication()` y `division()`.

**C++**: Utiliza **Bazelisk** (Bazel) como sistema de construcción moderno con targets declarativos en `BUILD`. Separa header (`calculator.h`), implementación (`calculator.cpp`) y tests (`calculator_tests.cpp`). Google Test se integra vía Bzlmod en `MODULE.bazel`.

**Go**: Se diferencia de los tres anteriores al no separar interfaz de implementación (un único archivo `.go` con funciones exportadas). Su gestor de dependencias **Go Modules** (`go.mod`) es más cercano conceptualmente a `alire.toml` (Ada) y `MODULE.bazel` (C++) que al Makefile de C. El descubrimiento de tests es automático por convención de nombres (`TestXxx`), similar a C y C++, y no requiere registro manual como en Ada. La implementación sigue el mismo patrón educacional que Ada (llamando a operaciones básicas dentro de las compuestas).

---

### 🌐 Otras implementaciones / Other implementations

Este proyecto también está implementado en otros lenguajes. Explora el [repositorio principal](https://github.com/yorche3/programming_languages) para ver todas las versiones.

---

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
