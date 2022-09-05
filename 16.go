// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

package task

type Number interface {
	int64 | float64
}

func partition[T Number](arr []T, low, high int) ([]T, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func QuickSort[T Number](arr []T, low, high int) []T {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}
