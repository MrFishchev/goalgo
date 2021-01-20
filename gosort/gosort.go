//Provides visualized algorithms
//for sorting and searching
package gosort

import (
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

//endregion
