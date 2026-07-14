# Go

Proyectos en **Go**, con gestión de dependencias mediante **Go Modules** y el framework de pruebas **testify**.

---

## 📂 Módulos / Modules

| Módulo | Descripción |
|--------|-------------|
| [`core/foundations/`](core/foundations/) | **Fase 0 — Fundamentos**: `hello_world`, `hello_user`, `calculator`, `numbers` |

---

## ▶️ Comenzar / Getting Started

```bash
# Hello, World!
cd core/foundations/helloworld
go run hello_world.go

# Hello, User!
cd core/foundations/hellouser
go run hello_user.go

# Calculator Tests
cd core/foundations/unit_test/calculator
go test ./tests/ -v

# Numbers Tests
cd core/foundations/numbers
go test ./tests/ -v
```

---

## 📦 Requisitos / Requirements

| Herramienta | Instalación |
|-------------|-------------|
| [Go](https://go.dev/dl/) | `sudo apt install golang-go` (Linux) / `winget install GoLang.Go` (Windows) / [Descargar](https://go.dev/dl/) |

```bash
# Verificar instalación
go version
```

---

## 🏗️ Tipos de proyecto / Project Types

### 1. Programa simple (sin módulo)

**ES:** Un único archivo fuente, sin dependencias externas, ejecutable directamente con `go run`. Ideal para `hello_world` y `hello_user`. No requiere `go.mod`.

**EN:** A single source file, no external dependencies, executable directly with `go run`. Ideal for `hello_world` and `hello_user`. No `go.mod` required.

```bash
go run <file>.go
```

### 2. Proyecto con pruebas unitarias (Go Modules + testify)

**ES:** Para proyectos que requieren pruebas unitarias, se inicializa un módulo con `go mod init` y se agrega `testify` como dependencia con `go get`. El código fuente se organiza en `src/` y las pruebas en `tests/`, con descubrimiento automático de tests mediante la convención `TestXxx`.

**EN:** For projects that require unit tests, initialize a module with `go mod init` and add `testify` as a dependency with `go get`. Source code goes in `src/` and tests in `tests/`, with automatic test discovery via the `TestXxx` convention.

```bash
go mod init example.com/<name>
go get github.com/stretchr/testify
go test ./tests/ -v
```

---

## 🌐 Otras implementaciones / Other implementations

Este proyecto también está implementado en otros lenguajes. Explora el [repositorio principal](https://github.com/yorche3/programming_languages) para ver todas las versiones.

---

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
