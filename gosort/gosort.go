//Provides visualized algorithms
//for sorting and searching
package gosort

import (
	"math/rand"

	vs "../visualizer"
)

//region Members

var NumberOfElements int

type Sorter func([]int, vs.FrameGen)

//endregion

//region Public Methods

//Stable. Complexity: M(1) O(n) O(n^2) O(n^2)
func BubbleSort(arr []int, frameGen vs.FrameGen) {
	generateFrame(arr, frameGen)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				generateHighlightedFrame(arr, j+1, frameGen)
			}
			generateHighlightedFrame(arr, j+1, frameGen)

		}
	}
}

//Unstable. Complexity: M(1) O(n^2) O(n^2) O(n^2)
//Improved buble sort with less exchanges during sorting
func SelectionSort(arr []int, frameGen vs.FrameGen) {
	generateFrame(arr, frameGen)
	for i := 0; i < len(arr)-1; i++ {
		positionOfMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[positionOfMin] {
				positionOfMin = j
			}
			generateHighlightedFrame(arr, j, frameGen)
		}
		if positionOfMin != i {
			arr[i], arr[positionOfMin] = arr[positionOfMin], arr[i]
			generateHighlightedFrame(arr, i, frameGen)
		}
	}
}

//Stable. Complexity: M(1) O(n) O(n^2) O(n^2)
func InsertionSort(arr []int, frameGen vs.FrameGen) {
	generateFrame(arr, frameGen)

	for i := 1; i < len(arr); i++ {
		currentValue := arr[i]
		position := i

		for position > 0 && arr[position-1] > currentValue {
			generateHighlightedFrame(arr, position, frameGen)
			arr[position] = arr[position-1]
			position--
			generateHighlightedFrame(arr, position, frameGen)
		}

		arr[position] = currentValue
		generateHighlightedFrame(arr, position, frameGen)
	}
}

//O(n*n!)
//Learning purpose only. Don't use it in a real code.
func BogoSort(arr []int, frameGen vs.FrameGen) {
	generateFrame(arr, frameGen)
	for !isSorted(arr) {
		arr = shuffle(arr)
		generateFrame(arr, frameGen)
	}
}

//Stable. Complexity: M(1) O(n) O(n^2) O(n^2)
func CocktailSort(arr []int, frameGen vs.FrameGen) {
	generateFrame(arr, frameGen)

	last := len(arr) - 1
	for {
		swapped := false
		for i := 0; i < last; i++ {
			if arr[i] > arr[i+1] {
				generateHighlightedFrame(arr, i+1, frameGen)
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
			generateHighlightedFrame(arr, i, frameGen)
		}

		if !swapped {
			return
		}

		generateFrame(arr, frameGen)
		for i := last - 1; i >= 0; i-- {
			if arr[i] > arr[i+1] {
				generateHighlightedFrame(arr, i+1, frameGen)
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
			generateHighlightedFrame(arr, i, frameGen)
		}

		if !swapped {
			return
		}
	}
}

//Unstable. Complexity: M(1) O(n*logn) O(n^4/3) O(n^3/2)
func ShellSort(arr []int, frameGen vs.FrameGen) {
	generateFrame(arr, frameGen)
	count := len(arr)

	h := 1
	for h < count/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < count; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
				generateHighlightedFrame(arr, j, frameGen)
				generateHighlightedFrame(arr, j-h, frameGen)
			}

			generateHighlightedFrame(arr, i, frameGen)
		}
		h /= 3
	}
}

//endregion

//region Private Methods

func generateFrame(arr []int, frameGen vs.FrameGen) {
	if frameGen != nil {
		frameGen(arr, -1)
	}
}

func generateHighlightedFrame(arr []int, currentValue int, frameGen vs.FrameGen) {
	if frameGen != nil {
		frameGen(arr, currentValue)
	}
}

//Helping function for BogoSort
func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

//Helping function for BogoSort
//Shuffle array in random way
func shuffle(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

//endregion
