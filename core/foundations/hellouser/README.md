# Hello, User! — Go

Implementación de la especificación [02_Hello_User](https://yorche3.github.io/programming_languages/core/foundations/02_Hello_User/) en **Go**, con el mismo enfoque **manual y minimalista** que [`hello_world`](../helloworld/).

---

## 📂 Archivos y estructura / Files & Structure

| Archivo / File | Propósito / Purpose |
|----------------|---------------------|
| [`hello_user.go`](hello_user.go) | Código fuente: solicita un nombre al usuario, lo almacena en una variable `name`, y muestra un saludo personalizado. |
| [`.gitignore`](../../../.gitignore) | Patrones globales para Go (idéntico al usado en `hello_world`). |

> **ES:** Al igual que `hello_world`, Go no requiere un `go.mod` para programas monofichero que solo usan la biblioteca estándar. Este proyecto se ejecuta directamente con `go run`.
>
> **EN:** Just like `hello_world`, Go does not require a `go.mod` for single-file programs using only the standard library. This project runs directly with `go run`.

**Estructura de directorios esperada:**

```text
go/
└── core/
    └── foundations/
        └── hellouser/
            ├── hello_user.go    # Código fuente
            └── README.md        # Este archivo
```

**ES:** Misma estructura que `hello_world`, solo cambia el nombre del archivo fuente y del directorio.

**EN:** Same structure as `hello_world`, only the source file and directory names change.

---

## 🛠️ Enfoque y construcción / Approach & Build

**ES:** Este proyecto fue creado **manualmente**, sin usar `go mod init`. Sigue la misma filosofía que `hello_world`:

1. Crear el directorio `core/foundations/hellouser/` dentro del árbol `go/`.
2. Escribir el archivo fuente `hello_user.go` con la estructura mínima de un programa Go.
3. Usar únicamente `fmt` de la biblioteca estándar para entrada (`fmt.Scanf`) y salida (`fmt.Print`, `fmt.Sprintf`, `fmt.Println`).

**EN:** This project was created **manually**, without using `go mod init`. It follows the same philosophy as `hello_world`:

1. Create the `core/foundations/hellouser/` directory inside the `go/` tree.
2. Write the source file `hello_user.go` with the minimal Go program structure.
3. Use only `fmt` from the standard library for input (`fmt.Scanf`) and output (`fmt.Print`, `fmt.Sprintf`, `fmt.Println`).

### Comparación con Ada / Comparison with Ada

| Aspecto | Ada | Go |
|---------|-----|----|
| Lectura de entrada | `Get_Line(Name, Length)` con buffer de tamaño fijo | `fmt.Scanf("%s", &name)` con string dinámico |
| Construcción de salida | Concatenación con `&` | `fmt.Sprintf("Hello, %s!", name)` |
| Gestión de memoria | Buffer fijo de 100 caracteres | `string` dinámico gestionado por el runtime |
| Archivos de configuración | `alire.toml` + `hello_user.gpr` + `.gitignore` | Ninguno (solo `.gitignore` global) |

---

## 📄 Archivos de configuración clave / Key Configuration Files

### `.gitignore` (raíz de `go/`)

**ES:** Mismo `.gitignore` global que `hello_world`. Ver [`hello_world/README.md`](../helloworld/README.md#gitignore-raíz-de-go) para la explicación detallada.

**EN:** Same global `.gitignore` as `hello_world`. See [`hello_world/README.md`](../helloworld/README.md#gitignore-raíz-de-go) for the detailed explanation.

---

## 🚀 Compilación y ejecución / Build & Run

### Ejecutar directamente (sin compilar explícito)

```bash
# Desde go/core/foundations/hellouser/
go run hello_user.go
```

**Salida esperada / Expected output:**

```text
What is your name? Go
Hello, Go!
```

### Compilar y ejecutar

```bash
# Compilar (genera hello_user.exe en Windows, hello_user en Unix)
go build hello_user.go

# Ejecutar
# En Windows:
hello_user.exe
# En Unix/macOS:
./hello_user
```

**Salida esperada / Expected output:**

```text
What is your name? Gopher
Hello, Gopher!
```

### Desde cualquier directorio (ruta relativa)

```bash
# Desde la raíz del repositorio
go run go/core/foundations/hellouser/hello_user.go
```

**Salida esperada / Expected output (ejemplo con entrada):**

```text
What is your name? Gopher
Hello, Gopher!
```

> **ES:** El programa espera una línea de entrada del usuario. Tras escribir el nombre y presionar Enter, muestra el saludo personalizado.
>
> **EN:** The program waits for a line of user input. After typing the name and pressing Enter, it displays the personalized greeting.

---

## 📝 Notas de implementación / Implementation Notes

- **ES:** A diferencia de Ada, que usa un buffer de tamaño fijo (`String(1 .. 100)`), Go maneja `string` de forma dinámica, sin límite predefinido de caracteres.
- **EN:** Unlike Ada, which uses a fixed-size buffer (`String(1 .. 100)`), Go handles `string` dynamically, with no predefined character limit.
- **ES:** `fmt.Scanf("%s", &name)` lee hasta el primer espacio en blanco. Para leer una línea completa (incluyendo espacios), se usaría `bufio.NewReader(os.Stdin).ReadString('\n')`. En este programa se optó por `fmt.Scanf` por simplicidad, siguiendo el mismo criterio minimalista del proyecto.
- **EN:** `fmt.Scanf("%s", &name)` reads until the first whitespace. To read a full line (including spaces), `bufio.NewReader(os.Stdin).ReadString('\n')` would be used. In this program `fmt.Scanf` was chosen for simplicity, following the same minimalist criterion as the project.
- **ES:** `fmt.Sprint` / `fmt.Sprintf` formatean cadenas sin imprimirlas directamente, permitiendo construir el mensaje antes de pasarlo a `fmt.Println`.
- **EN:** `fmt.Sprint` / `fmt.Sprintf` format strings without printing them directly, allowing the message to be built before passing it to `fmt.Println`.

---

### 🌐 Otras implementaciones / Other implementations

Este proyecto también está implementado en otros lenguajes. Explora el [repositorio principal](https://github.com/yorche3/programming_languages) para ver todas las versiones.

---

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
