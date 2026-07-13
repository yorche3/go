# 🚀 Fundamentos / Foundations — Go

Implementación de los ejercicios de la sección [Fundamentos / Foundations](https://yorche3.github.io/programming_languages/core/foundations/) del repositorio principal en **Go**.

---

## 📖 Descripción / Description

**ES:** Esta sección reúne los conceptos esenciales para empezar a trabajar con **Go**. Cubre desde el programa más básico (`Hello, World!`) hasta la implementación de una calculadora con pruebas unitarias y algoritmos numéricos en tres enfoques progresivos (recursivo directo, recursivo con acumulador e iterativo).

**EN:** This section brings together the essential concepts to start working with **Go**. It covers everything from the most basic program (`Hello, World!`) to the implementation of a calculator with unit tests and numerical algorithms in three progressive approaches (direct recursion, accumulator recursion, and iterative).

---

## 📁 Estructura / Structure

```text
go/
└── core/
    └── foundations/
        ├── README.md              # Este archivo / This file
        ├── helloworld/            # 01_Hello_World — Primer programa
        │   └── hello_world.go
        ├── hellouser/             # 02_Hello_User — Entrada y salida
        │   └── hello_user.go
        ├── unit_test/
        │   └── calculator/        # 03_Unit_Test_Calculator — Pruebas unitarias
        │       ├── src/
        │       │   └── calculator.go
        │       ├── tests/
        │       │   └── calculator_test.go
        │       ├── go.mod
        │       └── go.sum
        └── numbers/               # 04_Numbers — Algoritmos numéricos
            ├── src/
            │   └── numbers.go
            ├── tests/
            │   ├── numbers_rec_test.go
            │   └── numbers_iter_test.go
            ├── go.mod
            └── go.sum
```

---

## 🔢 Progresión / Progression

| Especificación | Proyecto | Conceptos | Dependencias externas |
|----------------|----------|-----------|:---------------------:|
| [`01_Hello_World`](https://yorche3.github.io/programming_languages/core/foundations/01_Hello_World/) | [`helloworld/`](helloworld/) | `package main`, `func main()`, `fmt.Printf` | ❌ Solo stdlib |
| [`02_Hello_User`](https://yorche3.github.io/programming_languages/core/foundations/02_Hello_User/) | [`hellouser/`](hellouser/) | Variables (`var`), `fmt.Scanf`, `fmt.Sprintf`, `fmt.Println` | ❌ Solo stdlib |
| [`03_Unit_Test_Calculator`](https://yorche3.github.io/programming_languages/core/foundations/03_Unit_Test_Calculator/) | [`unit_test/calculator/`](unit_test/calculator/) | Paquetes, pruebas unitarias, testify, `go mod init` | ✅ testify |
| [`04_Numbers`](https://yorche3.github.io/programming_languages/core/foundations/04_Numbers/) | [`numbers/`](numbers/) | Recursión, iteración, acumuladores, helpers privados, TCO | ✅ testify |

---

## 🛠️ Enfoque general / General Approach

**ES:** Los proyectos en esta sección siguen un patrón progresivo:

1. **Hello World** y **Hello User**: Programas monofichero sin `go.mod`, ejecutables con `go run`. Usan exclusivamente la biblioteca estándar (`fmt`).
2. **Calculator**: Primer proyecto con dependencia externa (`testify`). Introduce el uso de **Go Modules** (`go mod init`), separación `src/` + `tests/`, y el ecosistema de `go test`.
3. **Numbers**: Expande el patrón de Calculator a múltiples archivos de prueba (uno por enfoque) y más funciones. Introduce helpers privados (convención de minúscula inicial).

**EN:** The projects in this section follow a progressive pattern:

1. **Hello World** and **Hello User**: Single-file programs without `go.mod`, executable with `go run`. Use only the standard library (`fmt`).
2. **Calculator**: First project with an external dependency (`testify`). Introduces **Go Modules** (`go mod init`), `src/` + `tests/` separation, and the `go test` ecosystem.
3. **Numbers**: Expands the Calculator pattern to multiple test files (one per approach) and more functions. Introduces private helpers (lowercase initial convention).

---

## 🚀 Ejecución rápida / Quick Start

### Hello World

```bash
cd go/core/foundations/helloworld
go run hello_world.go
```

### Hello User

```bash
cd go/core/foundations/hellouser
go run hello_user.go
```

### Calculator (pruebas)

```bash
cd go/core/foundations/unit_test/calculator
go test ./tests/ -v
```

### Numbers (pruebas)

```bash
cd go/core/foundations/numbers
go test ./tests/ -v
```

---

## 📚 Especificaciones / Specifications

| # | Especificación | Proyecto | Estado |
|:-:|---------------|----------|:------:|
| 01 | [Hello World](https://yorche3.github.io/programming_languages/core/foundations/01_Hello_World/) | [`helloworld/`](helloworld/) | ✅ Completado |
| 02 | [Hello User](https://yorche3.github.io/programming_languages/core/foundations/02_Hello_User/) | [`hellouser/`](hellouser/) | ✅ Completado |
| 03 | [Unit Test Calculator](https://yorche3.github.io/programming_languages/core/foundations/03_Unit_Test_Calculator/) | [`unit_test/calculator/`](unit_test/calculator/) | ✅ Completado |
| 04 | [Numbers](https://yorche3.github.io/programming_languages/core/foundations/04_Numbers/) | [`numbers/`](numbers/) | ✅ Completado |

---

## ▶️ Siguiente / Next

👉 Después de completar los fundamentos, continúa explorando otros módulos en [`go/core/`](../).  
👉 After completing the foundations, continue exploring other modules in [`go/core/`](../).

---

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
