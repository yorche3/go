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

func SelectionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}
	return arr
}

func InsertionSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}
