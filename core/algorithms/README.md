# Algorithms Pure — Go

Implementaciones de la [Fase 1 — Algoritmos Puros](https://yorche3.github.io/programming_languages/ROADMAP/#fase-1--algoritmos-puros--algorithms-pure-) en **Go**: ordenamientos elementales, estructuras de datos propias, ordenamientos óptimos y distribuidos, y búsqueda.

Los módulos de esta fase trabajan sobre slices (`[]int`), que **sí admiten el indicador de fallo del lenguaje** (`nil`) y pueden ordenarse *in-place*.

---

## 📂 Módulos / Modules

| Módulo | Especificación | Enfoque | Tests | Estado |
|--------|---------------|---------|:-----:|:------:|
| [`naive_sort/`](naive_sort/) | [05_Naive_Sort](https://yorche3.github.io/programming_languages/core/algorithms/05_Naive_Sort/) | `go test` + `testify` | 3 | ✅ |

---

## 📁 Estructura / Structure

```text
algorithms/
└── naive_sort/                     # 05_Naive_Sort
    ├── go.mod
    ├── go.sum
    ├── src/
    │   └── naive_sort.go           # SelectionSort, BubbleSort, InsertionSort
    ├── tests/
    │   └── naive_sort_test.go      # 3 tests × (7 casos + caso nulo)
    └── README.md
```

---

## 🛠️ Patrón común / Common Pattern

| Característica | Descripción |
|---------------|-------------|
| **Runtime** | Go (gc) — binario nativo, sin máquina virtual |
| **CLI** | `go build`, `go test`, `go vet`, `gofmt` |
| **Manifiesto** | `go.mod` — declara el módulo (`example.com/<name>`) y sus dependencias |
| **Lock** | `go.sum` — checksums de las dependencias |
| **Framework de tests** | `testify` — aserciones con mensaje (`assert.Equal(t, expected, actual, msg)`) |
| **Descubrimiento** | `go test ./...` ejecuta toda función `func TestXxx(t *testing.T)` |
| **Separación** | `src/` (paquete `src`) ↔ `tests/` (paquete `tests` con *dot-import*) |
| **Iteración** | Bucles `for` nativos; los slices se mutan *in-place* |
| **Naming** | Exportado en `PascalCase` (`SelectionSort`), exigido para cruzar de `src` a `tests` |
| **Indicador de fallo** | `nil` — el slice nulo se devuelve tal cual, sin `panic` |
| **Formato** | `gofmt -l .` sin salida (el CI no lo exige, pero se verifica igual) |

---

## 🚀 Compilación rápida / Quick Build

```bash
# Naive Sort Tests
cd naive_sort
go vet ./...
go test ./... -v
```

---

### 🌐 Otras implementaciones / Other implementations

Este proyecto también está implementado en otros lenguajes. Explora el [repositorio principal](https://github.com/yorche3/programming_languages) para ver todas las versiones.

---

## ▶️ Siguiente / Next

👉 Continúa con los módulos pendientes de esta fase en el [Roadmap](https://yorche3.github.io/programming_languages/ROADMAP/).
👉 Continue with the pending modules of this phase in the [Roadmap](https://yorche3.github.io/programming_languages/ROADMAP/).

---

*[← Volver a Core](../README.md)*

*🌐 [github.com/yorche3/programming_languages](https://github.com/yorche3/programming_languages) · [GitHub Pages](https://yorche3.github.io/programming_languages/)*
