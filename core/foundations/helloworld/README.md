# Hello, World! — Go

Implementación de la especificación [01_Hello_World](https://yorche3.github.io/programming_languages/core/foundations/01_Hello_World/) en **Go**, con un enfoque **manual y minimalista** (sin `go mod init` de scaffolding).

---

## 📂 Archivos y estructura / Files & Structure

| Archivo / File | Propósito / Purpose |
|----------------|---------------------|
| [`hello_world.go`](hello_world.go) | Código fuente principal: imprime `"Hello, World! from Go!"` en la consola. |
| [`.gitignore`](../../../.gitignore) | Patrones globales para Go: binarios, binarios de test, perfiles de cobertura, etc. |

> **ES:** Go no requiere un `go.mod` para programas monofichero usando solo la biblioteca estándar. El proyecto se ejecuta directamente con `go run`.
>
> **EN:** Go does not require a `go.mod` for single-file programs using only the standard library. The project runs directly with `go run`.

**Estructura de directorios esperada:**

```text
go/
└── core/
    └── foundations/
        └── helloworld/
            ├── hello_world.go    # Código fuente
            └── README.md         # Este archivo
```

**ES:** La ubicación sigue la convención `{lenguaje}/core/foundations/hello_world/` del repositorio principal. Al no haber `go.mod` ni subproyectos de pruebas, el proyecto consta exclusivamente del archivo fuente y su documentación.

**EN:** The location follows the `{language}/core/foundations/hello_world/` convention of the main repository. Since there is no `go.mod` nor test subprojects, the project consists exclusively of the source file and its documentation.

---

## 🛠️ Enfoque y construcción / Approach & Build

**ES:** Este proyecto fue creado **manualmente**, sin usar `go mod init`. Se optó por el enfoque más simple posible:

1. Crear el directorio `core/foundations/helloworld/` dentro del árbol `go/`.
2. Escribir el archivo fuente `hello_world.go` con la estructura mínima de un programa Go (`package main` + `func main()`).
3. Usar únicamente `fmt.Printf` de la biblioteca estándar, sin dependencias externas.

**EN:** This project was created **manually**, without using `go mod init`. The simplest possible approach was chosen:

1. Create the `core/foundations/helloworld/` directory inside the `go/` tree.
2. Write the source file `hello_world.go` with the minimal Go program structure (`package main` + `func main()`).
3. Use only `fmt.Printf` from the standard library, with no external dependencies.

### Comparación con Ada / Comparison with Ada

| Aspecto | Ada | Go |
|---------|-----|----|
| Scaffolding | `alire.toml` + `hello_world.gpr` + `.gitignore` | No requiere archivos de proyecto para programas simples |
| Compilación | `gprbuild -P hello_world.gpr` | `go build hello_world.go` o `go run hello_world.go` |
| Archivos mínimos | 3 archivos (fuente + proyecto + manifiesto) | 1 archivo (solo fuente) |
| Ejecutable | Se genera en `bin/` | Se genera en el directorio actual (con `go build`) |

---

## 📄 Archivos de configuración clave / Key Configuration Files

### `.gitignore` (raíz de `go/`)

**ES:** El `.gitignore` global del proyecto `go/` cubre los patrones típicos de Go: binarios (`*.exe`, `*.dll`, `*.so`, `*.dylib`), binarios de test (`*.test`), artefactos de cobertura (`*.out`, `coverage.*`), y el archivo de workspace (`go.work`). Para este programa monofichero no se genera ninguno de estos archivos, pero el patrón está definido para todo el árbol `go/`.

**EN:** The global `.gitignore` of the `go/` project covers typical Go patterns: binaries (`*.exe`, `*.dll`, `*.so`, `*.dylib`), test binaries (`*.test`), coverage artifacts (`*.out`, `coverage.*`), and the workspace file (`go.work`). For this single-file program none of these files are generated, but the pattern is defined for the entire `go/` tree.

```gitignore
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Code coverage profiles and other test artifacts
*.out
coverage.*
*.coverprofile
profile.cov

# Go workspace file
go.work
go.work.sum

# env file
.env
```

---

## 🚀 Compilación y ejecución / Build & Run

### Ejecutar directamente (sin compilar explícito)

```bash
# Desde go/core/foundations/helloworld/
go run hello_world.go
```

**Salida esperada / Expected output:**

```text
Hello, World! from Go!
```

### Compilar y ejecutar

```bash
# Compilar (genera hello_world.exe en Windows, hello_world en Unix)
go build hello_world.go

# Ejecutar
# En Windows:
hello_world.exe
# En Unix/macOS:
./hello_world
```

**Salida esperada / Expected output:**

```text
Hello, World! from Go!
```

### Desde cualquier directorio (ruta relativa)

```bash
# Desde la raíz del repositorio
go run go/core/foundations/helloworld/hello_world.go
```

---

## 📝 Notas de implementación / Implementation Notes

- **ES:** Go usa el paquete `fmt` de la biblioteca estándar para entrada/salida formateada. `fmt.Printf` funciona de manera similar a `printf` de C.
- **EN:** Go uses the `fmt` package from the standard library for formatted I/O. `fmt.Printf` works similarly to C's `printf`.
- **ES:** No se necesita un archivo `go.mod` porque el programa usa solo la biblioteca estándar y se ejecuta desde su propio directorio. Go 1.21+ permite ejecutar programas monofichero sin módulo.
- **EN:** No `go.mod` file is needed because the program uses only the standard library and runs from its own directory. Go 1.21+ allows running single-file programs without a module.
- **ES:** La función `main()` no recibe argumentos ni devuelve valores en el caso más simple; para programas que procesan argumentos se usa `os.Args`.
- **EN:** The `main()` function takes no arguments and returns no values in the simplest case; for programs that process arguments, `os.Args` is used.

---

### 🌐 Otras implementaciones / Other implementations

Este proyecto también está implementado en otros lenguajes. Explora el [repositorio principal](https://github.com/yorche3/programming_languages) para ver todas las versiones.

---

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
