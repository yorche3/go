// naive_sort.go — Módulo Naive Sort con ordenamientos elementales
//
// Especificación: 05_Naive_Sort
//
// Funciones del contrato ([]int -> []int), de menor a mayor:
//
//	SelectionSort   — encuentra el mínimo del tramo no ordenado
//	BubbleSort      — compara e intercambia adyacentes, con bandera `swapped`
//	InsertionSort   — inserta cada elemento en su sub-array ordenado
//
// Caso nulo: Go representa la entrada inválida con el slice `nil`, que se
// retorna tal cual como indicador de fallo (equivalente al `int[]?` de C#).
package src
